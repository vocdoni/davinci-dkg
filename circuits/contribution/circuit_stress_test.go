package contribution

// circuit_stress_test.go – exhaustive prove+verify stress test for the
// contribution circuit.  Reproduces the non-deterministic "pairing doesn't
// match" error observed during the MaxN=32 gas benchmark.
//
// Run (no integration infra needed):
//
//	go test -v -count=1 -run TestContributionCircuitStress \
//	  -timeout 30m ./circuits/contribution/...
//
// To detect data races in gnark's prover:
//
//	go test -race -count=1 -run TestContributionCircuitStress \
//	  -timeout 60m ./circuits/contribution/...

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/types"
)

// anvilPrivateKeys mirrors tests/helpers/constants.go – the standard Anvil
// deterministic accounts used by the integration tests.
var anvilPrivateKeys = []string{
	"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
	"0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
	"0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
	"0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e",
	"0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356",
	"0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
	"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
	"0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897",
	"0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82",
	"0xa267530f49f8280200edf313ee7af6b827f2a8bce2897751d06a843f644967b1",
	"0x47c99abed3324a2707c28affff1267e45918ec8c3f20b8aa892e8b065d2942dd",
	"0xc526ee95bf44d8fc405a158bb884d9d1238d99f0612e9f33d006bb0789009aaa",
	"0x8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61",
	"0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0",
	"0x689af8efa8c651a91ad287602527f3af2fe9f6501a7ac4b061667b5a93e037fd",
	"0xde9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0",
	"0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e",
	"0xeaa861a9a01391ed3d587d8a5a84ca56ee277629a8b02c22093a419bf240e65d",
	"0xc511b2aa70776d4ff1d376e8537903dae36896132c90b91d52c1dfbae267cd8b",
	"0x224b7eb7449992aac96d631d9677f7bf5888245eef6d6eeda31e62d2f29a83e4",
	"0x4624e0802698b9769f5bdb260a3777fbd4941ad2901f5966b854f953497eec1b",
	"0x375ad145df13ed97f8ca8e27bb21ebf2a3819e9e0a06509a812db377e533def7",
	"0x18743e59419b01d1d846d97ea070b5a3368a3e7f6f0242cf497e1baac6972427",
	"0xe383b226df7c8282489889170b0f68f66af6459261f4833a781acd0804fafe7a",
	"0xf3a6b71b94f5cd909fb2dbb287da47badaa6d8bcdc45d595e2884835d8749001",
	"0x4e249d317253b9641e477aba8dd5d8f1f7cf5250a5acadd1229693e262720a19",
	"0x233c86e887ac435d7f7dc64979d7758d69320906a0d340d2b6518b0fd20aa998",
	"0x85a74ca11529e215137ccffd9c95b2c72c5fb0295c973eb21032e823329b3d2d",
	"0xac8698a440d33b866b6ffe8775621ce1a4e6ebd04ab7980deb97b3d997fc64fb",
}

const localNodeKeyDerivationDomain = "davinci-dkg/bjj-key/v1"

// actorBJJKey derives the BJJ public key for an Anvil account — mirrors
// tests/helpers.deterministicNodeKeyMaterial.
func actorBJJKey(privKeyHex string) types.NodeKey {
	preimage := append(common.FromHex(privKeyHex), []byte(localNodeKeyDerivationDomain)...)
	digest := ethcrypto.Keccak256(preimage)
	lo := new(big.Int).SetBytes(digest[:16])
	hi := new(big.Int).SetBytes(digest[16:])
	secret, _ := dkghash.HashFieldElements(lo, hi)
	secret.Mod(secret, group.ScalarField())
	if secret.Sign() == 0 {
		secret.SetInt64(1)
	}
	pub := group.NewPoint()
	pub.ScalarBaseMult(secret)
	enc := group.Encode(pub)
	return types.NodeKey{PubX: enc.X, PubY: enc.Y}
}

