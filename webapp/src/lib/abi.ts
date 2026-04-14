// Minimal ABI subsets for the DKG explorer. Only the read-only views and the
// events the explorer actually consumes are listed here.

export const dkgManagerAbi = [
  // ── views ────────────────────────────────────────────────────────────────
  {
    type: 'function',
    name: 'roundNonce',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint64' }],
  },
  {
    type: 'function',
    name: 'CHAIN_ID',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint32' }],
  },
  {
    type: 'function',
    name: 'ROUND_PREFIX',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint32' }],
  },
  {
    type: 'function',
    name: 'REGISTRY',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
  },
  {
    type: 'function',
    name: 'getRound',
    stateMutability: 'view',
    inputs: [{ name: 'roundId', type: 'bytes12' }],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'organizer', type: 'address' },
          {
            name: 'policy',
            type: 'tuple',
            components: [
              { name: 'threshold', type: 'uint16' },
              { name: 'committeeSize', type: 'uint16' },
              { name: 'minValidContributions', type: 'uint16' },
              { name: 'lotteryAlphaBps', type: 'uint16' },
              { name: 'seedDelay', type: 'uint16' },
              { name: 'registrationDeadlineBlock', type: 'uint64' },
              { name: 'contributionDeadlineBlock', type: 'uint64' },
              { name: 'disclosureAllowed', type: 'bool' },
            ],
          },
          { name: 'status', type: 'uint8' },
          { name: 'nonce', type: 'uint64' },
          { name: 'seedBlock', type: 'uint64' },
          { name: 'seed', type: 'bytes32' },
          { name: 'lotteryThreshold', type: 'uint256' },
          { name: 'claimedCount', type: 'uint16' },
          { name: 'contributionCount', type: 'uint16' },
          { name: 'partialDecryptionCount', type: 'uint16' },
          { name: 'revealedShareCount', type: 'uint16' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'selectedParticipants',
    stateMutability: 'view',
    inputs: [{ name: 'roundId', type: 'bytes12' }],
    outputs: [{ name: '', type: 'address[]' }],
  },
  {
    type: 'function',
    name: 'getContribution',
    stateMutability: 'view',
    inputs: [
      { name: 'roundId', type: 'bytes12' },
      { name: 'contributor', type: 'address' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'contributor', type: 'address' },
          { name: 'contributorIndex', type: 'uint16' },
          { name: 'commitmentsHash', type: 'bytes32' },
          { name: 'encryptedSharesHash', type: 'bytes32' },
          { name: 'commitmentVectorDigest', type: 'bytes32' },
          { name: 'accepted', type: 'bool' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'getPartialDecryption',
    stateMutability: 'view',
    inputs: [
      { name: 'roundId', type: 'bytes12' },
      { name: 'participant', type: 'address' },
      { name: 'ciphertextIndex', type: 'uint16' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'participant', type: 'address' },
          { name: 'participantIndex', type: 'uint16' },
          { name: 'ciphertextIndex', type: 'uint16' },
          { name: 'deltaHash', type: 'bytes32' },
          {
            name: 'delta',
            type: 'tuple',
            components: [
              { name: 'x', type: 'uint256' },
              { name: 'y', type: 'uint256' },
            ],
          },
          { name: 'accepted', type: 'bool' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'getCombinedDecryption',
    stateMutability: 'view',
    inputs: [
      { name: 'roundId', type: 'bytes12' },
      { name: 'ciphertextIndex', type: 'uint16' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'ciphertextIndex', type: 'uint16' },
          { name: 'combineHash', type: 'bytes32' },
          { name: 'plaintextHash', type: 'bytes32' },
          { name: 'completed', type: 'bool' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'getShareCommitmentHash',
    stateMutability: 'view',
    inputs: [
      { name: 'roundId', type: 'bytes12' },
      { name: 'participantIndex', type: 'uint16' },
    ],
    outputs: [{ name: '', type: 'bytes32' }],
  },

  // ── events ────────────────────────────────────────────────────────────────
  {
    type: 'event',
    name: 'RoundCreated',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'organizer', type: 'address', indexed: true },
      { name: 'seedBlock', type: 'uint64', indexed: false },
      { name: 'lotteryThreshold', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'SeedResolved',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'seed', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'SlotClaimed',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'claimer', type: 'address', indexed: true },
      { name: 'slot', type: 'uint16', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'RegistrationClosed',
    inputs: [{ name: 'roundId', type: 'bytes12', indexed: true }],
  },
  {
    type: 'event',
    name: 'ContributionSubmitted',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'contributor', type: 'address', indexed: true },
      { name: 'contributorIndex', type: 'uint16', indexed: false },
      { name: 'commitmentsHash', type: 'bytes32', indexed: false },
      { name: 'encryptedSharesHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'RoundFinalized',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'aggregateCommitmentsHash', type: 'bytes32', indexed: false },
      { name: 'collectivePublicKeyHash', type: 'bytes32', indexed: false },
      { name: 'shareCommitmentHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'PartialDecryptionSubmitted',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'participant', type: 'address', indexed: true },
      { name: 'participantIndex', type: 'uint16', indexed: false },
      { name: 'ciphertextIndex', type: 'uint16', indexed: false },
      { name: 'deltaHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'DecryptionCombined',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'ciphertextIndex', type: 'uint16', indexed: true },
      { name: 'combineHash', type: 'bytes32', indexed: false },
      { name: 'plaintextHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'RevealedShareSubmitted',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'participant', type: 'address', indexed: true },
      { name: 'participantIndex', type: 'uint16', indexed: false },
      { name: 'shareHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'SecretReconstructed',
    inputs: [
      { name: 'roundId', type: 'bytes12', indexed: true },
      { name: 'disclosureHash', type: 'bytes32', indexed: false },
      { name: 'reconstructedSecretHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'RoundEvicted',
    inputs: [{ name: 'roundId', type: 'bytes12', indexed: true }],
  },
  {
    type: 'event',
    name: 'RoundAborted',
    inputs: [{ name: 'roundId', type: 'bytes12', indexed: true }],
  },
] as const;

export const dkgRegistryAbi = [
  {
    type: 'function',
    name: 'nodeCount',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint64' }],
  },
  {
    type: 'function',
    name: 'activeCount',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint64' }],
  },
  {
    type: 'function',
    name: 'INACTIVITY_WINDOW',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint64' }],
  },
  {
    type: 'function',
    name: 'isActive',
    stateMutability: 'view',
    inputs: [{ name: 'operator', type: 'address' }],
    outputs: [{ name: '', type: 'bool' }],
  },
  {
    type: 'function',
    name: 'getNode',
    stateMutability: 'view',
    inputs: [{ name: 'operator', type: 'address' }],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'operator', type: 'address' },
          { name: 'pubX', type: 'uint256' },
          { name: 'pubY', type: 'uint256' },
          { name: 'status', type: 'uint8' },
          { name: 'lastActiveBlock', type: 'uint64' },
        ],
      },
    ],
  },
  {
    type: 'event',
    name: 'NodeRegistered',
    inputs: [
      { name: 'operator', type: 'address', indexed: true },
      { name: 'pubX', type: 'uint256', indexed: false },
      { name: 'pubY', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'NodeUpdated',
    inputs: [
      { name: 'operator', type: 'address', indexed: true },
      { name: 'pubX', type: 'uint256', indexed: false },
      { name: 'pubY', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'NodeMarkedActive',
    inputs: [
      { name: 'operator', type: 'address', indexed: true },
      { name: 'atBlock', type: 'uint64', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'NodeReaped',
    inputs: [
      { name: 'operator', type: 'address', indexed: true },
      { name: 'lastActiveBlock', type: 'uint64', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'NodeReactivated',
    inputs: [{ name: 'operator', type: 'address', indexed: true }],
  },
] as const;

export const NodeStatus = {
  None: 0,
  Active: 1,
  Inactive: 2,
} as const;

export const nodeStatusLabel = (status: number): string => {
  switch (status) {
    case NodeStatus.None:
      return 'None';
    case NodeStatus.Active:
      return 'Active';
    case NodeStatus.Inactive:
      return 'Inactive';
    default:
      return `Unknown (${status})`;
  }
};

export const nodeStatusColor = (status: number): string => {
  switch (status) {
    case NodeStatus.Active:
      return 'green';
    case NodeStatus.Inactive:
      return 'orange';
    default:
      return 'gray';
  }
};

// Round status enum values mirror DKGTypes.RoundStatus.
export const RoundStatus = {
  None: 0,
  Registration: 1,
  Contribution: 2,
  Finalized: 3,
  Aborted: 4,
  Completed: 5,
} as const;

export const roundStatusLabel = (status: number): string => {
  switch (status) {
    case 0:
      return 'None';
    case 1:
      return 'Registration';
    case 2:
      return 'Contribution';
    case 3:
      return 'Finalized';
    case 4:
      return 'Aborted';
    case 5:
      return 'Completed';
    default:
      return `Unknown (${status})`;
  }
};

export const roundStatusColor = (status: number): string => {
  switch (status) {
    case 1:
      return 'yellow';
    case 2:
      return 'blue';
    case 3:
      return 'cyan';
    case 4:
      return 'red';
    case 5:
      return 'green';
    default:
      return 'gray';
  }
};
