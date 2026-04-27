import { sepolia, foundry, type Chain } from 'wagmi/chains'

// The chains the wallet picker will offer. Wagmi will reject any other
// chain at connect time, so keep this list strictly to deployments we
// actually target. Add new entries when DKGManager is deployed somewhere new.
export const supportedChains: readonly [Chain, ...Chain[]] = [sepolia, foundry]