// stressRoundHashes returns a representative set of roundHash values:
// small, medium, near-max, and random 12-byte values (simulating real roundIDs).
func stressRoundHashes(n int) []*big.Int {
	hashes := []*big.Int{
		big.NewInt(1),
		big.NewInt(12345),
		big.NewInt(99999),
		new(big.Int).SetBytes([]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe, 0x00, 0x00, 0x00, 0x01}),
		new(big.Int).SetBytes([]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe, 0x00, 0x00, 0x00, 0x02}),
		new(big.Int).SetBytes([]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe, 0x00, 0x00, 0x00, 0x03}),
		new(big.Int).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}),
	}
	// Add random 12-byte values mimicking Solidity round IDs.
	rnd := mathrand.New(mathrand.NewSource(42))
	for i := 0; i < n; i++ {
		buf := make([]byte, 12)
		_, _ = rnd.Read(buf)
		hashes = append(hashes, new(big.Int).SetBytes(buf))
	}
	return hashes
}

// randCoefficients generates random-ish coefficients from a seed so that
// the test is reproducible but exercises many different values.
func randCoefficients(seed int64, threshold int) []*big.Int {
	rbjj := group.ScalarField()
	rnd := mathrand.New(mathrand.NewSource(seed))
	coeffs := make([]*big.Int, threshold)
	for k := range threshold {
		b := make([]byte, 32)
		_, _ = rnd.Read(b)
		v := new(big.Int).SetBytes(b)
		v.Mod(v, rbjj)
		coeffs[k] = v
	}
	return coeffs
}

// randCryptoCoefficients generates cryptographically random r_bjj-scale
// coefficients to stress the AddModSubgroupOrder path.
func randCryptoCoefficients(threshold int) ([]*big.Int, error) {
	rbjj := group.ScalarField()
	coeffs := make([]*big.Int, threshold)
	for k := range threshold {
		v, err := rand.Int(rand.Reader, rbjj)
		if err != nil {
			return nil, err
		}
		coeffs[k] = v
	}
	return coeffs, nil
}

