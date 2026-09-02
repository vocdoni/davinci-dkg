package dleq

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

type organizerFixture struct {
	eid   [12]byte
	aid   [32]byte
	ctIdx uint16
	skOrg *big.Int
	pkOrg types.CurvePoint
	c1    types.CurvePoint
	delta types.CurvePoint
	proof Proof
}

func newOrganizerFixture(t *testing.T) organizerFixture {
	t.Helper()
	c := qt.New(t)

	f := organizerFixture{
		eid:   [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x77},
		aid:   [32]byte{31: 0x07},
		ctIdx: 3,
		skOrg: big.NewInt(1234567890),
	}
	pk := group.NewPoint()
	pk.ScalarBaseMult(f.skOrg)
	f.pkOrg = group.Encode(pk)
	c1 := group.NewPoint()
	c1.ScalarBaseMult(big.NewInt(999999))
	f.c1 = group.Encode(c1)

	delta, proof, err := ProveOrganizerShare(f.eid, f.aid, f.ctIdx, f.skOrg, f.c1)
	c.Assert(err, qt.IsNil)
	f.delta, f.proof = delta, proof
	return f
}

func TestProveOrganizerShareRoundTrip(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	// Δ must be the real sk_org·C1, not just something the DLEQ closes over.
	c1Point, err := group.Decode(f.c1)
	c.Assert(err, qt.IsNil)
	want := group.NewPoint()
	want.ScalarMult(c1Point, f.skOrg)
	c.Assert(f.delta.X.Cmp(group.Encode(want).X), qt.Equals, 0)
	c.Assert(f.delta.Y.Cmp(group.Encode(want).Y), qt.Equals, 0)

	c.Assert(f.proof.Response.Cmp(group.ScalarField()), qt.Equals, -1)
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, f.proof), qt.IsTrue)
}

func TestProveOrganizerShareUsesFreshNonce(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	_, second, err := ProveOrganizerShare(f.eid, f.aid, f.ctIdx, f.skOrg, f.c1)
	c.Assert(err, qt.IsNil)
	// A deterministic nonce would leak sk_org the moment two challenges differ.
	c.Assert(second.A1.X.Cmp(f.proof.A1.X) == 0 && second.A1.Y.Cmp(f.proof.A1.Y) == 0, qt.IsFalse)
}

func TestVerifyOrganizerShareRejectsTamperedResponse(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	tampered := f.proof
	tampered.Response = new(big.Int).Add(f.proof.Response, big.NewInt(1))
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, tampered), qt.IsFalse)

	// A non-canonical z (z + q verifies algebraically) must be rejected too,
	// otherwise one share has two on-chain encodings.
	tampered.Response = new(big.Int).Add(f.proof.Response, group.ScalarField())
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, tampered), qt.IsFalse)

	tampered.Response = nil
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, tampered), qt.IsFalse)
}

func TestVerifyOrganizerShareBindsContext(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	otherAid := f.aid
	otherAid[31] = 0x08
	c.Assert(VerifyOrganizerShare(f.eid, otherAid, f.ctIdx, f.pkOrg, f.c1, f.delta, f.proof), qt.IsFalse)

	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx+1, f.pkOrg, f.c1, f.delta, f.proof), qt.IsFalse)

	otherEid := f.eid
	otherEid[11] = 0x78
	c.Assert(VerifyOrganizerShare(otherEid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, f.proof), qt.IsFalse)
}

func TestVerifyOrganizerShareRejectsMalformedPoints(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	identity := types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	// Δ = O would let a combine proceed with no organizer contribution at all.
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, identity, f.proof), qt.IsFalse)
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, identity, f.c1, f.delta, f.proof), qt.IsFalse)

	offCurve := types.CurvePoint{X: big.NewInt(2), Y: big.NewInt(3)}
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, offCurve, f.proof), qt.IsFalse)

	tampered := f.proof
	tampered.A1 = offCurve
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, tampered), qt.IsFalse)

	// A share proved for another organizer key must not verify against this one.
	otherPK := group.NewPoint()
	otherPK.ScalarBaseMult(big.NewInt(42))
	c.Assert(VerifyOrganizerShare(f.eid, f.aid, f.ctIdx, group.Encode(otherPK), f.c1, f.delta, f.proof), qt.IsFalse)
}

func TestProveOrganizerShareRejectsBadInputs(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	_, _, err := ProveOrganizerShare(f.eid, f.aid, f.ctIdx, nil, f.c1)
	c.Assert(err, qt.Not(qt.IsNil))

	_, _, err = ProveOrganizerShare(f.eid, f.aid, f.ctIdx, group.ScalarField(), f.c1)
	c.Assert(err, qt.Not(qt.IsNil))

	_, _, err = ProveOrganizerShare(f.eid, f.aid, f.ctIdx, f.skOrg, types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)})
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestOrganizerShareChallengeIsDomainSeparated(t *testing.T) {
	c := qt.New(t)
	f := newOrganizerFixture(t)

	e := OrganizerShareChallenge(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, f.proof.A1, f.proof.A2)
	c.Assert(e.Sign() > 0, qt.IsTrue)
	c.Assert(e.Cmp(group.ScalarField()), qt.Equals, -1)
	// Deterministic in its inputs, and every input is bound.
	c.Assert(OrganizerShareChallenge(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, f.proof.A1, f.proof.A2).Cmp(e),
		qt.Equals, 0)
	c.Assert(OrganizerShareChallenge(f.eid, f.aid, f.ctIdx+1, f.pkOrg, f.c1, f.delta, f.proof.A1, f.proof.A2).Cmp(e),
		qt.Not(qt.Equals), 0)
	c.Assert(OrganizerShareChallenge(f.eid, f.aid, f.ctIdx, f.pkOrg, f.c1, f.delta, f.proof.A2, f.proof.A1).Cmp(e),
		qt.Not(qt.Equals), 0)
}
