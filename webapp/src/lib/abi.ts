// ABI and type re-exports from the DKG SDK.
// The SDK owns the canonical ABI definitions; the webapp imports them here.
// Display helpers (color/label) that map SDK values to Chakra UI tokens stay local.

export {
  dkgManagerAbi,
  dkgRegistryAbi,
  RoundStatus,
  NodeStatus,
  roundStatusLabel,
} from '@vocdoni/davinci-dkg-sdk';

export const nodeStatusLabel = (status: number): string => {
  switch (status) {
    case 0: return 'None';
    case 1: return 'Active';
    case 2: return 'Inactive';
    default: return `Unknown (${status})`;
  }
};

export const nodeStatusColor = (status: number): string => {
  switch (status) {
    case 1: return 'green';
    case 2: return 'orange';
    default: return 'gray';
  }
};

export const roundStatusColor = (status: number): string => {
  switch (status) {
    case 1: return 'yellow';
    case 2: return 'blue';
    case 3: return 'cyan';
    case 4: return 'red';
    case 5: return 'green';
    default: return 'gray';
  }
};
