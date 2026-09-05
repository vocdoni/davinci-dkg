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
// carries MaxK polynomials and each application claims one key. Mirror of
// `MAX_K` in `solidity/src/libraries/Sizes.sol`. See docs/pool-keys.md.
const MaxK = 8

// MerkleDepth is log2(MaxN): the depth of the keccak Merkle tree over the
// per-member share commitments of one pool key that `activatePoolKey`
// stores and `submitPartialDecryption` proves membership in. Mirror of
// `MERKLE_DEPTH` in Sizes.sol.
const MerkleDepth = 5
