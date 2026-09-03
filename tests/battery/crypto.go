package battery

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/finalizer"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// nodeKeyDomain must match bjjKeyDomain in cmd/davinci-dkg-node and
// localNodeKeyDerivationDomain in tests/helpers/nodekeys.go: the battery
// registers its adversarial operators through helpers.EnsureNodeKeyRegistered
// and then needs the matching secret to decrypt the shares sent to it.
const nodeKeyDomain = "davinci-dkg/bjj-key/v1"

// randomAid returns an application id below the BN254 scalar field.
func randomAid() ([32]byte, error) {
	var aid [32]byte
	if _, err := rand.Read(aid[:]); err != nil {
		return aid, err
	}
	aid[0] &= 0x1f
	return aid, nil
}

// randomScalar returns a uniform element of [1, q).
func randomScalar() (*big.Int, error) {
	order := group.ScalarField()
	s, err := rand.Int(rand.Reader, new(big.Int).Sub(order, big.NewInt(1)))
	if err != nil {
		return nil, err
	}
	return s.Add(s, big.NewInt(1)), nil
}

// randomScalars returns k random non-zero scalars.
func randomScalars(k int) ([]*big.Int, error) {
	out := make([]*big.Int, k)
	for i := range out {
		s, err := randomScalar()
		if err != nil {
			return nil, err
		}
		out[i] = s
	}
	return out, nil
}

// randomPlaintext returns a uniform integer in [1, 2^bits).
func randomPlaintext(bits uint) (*big.Int, error) {
	m, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), bits))
	if err != nil {
		return nil, err
	}
	if m.Sign() == 0 {
		m.SetInt64(1)
	}
	return m, nil
}

// nodeSecret mirrors the node binary's deterministic BabyJubJub key
// derivation from the operator's EVM private key.
func nodeSecret(privKeyHex string) (*big.Int, error) {
	preimage := append(common.FromHex(privKeyHex), []byte(nodeKeyDomain)...)
	digest := ethcrypto.Keccak256(preimage)
	lo := new(big.Int).SetBytes(digest[:16])
	hi := new(big.Int).SetBytes(digest[16:])
	secret, err := dkghash.HashFieldElements(lo, hi)
	if err != nil {
		return nil, fmt.Errorf("poseidon: %w", err)
	}
	secret.Mod(secret, group.ScalarField())
	if secret.Sign() == 0 {
		secret.SetInt64(1)
	}
	return secret, nil
}

// torsionPoint returns (0, −1): the order-2 point of BabyJubJub in reduced
// twisted-Edwards form (−x² + y² = 1 + d·x²·y²). It is canonical, on the
// curve and non-identity — so the contract accepts it — but lies in the
// cofactor subgroup, so no honest ciphertext ever contains it.
func torsionPoint() types.CurvePoint {
	return types.CurvePoint{X: big.NewInt(0), Y: new(big.Int).Sub(group.BaseField(), big.NewInt(1))}
}

// addPoints returns a + b on the full curve (either operand may sit outside
// the prime-order subgroup).
func addPoints(a, b types.CurvePoint) (types.CurvePoint, error) {
	pa, err := group.Decode(a)
	if err != nil {
		return types.CurvePoint{}, err
	}
	pb, err := group.Decode(b)
	if err != nil {
		return types.CurvePoint{}, err
	}
	sum := group.NewPoint()
	sum.Add(pa, pb)
	return group.Encode(sum), nil
}

// randomSubgroupPoint returns r·G for a random r.
func randomSubgroupPoint() (types.CurvePoint, error) {
	r, err := randomScalar()
	if err != nil {
		return types.CurvePoint{}, err
	}
	return helpers.ScalarBasePoint(r), nil
}

// ownContribution is what the battery keeps of a contribution it made
// itself: enough to evaluate its own polynomial at its own index.
type ownContribution struct {
	Index        uint16
	Coefficients []*big.Int
}

