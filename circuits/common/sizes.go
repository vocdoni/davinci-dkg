package common

// MaxN is the single source of truth for the compile-time array bound used
// across every DKG circuit. It governs MaxCoefficients, MaxRecipients,
// MaxParticipants and MaxShares — they are all aliases of the same value.
//
// To change the maximum committee size, edit this one constant and then
// recompile every circuit (which regenerates proving keys, verifying keys
// and the corresponding Solidity verifier wrappers). The Solidity contract
// reads `MAX_N` from `solidity/src/DKGManager.sol`, which must be set to the
// same number; the test `TestSolidityMaxNMatchesGoMaxN` enforces this.
const MaxN = 32
