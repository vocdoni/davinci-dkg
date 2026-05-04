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
