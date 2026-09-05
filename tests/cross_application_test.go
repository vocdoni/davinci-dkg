package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestCrossApplicationCiphertextIsUseless is the regression test for the
// decryption oracle pool keys close. Two applications of the same epoch hold
// different committee keys (P_0 and P_1), so a ciphertext of application A
// copied verbatim into application B yields partials under B's key and
// decrypts to nothing: the residual C2 − Σ λ δ' is r·(P_0 − P_1) away from
// the real plaintext, and no combine over B can ever be proven for it.
//
// It also asserts the contract-level half of the fix: a partial computed
// with A's share cannot be posted under B, because its share commitment is
// not a leaf of B's pool share root.
func TestCrossApplicationCiphertextIsUseless(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()
	_, err := helpers.ActivateRoundPoolKey(ctx, services, res, 1)
	c.Assert(err, qt.IsNil)

	automatic := golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)}
	aidA, aidB := randomAid(c), randomAid(c)
	c.Assert(helpers.RegisterApplication(ctx, self, services.AppManager, res.EpochID, aidA, nil, automatic), qt.IsNil)
	c.Assert(helpers.RegisterApplication(ctx, self, services.AppManager, res.EpochID, aidB, nil, automatic), qt.IsNil)

	appA, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aidA)
	c.Assert(err, qt.IsNil)
	appB, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aidB)
	c.Assert(err, qt.IsNil)
	c.Assert(appA.PoolIndex, qt.Not(qt.Equals), appB.PoolIndex, qt.Commentf("each application holds its own committee key"))

	activationA, activationB := res.Activation(appA.PoolIndex), res.Activation(appB.PoolIndex)
	c.Assert(activationA.PoolKey.X.Cmp(activationB.PoolKey.X), qt.Not(qt.Equals), 0)
	shareA := poolShare(c, res, appA.PoolIndex, 1)
	shareB := poolShare(c, res, appB.PoolIndex, 1)

	// A genuine ciphertext for A, with randomness nobody else knows.
	plaintext := big.NewInt(4242)
	c1, c2, err := elgamal.Encrypt(activationA.PoolKey, plaintext)
	c.Assert(err, qt.IsNil)

	idxA, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aidA, c1, c2)
	c.Assert(err, qt.IsNil)

	// ── the honest path: A's own key recovers the plaintext ───────────────
	partialA, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, res.EpochID, aidA, idxA, 1, c1, c2, shareA, big.NewInt(7), activationA.Shares,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aidA, 1, idxA, partialA), qt.IsNil)
	combineA, err := helpers.BuildDecryptCombineOutputFromCiphertext(
		ctx, res.EpochID, aidA, idxA, 1, c1, c2, helpers.IdentityPoint(), nil,
		[]uint16{1}, []types.CurvePoint{partialA.Delta}, plaintext,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aidA, idxA, combineA), qt.IsNil)
	record, err := helpers.WaitCombinedDecryption(ctx, services, res.EpochID, aidA, idxA)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Plaintext.String(), qt.Equals, plaintext.String())

	// ── the attack: the same ciphertext copied into B ─────────────────────
	idxB, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aidB, c1, c2)
	c.Assert(err, qt.IsNil)

	// A's partial cannot be replayed under B: its share commitment is not a
	// leaf of B's pool share root.
	replayed, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, res.EpochID, aidB, idxB, 1, c1, c2, shareA, big.NewInt(8), activationB.Shares,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aidB, 1, idxB, replayed), qt.IsNotNil,
		qt.Commentf("A's share must not verify against B's pool share root"))

	// The committee's honest work under B produces δ' = e_{1,i}·C1 for B's
	// key. The contract accepts it — and it is worthless.
	partialB, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, res.EpochID, aidB, idxB, 1, c1, c2, shareB, big.NewInt(9), activationB.Shares,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aidB, 1, idxB, partialB), qt.IsNil)

	// The residual is not the plaintext: C2 − δ' = m·G + r·(P_0 − P_1).
	residual := subtractPoints(c, c2, partialB.Delta)
	c.Assert(residual, qt.Not(qt.Equals), pointString(helpers.ScalarBasePoint(plaintext)),
		qt.Commentf("B's partials must not reveal A's plaintext"))

	// And no combine over B can be proven for it, so the plaintext never
	// reaches the chain either.
	_, err = helpers.BuildDecryptCombineOutputFromCiphertext(
		ctx, res.EpochID, aidB, idxB, 1, c1, c2, helpers.IdentityPoint(), nil,
		[]uint16{1}, []types.CurvePoint{partialB.Delta}, plaintext,
	)
	c.Assert(err, qt.IsNotNil, qt.Commentf("the combine equation cannot hold under a different pool key"))

	combined, err := services.Contracts.GetCombinedDecryption(ctx, res.EpochID, aidB, idxB)
	c.Assert(err, qt.IsNil)
	c.Assert(combined.Completed, qt.IsFalse)
}

// subtractPoints returns a − b as a comparable string.
func subtractPoints(c *qt.C, a, b types.CurvePoint) string {
	left, err := group.Decode(a)
	c.Assert(err, qt.IsNil)
	right, err := group.Decode(b)
	c.Assert(err, qt.IsNil)
	neg := group.NewPoint()
	neg.Neg(right)
	sum := group.NewPoint()
	sum.Add(left, neg)
	return pointString(group.Encode(sum))
}

func pointString(p types.CurvePoint) string {
	return p.X.String() + ":" + p.Y.String()
}
