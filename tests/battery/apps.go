package battery

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
)

// application is one registered (epoch, aid) with the organizer secret the
// battery holds for it. PKAid = PK_ep + PK_org is what ciphertexts are
// encrypted under.
type application struct {
	Epoch     [12]byte
	Aid       [32]byte
	SkOrg     *big.Int
	PKOrg     types.CurvePoint
	PKAid     types.CurvePoint
	Organizer *actor
}

// registerApplication registers a fresh aid with a fresh organizer secret.
func (f *Fleet) registerApplication(
	ctx context.Context, a *actor, epoch [12]byte, policy golangtypes.DKGTypesAppPolicy,
) (*application, txOutcome, error) {
	aid, err := randomAid()
	if err != nil {
		return nil, txOutcome{}, err
	}
	skOrg, err := randomScalar()
	if err != nil {
		return nil, txOutcome{}, err
	}
	return f.registerApplicationWith(ctx, a, epoch, aid, skOrg, policy)
}

// registerApplicationWith registers a given aid / organizer secret and
// derives the application key from the epoch's collective key.
func (f *Fleet) registerApplicationWith(
	ctx context.Context, a *actor, epoch [12]byte, aid [32]byte, skOrg *big.Int, policy golangtypes.DKGTypesAppPolicy,
) (*application, txOutcome, error) {
	pkX, pkY, proof, err := schnorr.ProveOrganizerRegister(skOrg, epoch, aid)
	if err != nil {
		return nil, txOutcome{}, fmt.Errorf("organizer schnorr proof: %w", err)
	}
	out, err := f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.AppManager.RegisterApplication(auth, epoch, aid, policy, pkX, pkY, proof.Ax, proof.Ay, proof.Z)
	})
	if err != nil {
		return nil, out, err
	}
	pkEp, err := f.collectiveKey(ctx, epoch)
	if err != nil {
		return nil, out, fmt.Errorf("collective key: %w", err)
	}
	pkOrg := types.CurvePoint{X: pkX, Y: pkY}
	pkAid, err := elgamal.ApplicationKey(pkEp, pkOrg)
	if err != nil {
		return nil, out, fmt.Errorf("application key: %w", err)
	}
	return &application{Epoch: epoch, Aid: aid, SkOrg: skOrg, PKOrg: pkOrg, PKAid: pkAid, Organizer: a}, out, nil
}

// encrypt produces an honest ElGamal ciphertext of m under PK_aid.
func (app *application) encrypt(m *big.Int) (c1, c2 types.CurvePoint, err error) {
	return elgamal.Encrypt(app.PKAid, m)
}

// share derives the organizer words for one ciphertext. The DLEQ nonce is
// fresh on every call, so the words that land on chain are the only ones a
// combine may carry.
func (app *application) share(idx uint16, c1 types.CurvePoint) (types.CurvePoint, dleq.Proof, error) {
	return dleq.ProveOrganizerShare(app.Epoch, app.Aid, idx, app.SkOrg, c1)
}

// submitCiphertext posts (c1, c2) and returns the index the contract
// assigned from the CiphertextSubmitted event in the receipt.
func (f *Fleet) submitCiphertext(
	ctx context.Context, a *actor, epoch [12]byte, aid [32]byte, c1, c2 types.CurvePoint,
) (uint16, txOutcome, error) {
	out, err := f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.Manager.SubmitCiphertext(auth, epoch, aid, c1.X, c1.Y, c2.X, c2.Y)
	})
	if err != nil {
		return 0, out, err
	}
	receipt, err := f.Services.Contracts.Client().TransactionReceipt(ctx, out.Hash)
	if err != nil {
		return 0, out, fmt.Errorf("ciphertext receipt: %w", err)
	}
	for _, lg := range receipt.Logs {
		if ev, perr := f.Services.Manager.ParseCiphertextSubmitted(*lg); perr == nil {
			return ev.CiphertextIndex, out, nil
		}
	}
	return 0, out, fmt.Errorf("CiphertextSubmitted event not found in tx %s", out.Hash.Hex())
}

// submitShareWords posts arbitrary organizer-share words for a slot. Used
// both for honest releases and for the adversarial (tampered / replayed /
// relayed) variants, so it takes the words explicitly.
func (f *Fleet) submitShareWords(
	ctx context.Context, a *actor, epoch [12]byte, aid [32]byte, idx uint16,
	c1, c2, delta types.CurvePoint, proof dleq.Proof,
) (txOutcome, error) {
	return f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.AppManager.SubmitOrganizerShare(auth, epoch, aid, idx,
			c1.X, c1.Y, c2.X, c2.Y, delta.X, delta.Y,
			proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y, proof.Response)
	})
}

// releaseShare proves and posts the honest organizer share from the
// organizer's own address.
func (f *Fleet) releaseShare(
	ctx context.Context, app *application, idx uint16, c1, c2 types.CurvePoint,
) (types.CurvePoint, dleq.Proof, txOutcome, error) {
	delta, proof, err := app.share(idx, c1)
	if err != nil {
		return delta, proof, txOutcome{}, err
	}
	out, err := f.submitShareWords(ctx, app.Organizer, app.Epoch, app.Aid, idx, c1, c2, delta, proof)
	return delta, proof, out, err
}
