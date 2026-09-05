package common

// MaxN is the single source of truth for the compile-time array bound used
// across every DKG circuit. It governs MaxCoefficients, MaxRecipients,
// MaxParticipants and MaxShares — they are all aliases of the same value.
//
// To change the maximum committee size, edit this one constant and the
// matching `MAX_N` in `solidity/src/libraries/Sizes.sol`, then run
// `make circuits` (recompiles every circuit, regenerates proving / verifying
// keys, the Solidity verifier wrappers, and the Go bindings).
const MaxN = 32

// MaxK is the number of pool keys every epoch deals: each contribution
// carries MaxK polynomials, finalization derives all MaxK keys at once and
// each application claims one key. Mirror of `MAX_K` in
// `solidity/src/libraries/Sizes.sol`. See docs/pool-keys-v4.md.
const MaxK = 16

// MerkleDepth is log2(MaxN): the depth of the keccak Merkle tree over the
// per-member share commitments of one pool key that `finalizeEpoch` stores
// and `submitPartialDecryption` proves membership in. Mirror of
// `MERKLE_DEPTH` in Sizes.sol.
const MerkleDepth = 5

// FinalizeTranscriptWords is the fixed word count L_F of the finalization
// calldata transcript (docs/pool-keys-v4.md §7): the accepted dealers'
// indexes and contribution hashes (2·MaxN words), then for every key its
// pool key P_j and the share commitment D_j,i of every committee slot
// (2 + 2·MaxN words per key). 1,120 at MaxN = 32, MaxK = 16. Mirror of
// `FINALIZE_TRANSCRIPT_WORDS` in Sizes.sol.
const FinalizeTranscriptWords = 2*MaxN + MaxK*(2+2*MaxN)

// CompactContributionWords returns L_C(t, n) = MaxK·(2t+n) + 5n, the word
// count of the compact contribution transcript (docs/pool-keys-v4.md §3) for
// threshold t and committee size n: 2t commitment coordinates per key, n
// recipient indexes, 2n public-key and 2n ephemeral coordinates, and n
// masked shares per key. No padding travels in calldata, so the length is a
// function of the epoch's public policy, not of the circuit bounds. The
// contract computes the same value from the epoch it verifies against.
func CompactContributionWords(threshold, committeeSize int) int {
	return MaxK*(2*threshold+committeeSize) + 5*committeeSize
}