// recoverPrivateShare rebuilds d_i = Σ_j f_j(i) for committee member myIdx
// exactly like node.buildPrivateShare: own contributions are evaluated
// locally, every other accepted contribution is located through its
// ContributionSubmitted event, its calldata transcript decoded and the share
// slot addressed to myIdx decrypted with the member's BabyJubJub secret.
func recoverPrivateShare(
	ctx context.Context, f *Fleet, epochID [12]byte, myIdx uint16,
	committee []common.Address, secret *big.Int, own *ownContribution, fromBlock uint64,
) (*big.Int, int, error) {
	roundHash := helpers.RoundScalar(epochID)
	modulus := group.ScalarField()
	total := new(big.Int)
	accepted := 0
	for i, addr := range committee {
		contribIdx := uint16(i + 1)
		rec, err := f.Services.Manager.GetContribution(f.callOpts(ctx), epochID, addr)
		if err != nil {
			return nil, 0, fmt.Errorf("get contribution of %s: %w", addr.Hex(), err)
		}
		if !rec.Accepted {
			continue
		}
		accepted++
		var share *big.Int
		if own != nil && own.Index == contribIdx {
			share, err = ccommon.EvaluatePolynomialNative(own.Coefficients, big.NewInt(int64(myIdx)))
		} else {
			share, err = decryptShareFrom(ctx, f, epochID, addr, contribIdx, roundHash, myIdx, secret, fromBlock)
		}
		if err != nil {
			return nil, 0, fmt.Errorf("share from %s (idx %d): %w", addr.Hex(), contribIdx, err)
		}
		total.Add(total, share)
		total.Mod(total, modulus)
	}
	if accepted == 0 {
		return nil, 0, fmt.Errorf("no accepted contributions")
	}
	return total, accepted, nil
}

func decryptShareFrom(
	ctx context.Context, f *Fleet, epochID [12]byte, contributor common.Address,
	contribIdx uint16, roundHash *big.Int, myIdx uint16, secret *big.Int, fromBlock uint64,
) (*big.Int, error) {
	client := f.Services.Contracts.Client()
	data, err := finalizer.ContributionCalldata(ctx, client, f.Services.Manager, epochID, contributor, fromBlock)
	if err != nil {
		return nil, err
	}
	transcript, err := finalizer.ContributionTranscript(data)
	if err != nil {
		return nil, err
	}
	// Layout (N = MaxN words of 32 bytes): [0..2N) commitments, [2N..3N)
	// recipient indexes, [3N..5N) recipient keys, [5N..7N) ephemerals,
	// [7N..8N) masked shares.
	const n = ccommon.MaxN
	word := func(i int) *big.Int { return new(big.Int).SetBytes(transcript[i*32 : (i+1)*32]) }
	for slot := range n {
		if uint16(word(2*n+slot).Uint64()) != myIdx {
			continue
		}
		ct := shareenc.Ciphertext{
			Ephemeral:   types.CurvePoint{X: word(5*n + 2*slot), Y: word(5*n + 2*slot + 1)},
			MaskedShare: word(7*n + slot),
		}
		return shareenc.DecryptShareRoundHash(roundHash, contribIdx, myIdx, ct, secret)
	}
	return nil, fmt.Errorf("no share slot for index %d", myIdx)
}

// shareCommitmentMatches checks d·G against the share commitment hash the
// finalize proof pinned on chain for participant idx.
func shareCommitmentMatches(ctx context.Context, f *Fleet, epochID [12]byte, idx uint16, d *big.Int) (bool, error) {
	stored, err := f.Services.Manager.GetShareCommitmentHash(f.callOpts(ctx), epochID, idx)
	if err != nil {
		return false, err
	}
	point := helpers.ScalarBasePoint(d)
	digest := ethcrypto.Keccak256Hash(common.LeftPadBytes(point.X.Bytes(), 32), common.LeftPadBytes(point.Y.Bytes(), 32))
	return digest == common.Hash(stored), nil
}