// TestContributionCircuitStress runs repeated prove+verify cycles with varied
// inputs to detect non-deterministic prover failures.
//
// It exercises:
//   - All committee sizes that fit in the current MaxN (step 4)
//   - Multiple contributor indexes per size
//   - Three coefficient regimes: benchmark-style small, pseudo-random r_bjj-scale,
//     and cryptographically random r_bjj-scale
//   - Multiple roundHash values (small, mid, max, random 12-byte IDs)
//   - Actual Anvil actor BJJ keys (same derivation as the integration tests)
func TestContributionCircuitStress(t *testing.T) {
	if MaxRecipients < 4 {
		t.Skip("MaxN too small")
	}

	c := qt.New(t)

	// Pre-derive the actual actor BJJ keys used by the integration tests.
	maxActors := MaxRecipients
	if maxActors > len(anvilPrivateKeys) {
		maxActors = len(anvilPrivateKeys)
	}
	actorKeys := make([]types.NodeKey, maxActors)
	for i := range maxActors {
		actorKeys[i] = actorBJJKey(anvilPrivateKeys[i])
	}

	// Load the circuit runtime once (reused across all subtests).
	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &ContributionCircuit{})
	c.Assert(err, qt.IsNil)

	roundHashes := stressRoundHashes(15) // 7 fixed + 15 random = 22 total

	// Committee sizes to test. Default is n=4 and n=8 (fast enough for CI).
	// Set STRESS_ALL_N=1 to test all sizes up to MaxN (takes 4+ hours at MaxN=32).
	var sizes []int
	if os.Getenv("STRESS_ALL_N") == "1" {
		for n := 4; n <= MaxRecipients; n += 4 {
			sizes = append(sizes, n)
		}
	} else {
		sizes = []int{4, 8}
		if MaxRecipients < 8 {
			sizes = []int{4}
		}
	}

	type failureInfo struct {
		n              int
		threshold      int
		contributorIdx int
		roundHash      *big.Int
		coefficients   []*big.Int
		err            error
	}
	var failures []failureInfo

	prove := func(n, threshold, contributorIdx int, roundHash *big.Int, coefficients []*big.Int) error {
		recipientIndexes := make([]uint16, n)
		recipientKeys := make([]types.NodeKey, n)
		nonces := make([]*big.Int, n)
		for j := range n {
			recipientIndexes[j] = uint16(j + 1)
			recipientKeys[j] = actorKeys[j]
			nonces[j] = big.NewInt(int64(1000 + j + 1))
		}
		assignment := Assignment{
			RoundHash:        roundHash,
			Threshold:        uint16(threshold),
			CommitteeSize:    uint16(n),
			ContributorIndex: uint16(contributorIdx),
			Coefficients:     coefficients,
			RecipientIndexes: recipientIndexes,
			RecipientKeys:    recipientKeys,
			EncryptionNonces: nonces,
		}
		witness, publicInputs, err := BuildWitness(assignment)
		if err != nil {
			return fmt.Errorf("BuildWitness: %w", err)
		}
		proof, err := runtime.ProveAndVerify(witness)
		if err != nil {
			return fmt.Errorf("ProveAndVerify: %w", err)
		}
		if err := runtime.Verify(proof, publicInputs.PublicWitness()); err != nil {
			return fmt.Errorf("Verify(publicWitness): %w", err)
		}
		return nil
	}

	for _, n := range sizes {
		n := n
		threshold := (2*n + 2) / 3 // ceil(2n/3)

		t.Run(fmt.Sprintf("n%d", n), func(t *testing.T) {
			// ── Regime 1: benchmark-style small coefficients, all contributors,
			//             all roundHash values.
			t.Run("small_coeffs", func(t *testing.T) {
				for _, rh := range roundHashes {
					for i := range n {
						coefficients := make([]*big.Int, threshold)
						for k := range threshold {
							coefficients[k] = big.NewInt(int64((i+1)*10 + k + 1))
						}
						if err := prove(n, threshold, i+1, rh, coefficients); err != nil {
							info := failureInfo{n: n, threshold: threshold, contributorIdx: i + 1, roundHash: rh, coefficients: coefficients, err: err}
							failures = append(failures, info)
							t.Errorf("FAIL n=%d t=%d contributor=%d roundHash=%x\n  coefficients=%v\n  error: %v",
								n, threshold, i+1, rh.Bytes(), coefficients, err)
						}
					}
				}
			})

			// ── Regime 2: pseudo-random r_bjj-scale coefficients (deterministic
			//             seed so the test is reproducible).
			t.Run("random_coeffs", func(t *testing.T) {
				for seedOffset, rh := range roundHashes[:5] { // 5 roundHashes × all contributors
					for i := range n {
						seed := int64(n*1000 + (i+1)*100 + seedOffset)
						coefficients := randCoefficients(seed, threshold)
						if err := prove(n, threshold, i+1, rh, coefficients); err != nil {
							info := failureInfo{n: n, threshold: threshold, contributorIdx: i + 1, roundHash: rh, coefficients: coefficients, err: err}
							failures = append(failures, info)
							t.Errorf("FAIL n=%d t=%d contributor=%d roundHash=%x\n  coefficients=%v\n  error: %v",
								n, threshold, i+1, rh.Bytes(), coefficients, err)
						}
					}
				}
			})

			// ── Regime 3: cryptographically random r_bjj-scale coefficients
			//             (exercises the AddModSubgroupOrder carry path).
			//             Run 3 iterations per contributor × 3 random roundHashes.
			t.Run("crypto_random_coeffs", func(t *testing.T) {
				for iter := 0; iter < 3; iter++ {
					rh, _ := rand.Int(rand.Reader, new(big.Int).SetBytes([]byte{
						0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
						0xff, 0xff, 0xff, 0xff,
					}))
					for i := range n {
						coefficients, err := randCryptoCoefficients(threshold)
						c.Assert(err, qt.IsNil)
						if err := prove(n, threshold, i+1, rh, coefficients); err != nil {
							info := failureInfo{n: n, threshold: threshold, contributorIdx: i + 1, roundHash: rh, coefficients: coefficients, err: err}
							failures = append(failures, info)
							t.Errorf("FAIL n=%d t=%d contributor=%d roundHash=%x\n  coefficients=%v\n  error: %v",
								n, threshold, i+1, rh.Bytes(), coefficients, err)
						}
					}
				}
			})
		})
	}

	if len(failures) > 0 {
		t.Logf("\n=== FAILURE SUMMARY (%d failures) ===", len(failures))
		for _, f := range failures {
			t.Logf("n=%d t=%d contributor=%d roundHash=%s err=%v",
				f.n, f.threshold, f.contributorIdx, f.roundHash.String(), f.err)
			for k, c := range f.coefficients {
				t.Logf("  coeff[%d] = %s", k, c.String())
			}
		}
	}
}

