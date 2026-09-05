package battery

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
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
// itself: enough to evaluate any of its MaxK polynomials at its own index.
// Coefficients is indexed by pool key, then coefficient.
type ownContribution struct {
	Index        uint16
	Coefficients [][]*big.Int
}

// recoverPrivateShare rebuilds d_i = Σ_j f_j(i) for committee member myIdx
// exactly like node.buildPrivateShare: own contributions are evaluated
// locally, every other accepted contribution is located through its
// ContributionSubmitted event, its calldata transcript decoded and the share
// slot addressed to myIdx decrypted with the member's BabyJubJub secret.
func recoverPrivateShare(
	ctx context.Context, f *Fleet, epochID [12]byte, myIdx uint16, keyIndex uint8,
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
			share, err = ccommon.EvaluatePolynomialNative(own.Coefficients[keyIndex], big.NewInt(int64(myIdx)))
		} else {
			share, err = decryptShareFrom(ctx, f, epochID, addr, contribIdx, roundHash, myIdx, keyIndex, secret, fromBlock)
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
	contribIdx uint16, roundHash *big.Int, myIdx uint16, keyIndex uint8, secret *big.Int, fromBlock uint64,
) (*big.Int, error) {
	client := f.Services.Contracts.Client()
	data, err := finalizer.ContributionCalldata(ctx, client, f.Services.Manager, epochID, contributor, fromBlock)
	if err != nil {
		return nil, err
	}
	// The compact transcript's offsets are functions of the epoch's (t, n),
	// which come from the epoch policy, never from the calldata
	// (docs/pool-keys-v4.md §3, §5).
	e, err := f.epoch(ctx, epochID)
	if err != nil {
		return nil, err
	}
	layout, err := contribution.NewLayout(int(e.Policy.Threshold), int(e.Policy.CommitteeSize))
	if err != nil {
		return nil, err
	}
	tr, err := finalizer.DecodeContribution(data, layout)
	if err != nil {
		return nil, err
	}
	if myIdx == 0 || int(myIdx) > layout.CommitteeSize {
		return nil, fmt.Errorf("no share slot for index %d", myIdx)
	}
	// Decode enforces that recipient slot i carries index i+1.
	slot := int(myIdx) - 1
	ct := shareenc.Ciphertext{
		Ephemeral:   tr.Ephemerals[slot],
		MaskedShare: tr.MaskedShares[keyIndex][slot],
	}
	return shareenc.DecryptShareRoundHash(roundHash, contribIdx, myIdx, keyIndex, ct, secret)
}

// shareCommitmentMatches checks d·G against the share commitment the
// finalization published for participant idx under `keyIndex`. The
// commitments travel as finalizeEpoch calldata; only their Merkle roots are
// stored, so the check reads the transcript back from that transaction.
func shareCommitmentMatches(
	ctx context.Context, f *Fleet, epochID [12]byte, keyIndex uint8, idx uint16, d *big.Int, fromBlock uint64,
) (bool, error) {
	commitment, err := poolShareCommitment(ctx, f, epochID, keyIndex, idx, fromBlock)
	if err != nil {
		return false, err
	}
	point := helpers.ScalarBasePoint(d)
	return point.X.Cmp(commitment.X) == 0 && point.Y.Cmp(commitment.Y) == 0, nil
}

// finalizeCalldata reads back the calldata of the finalizeEpoch transaction
// that made the epoch Live (located through its EpochLive event). Only the
// Merkle roots of the share commitments are stored on chain, so everything
// the battery needs about them comes from here, decoded with the finalizer's
// hostile-calldata parsers.
func finalizeCalldata(ctx context.Context, f *Fleet, epochID [12]byte, fromBlock uint64) ([]byte, error) {
	return finalizer.FinalizeCalldata(ctx, f.Services.Contracts.Client(), f.Services.Manager, epochID, fromBlock)
}

// poolShareTree rebuilds one pool key's share-commitment tree, so a battery
// member can produce the Merkle path submitPartialDecryption asks for. The
// transcript carries D_{j,p} for every committee member p (slot p−1), whether
// or not p contributed, and the tree is laid out over exactly those.
func poolShareTree(
	ctx context.Context, f *Fleet, epochID [12]byte, keyIndex uint8, fromBlock uint64,
) (helpers.ShareTree, error) {
	data, err := finalizeCalldata(ctx, f, epochID, fromBlock)
	if err != nil {
		return helpers.ShareTree{}, err
	}
	indexes, commitments, err := finalizer.FinalizeShareCommitments(data, keyIndex)
	if err != nil {
		return helpers.ShareTree{}, fmt.Errorf("decode finalization share commitments: %w", err)
	}
	return helpers.NewShareTree(indexes, commitments)
}

// poolShareCommitment reads D_{j,idx} for one committee member out of the
// finalization transcript, key `keyIndex`.
func poolShareCommitment(
	ctx context.Context, f *Fleet, epochID [12]byte, keyIndex uint8, idx uint16, fromBlock uint64,
) (types.CurvePoint, error) {
	data, err := finalizeCalldata(ctx, f, epochID, fromBlock)
	if err != nil {
		return types.CurvePoint{}, err
	}
	_, commitments, err := finalizer.FinalizeShareCommitments(data, keyIndex)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode finalization share commitments: %w", err)
	}
	if idx == 0 || int(idx) > len(commitments) {
		return types.CurvePoint{}, fmt.Errorf("member %d is outside the committee of %d", idx, len(commitments))
	}
	return commitments[idx-1], nil
}
