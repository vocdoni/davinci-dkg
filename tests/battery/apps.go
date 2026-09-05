package battery

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
)

// application is one registered (epoch, aid) with the organizer secret the
// battery holds for it. PKAid = P_j + PK_org is what ciphertexts are
// encrypted under, with P_j the pool key the registration claimed; an
// automatic application has PK_org = identity and SkOrg = nil.
type application struct {
	Epoch     [12]byte
	Aid       [32]byte
	SkOrg     *big.Int
	PKOrg     types.CurvePoint
	PKAid     types.CurvePoint
	PoolIndex uint8
	Automatic bool
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
// derives the application key from the pool key the registration claimed.
// In automatic mode the key and Schnorr words are zero and skOrg is ignored.
func (f *Fleet) registerApplicationWith(
	ctx context.Context, a *actor, epoch [12]byte, aid [32]byte, skOrg *big.Int, policy golangtypes.DKGTypesAppPolicy,
) (*application, txOutcome, error) {
	automatic := policy.Mode == uint8(types.AppModeAutomatic)
	// A registration claims the key at the pool cursor and reverts
	// PoolKeyNotActive until a node has activated it (one proof per key, a
	// tick or two after the previous claim); wait for that proof rather than
	// reporting the race as a failure.
	if _, err := f.waitPoolKeys(ctx, epoch, 1, activationWait(), 0); err != nil {
		return nil, txOutcome{}, err
	}
	pkX, pkY := new(big.Int), new(big.Int)
	ax, ay, z := new(big.Int), new(big.Int), new(big.Int)
	if !automatic {
		organizerX, organizerY, proof, err := schnorr.ProveOrganizerRegister(skOrg, epoch, aid)
		if err != nil {
			return nil, txOutcome{}, fmt.Errorf("organizer schnorr proof: %w", err)
		}
		pkX, pkY = organizerX, organizerY
		ax, ay, z = proof.Ax, proof.Ay, proof.Z
	}
	out, err := f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.AppManager.RegisterApplication(auth, epoch, aid, policy, pkX, pkY, ax, ay, z)
	})
	if err != nil {
		return nil, out, err
	}

	record, err := f.Services.AppManager.GetApplication(f.callOpts(ctx), epoch, aid)
	if err != nil {
		return nil, out, fmt.Errorf("read application: %w", err)
	}
	poolKey, err := f.poolKey(ctx, epoch, record.PoolIndex)
	if err != nil {
		return nil, out, fmt.Errorf("pool key %d: %w", record.PoolIndex, err)
	}
	pkOrg := types.CurvePoint{X: record.OrganizerPK.X, Y: record.OrganizerPK.Y}
	pkAid, err := elgamal.ApplicationKey(poolKey, pkOrg)
	if err != nil {
		return nil, out, fmt.Errorf("application key: %w", err)
	}
	app := &application{
		Epoch: epoch, Aid: aid, PKOrg: pkOrg, PKAid: pkAid,
		PoolIndex: record.PoolIndex, Automatic: automatic, Organizer: a,
	}
	if !automatic {
		app.SkOrg = skOrg
	}
	return app, out, nil
}

// encrypt produces an honest ElGamal ciphertext of m under PK_aid.
func (app *application) encrypt(m *big.Int) (c1, c2 types.CurvePoint, err error) {
	return elgamal.Encrypt(app.PKAid, m)
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

// revealSecret publishes sk_org for an organizer-locked application. The
// call is permissionless, so `a` need not be the organizer.
func (f *Fleet) revealSecret(ctx context.Context, a *actor, app *application, secret *big.Int) (txOutcome, error) {
	return f.send(ctx, a, func(auth *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return f.Services.AppManager.RevealOrganizerSecret(auth, app.Epoch, app.Aid, secret)
	})
}

// releaseSecret reveals the honest organizer secret from the organizer's own
// address: the one-shot that lets the committee finish every ciphertext of
// the application, past and future.
func (f *Fleet) releaseSecret(ctx context.Context, app *application) (txOutcome, error) {
	if app.Automatic {
		return txOutcome{}, fmt.Errorf("automatic application %x has no secret to reveal", app.Aid)
	}
	return f.revealSecret(ctx, app.Organizer, app, app.SkOrg)
}

// automaticPolicy is the default the battery registers with: no organizer
// half, so the fleet owns every ciphertext of the application end to end.
func automaticPolicy() golangtypes.DKGTypesAppPolicy {
	return golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)}
}

// automaticPolicyWith forces `policy` into automatic mode.
func automaticPolicyWith(policy golangtypes.DKGTypesAppPolicy) golangtypes.DKGTypesAppPolicy {
	policy.Mode = uint8(types.AppModeAutomatic)
	return policy
}