// TestContributionCircuitStressRepeat repeats the exact failing scenario
// (n=12, t=8, contributor 9, benchmark coefficients) many times to catch
// intermittent prover failures.
func TestContributionCircuitStressRepeat(t *testing.T) {
	const n, threshold, contributorI = 12, 8, 9
	const iterations = 30

	if MaxRecipients < n {
		t.Skipf("MaxN=%d < n=%d", MaxRecipients, n)
	}

	c := qt.New(t)

	// Real actor keys.
	recipientKeys := make([]types.NodeKey, n)
	for j := range n {
		recipientKeys[j] = actorBJJKey(anvilPrivateKeys[j])
	}
	recipientIndexes := make([]uint16, n)
	nonces := make([]*big.Int, n)
	for j := range n {
		recipientIndexes[j] = uint16(j + 1)
		nonces[j] = big.NewInt(int64(1000 + j + 1))
	}
	coefficients := make([]*big.Int, threshold)
	for k := range threshold {
		coefficients[k] = big.NewInt(int64(contributorI*10 + k + 1))
	}

	runtime, err := Artifacts.LoadOrSetupForCircuit(context.Background(), &ContributionCircuit{})
	c.Assert(err, qt.IsNil)

	// Vary the roundHash across iterations to hit different mask values.
	for iter := range iterations {
		var roundHash *big.Int
		switch {
		case iter < 5:
			roundHash = big.NewInt(int64(iter + 1))
		case iter < 10:
			roundHash = new(big.Int).SetBytes([]byte{
				0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe,
				0x00, 0x00, 0x00, byte(iter),
			})
		default:
			var err error
			roundHash, err = rand.Int(rand.Reader, new(big.Int).SetBytes([]byte{
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff,
			}))
			c.Assert(err, qt.IsNil)
		}

		assignment := Assignment{
			RoundHash:        roundHash,
			Threshold:        uint16(threshold),
			CommitteeSize:    uint16(n),
			ContributorIndex: uint16(contributorI),
			Coefficients:     coefficients,
			RecipientIndexes: recipientIndexes,
			RecipientKeys:    recipientKeys,
			EncryptionNonces: nonces,
		}
		witness, publicInputs, buildErr := BuildWitness(assignment)
		if buildErr != nil {
			t.Errorf("iter=%d roundHash=%s BuildWitness: %v", iter, roundHash, buildErr)
			continue
		}
		proof, proveErr := runtime.ProveAndVerify(witness)
		if proveErr != nil {
			t.Errorf("iter=%d roundHash=%s ProveAndVerify: %v", iter, roundHash, proveErr)
			t.Logf("  coefficients: %v", coefficients)
			continue
		}
		if verifyErr := runtime.Verify(proof, publicInputs.PublicWitness()); verifyErr != nil {
			t.Errorf("iter=%d roundHash=%s Verify(publicWitness): %v", iter, roundHash, verifyErr)
		}
	}
}
