// Complete ABI for the DKG Manager and Registry contracts.
// Includes all functions (read and write), events, and custom errors.

export const dkgManagerAbi = [
  // ── write functions ───────────────────────────────────────────────────────
  {
    type: 'function',
    name: 'createEpoch',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'threshold', type: 'uint16' },
      { name: 'committeeSize', type: 'uint16' },
      { name: 'minValidContributions', type: 'uint16' },
      { name: 'lotteryAlphaBps', type: 'uint16' },
    ],
    outputs: [{ name: '', type: 'bytes12' }],
  },
  {
    type: 'function',
    name: 'nextEpochStartBlock',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint64' }],
  },
  {
    type: 'function',
    name: 'epochDurationBlocks',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
  },
  {
    type: 'function',
    name: 'lastEpochStartBlock',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint64' }],
  },
  {
    type: 'function',
    name: 'EPOCH_DURATION_BLOCKS',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint256' }],
  },
  // Deploy-time bounds enforced by createEpoch (see DKGManager constructor).
  {
    type: 'function',
    name: 'MIN_THRESHOLD',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint16' }],
  },
  {
    type: 'function',
    name: 'MIN_COMMITTEE_SIZE',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint16' }],
  },
  {
    type: 'function',
    name: 'MAX_LOTTERY_ALPHA_BPS',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'uint16' }],
  },
  {
    type: 'function',
    name: 'submitCiphertext',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'c1x', type: 'uint256' },
      { name: 'c1y', type: 'uint256' },
      { name: 'c2x', type: 'uint256' },
      { name: 'c2y', type: 'uint256' },
    ],
    outputs: [{ name: 'ciphertextIndex', type: 'uint16' }],
  },
  {
    type: 'function',
    name: 'claimSlot',
    stateMutability: 'nonpayable',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
    outputs: [],
  },
  // `extendRegistration` was removed in the auto-cadence refactor.
  {
    type: 'function',
    name: 'submitContribution',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'contributorIndex', type: 'uint16' },
      { name: 'commitmentsHash', type: 'bytes32' },
      { name: 'encryptedSharesHash', type: 'bytes32' },
      { name: 'transcript', type: 'bytes' },
      { name: 'proof', type: 'bytes' },
      { name: 'input', type: 'bytes' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'finalizeEpoch',
    stateMutability: 'nonpayable',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
    outputs: [],
  },
  {
    type: 'function',
    name: 'activatePoolKey',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'keyIndex', type: 'uint8' },
      // Poseidon digest of the masked transcript words (public input 5); the
      // BRLC challenge is anchored on keccak(transcriptDigest ‖ keccak(transcript)).
      { name: 'transcriptDigest', type: 'bytes32' },
      { name: 'transcript', type: 'bytes' },
      { name: 'proof', type: 'bytes' },
      { name: 'input', type: 'bytes' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'claimPoolKey',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
    ],
    outputs: [{ name: 'keyIndex', type: 'uint8' }],
  },
  {
    type: 'function',
    name: 'submitPartialDecryption',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'participantIndex', type: 'uint16' },
      { name: 'ciphertextIndex', type: 'uint16' },
      { name: 'c1x', type: 'uint256' },
      { name: 'c1y', type: 'uint256' },
      { name: 'c2x', type: 'uint256' },
      { name: 'c2y', type: 'uint256' },
      { name: 'deltaHash', type: 'bytes32' },
      { name: 'proof', type: 'bytes' },
      { name: 'input', type: 'bytes' },
      // Merkle proof (bottom-up siblings, MERKLE_DEPTH=5) that this
      // participant's share commitment is in poolShareRoots[epochId][poolIndex].
      { name: 'shareProof', type: 'bytes32[]' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'combineDecryption',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'ciphertextIndex', type: 'uint16' },
      { name: 'combineHash', type: 'bytes32' },
      { name: 'plaintext', type: 'uint256' },
      { name: 'transcript', type: 'bytes' },
      { name: 'proof', type: 'bytes' },
      { name: 'input', type: 'bytes' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'abortEpoch',
    stateMutability: 'nonpayable',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
    outputs: [],
  },

  // ── view functions ────────────────────────────────────────────────────────
  {
    type: 'function',
    name: 'epochNonce',
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
    name: 'EPOCH_PREFIX',
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
    name: 'appManager',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
  },
  {
    type: 'function',
    name: 'getEpoch',
    stateMutability: 'view',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
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
              { name: 'committeeSelectionDeadlineBlock', type: 'uint64' },
              { name: 'keyAssemblyDeadlineBlock', type: 'uint64' },
              { name: 'liveNotBeforeBlock', type: 'uint64' },
            ],
          },
          { name: 'status', type: 'uint8' },
          { name: 'nonce', type: 'uint64' },
          { name: 'startBlock', type: 'uint64' },
          { name: 'seedBlock', type: 'uint64' },
          { name: 'seed', type: 'bytes32' },
          { name: 'lotteryThreshold', type: 'uint256' },
          { name: 'claimedCount', type: 'uint16' },
          { name: 'contributionCount', type: 'uint16' },
          { name: 'partialDecryptionCount', type: 'uint16' },
          { name: 'ciphertextCount', type: 'uint16' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'selectedParticipants',
    stateMutability: 'view',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
    outputs: [{ name: '', type: 'address[]' }],
  },
  {
    type: 'function',
    name: 'getContribution',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
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
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'participantIndex', type: 'uint16' },
      { name: 'ciphertextIndex', type: 'uint16' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'participantIndex', type: 'uint16' },
          { name: 'ciphertextIndex', type: 'uint16' },
          { name: 'deltaHash', type: 'bytes32' },
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
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'ciphertextIndex', type: 'uint16' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'ciphertextIndex', type: 'uint16' },
          { name: 'completed', type: 'bool' },
          { name: 'plaintext', type: 'uint256' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'getCiphertextHash',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'ciphertextIndex', type: 'uint16' },
    ],
    outputs: [{ name: '', type: 'bytes32' }],
  },
  {
    type: 'function',
    name: 'getPlaintext',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'ciphertextIndex', type: 'uint16' },
    ],
    outputs: [{ name: '', type: 'uint256' }],
  },
  {
    type: 'function',
    name: 'ciphertextCount',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
    ],
    outputs: [{ name: '', type: 'uint16' }],
  },
  {
    type: 'function',
    name: 'getContributionVerifierVKeyHash',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'bytes32' }],
  },
  {
    type: 'function',
    name: 'getPartialDecryptVerifierVKeyHash',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'bytes32' }],
  },
  {
    type: 'function',
    name: 'getPoolKeyVerifierVKeyHash',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'bytes32' }],
  },
  {
    type: 'function',
    name: 'getDecryptCombineVerifierVKeyHash',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'bytes32' }],
  },
  {
    type: 'function',
    name: 'getPoolKey',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'keyIndex', type: 'uint8' },
    ],
    outputs: [
      { name: '', type: 'uint256' },
      { name: '', type: 'uint256' },
    ],
  },
  {
    type: 'function',
    name: 'getPoolStatus',
    stateMutability: 'view',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
    outputs: [
      { name: 'nextIndex', type: 'uint8' },
      { name: 'activated', type: 'uint8' },
    ],
  },
  {
    type: 'function',
    name: 'getPoolShareRoot',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'keyIndex', type: 'uint8' },
    ],
    outputs: [{ name: '', type: 'bytes32' }],
  },
  {
    type: 'function',
    name: 'getAppPoolIndex',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
    ],
    outputs: [{ name: '', type: 'uint8' }],
  },

  // ── events ────────────────────────────────────────────────────────────────
  {
    type: 'event',
    name: 'EpochCreated',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'organizer', type: 'address', indexed: true },
      { name: 'startBlock', type: 'uint64', indexed: false },
      { name: 'seedBlock', type: 'uint64', indexed: false },
      { name: 'lotteryThreshold', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'SeedResolved',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'seed', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'SlotClaimed',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'claimer', type: 'address', indexed: true },
      { name: 'slot', type: 'uint16', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'CommitteeFilled',
    inputs: [{ name: 'epochId', type: 'bytes12', indexed: true }],
  },
  {
    type: 'event',
    name: 'ContributionSubmitted',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'contributor', type: 'address', indexed: true },
      { name: 'contributorIndex', type: 'uint16', indexed: false },
      { name: 'commitmentsHash', type: 'bytes32', indexed: false },
      { name: 'encryptedSharesHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'EpochLive',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'contributionCount', type: 'uint16', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'PoolKeyActivated',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'keyIndex', type: 'uint8', indexed: true },
      { name: 'x', type: 'uint256', indexed: false },
      { name: 'y', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'PoolKeyClaimed',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'aid', type: 'bytes32', indexed: true },
      { name: 'keyIndex', type: 'uint8', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'PartialDecryptionSubmitted',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'aid', type: 'bytes32', indexed: true },
      { name: 'participant', type: 'address', indexed: true },
      { name: 'participantIndex', type: 'uint16', indexed: false },
      { name: 'ciphertextIndex', type: 'uint16', indexed: false },
      { name: 'deltaX', type: 'uint256', indexed: false },
      { name: 'deltaY', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'CiphertextSubmitted',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'aid', type: 'bytes32', indexed: true },
      { name: 'ciphertextIndex', type: 'uint16', indexed: true },
      { name: 'submitter', type: 'address', indexed: false },
      { name: 'c1x', type: 'uint256', indexed: false },
      { name: 'c1y', type: 'uint256', indexed: false },
      { name: 'c2x', type: 'uint256', indexed: false },
      { name: 'c2y', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'DecryptionCombined',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'aid', type: 'bytes32', indexed: true },
      { name: 'ciphertextIndex', type: 'uint16', indexed: true },
      { name: 'combineHash', type: 'bytes32', indexed: false },
      { name: 'plaintext', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'EpochAborted',
    inputs: [{ name: 'epochId', type: 'bytes12', indexed: true }],
  },

  // ── errors ────────────────────────────────────────────────────────────────
  { type: 'error', name: 'InvalidPolicy', inputs: [] },
  { type: 'error', name: 'InvalidChainId', inputs: [] },
  { type: 'error', name: 'InvalidAddress', inputs: [] },
  { type: 'error', name: 'InvalidEpoch', inputs: [] },
  { type: 'error', name: 'InvalidPhase', inputs: [] },
  { type: 'error', name: 'NotEligible', inputs: [] },
  { type: 'error', name: 'AlreadyClaimed', inputs: [] },
  { type: 'error', name: 'SlotsFull', inputs: [] },
  { type: 'error', name: 'SeedNotReady', inputs: [] },
  { type: 'error', name: 'SeedExpired', inputs: [] },
  { type: 'error', name: 'NotRegistered', inputs: [] },
  // The operator registered after the epoch was created; only the registry
  // snapshot taken at createEpoch enters the lottery.
  { type: 'error', name: 'NotInSnapshot', inputs: [] },
  { type: 'error', name: 'NotSelectedParticipant', inputs: [] },
  { type: 'error', name: 'AlreadyContributed', inputs: [] },
  { type: 'error', name: 'AlreadyLive', inputs: [] },
  { type: 'error', name: 'AlreadyPartiallyDecrypted', inputs: [] },
  { type: 'error', name: 'InvalidCommitteeSize', inputs: [] },
  { type: 'error', name: 'InvalidContribution', inputs: [] },
  { type: 'error', name: 'InvalidFinalization', inputs: [] },
  { type: 'error', name: 'InvalidPartialDecryption', inputs: [] },
  { type: 'error', name: 'InsufficientContributions', inputs: [] },
  { type: 'error', name: 'InvalidVerifier', inputs: [] },
  { type: 'error', name: 'Unauthorized', inputs: [] },
  { type: 'error', name: 'AlreadyCombined', inputs: [] },
  { type: 'error', name: 'InvalidCombinedDecryption', inputs: [] },
  { type: 'error', name: 'InsufficientPartialDecryptions', inputs: [] },
  { type: 'error', name: 'InvalidProofInput', inputs: [] },
  { type: 'error', name: 'DecryptionLimitReached', inputs: [] },
  { type: 'error', name: 'CiphertextAlreadySubmitted', inputs: [] },
  { type: 'error', name: 'CiphertextNotSubmitted', inputs: [] },
  { type: 'error', name: 'InvalidCiphertext', inputs: [] },
  { type: 'error', name: 'AppManagerAlreadySet', inputs: [] },
  { type: 'error', name: 'AppManagerNotSet', inputs: [] },
  { type: 'error', name: 'PoolExhausted', inputs: [] },
  { type: 'error', name: 'PoolKeyAlreadyActive', inputs: [] },
  { type: 'error', name: 'PoolKeyNotActive', inputs: [] },
  // BRLC: a transcript word >= the BN254 scalar field on any proof-carrying
  // call (submitContribution, activatePoolKey, combineDecryption).
  { type: 'error', name: 'TranscriptWordNotInField', inputs: [] },
  // Raised by DKGAppManager and bubbled through submitCiphertext,
  // submitPartialDecryption and combineDecryption; listed here so viem
  // decodes them on manager calls.
  { type: 'error', name: 'InvalidApplication', inputs: [] },
  { type: 'error', name: 'NotOwner', inputs: [] },
  { type: 'error', name: 'DecryptionNotYetAllowed', inputs: [] },
  { type: 'error', name: 'DecryptionExpired', inputs: [] },
  { type: 'error', name: 'DecryptionNotOpen', inputs: [] },
  { type: 'error', name: 'DecryptionClosed', inputs: [] },
  { type: 'error', name: 'OrganizerSecretNotRevealed', inputs: [] },
] as const;

export const dkgRegistryAbi = [
  // ── write functions ───────────────────────────────────────────────────────
  {
    type: 'function',
    name: 'registerKey',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'pubX', type: 'uint256' },
      { name: 'pubY', type: 'uint256' },
      { name: 'schnorrAx', type: 'uint256' },
      { name: 'schnorrAy', type: 'uint256' },
      { name: 'schnorrZ', type: 'uint256' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'updateKey',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'pubX', type: 'uint256' },
      { name: 'pubY', type: 'uint256' },
      { name: 'schnorrAx', type: 'uint256' },
      { name: 'schnorrAy', type: 'uint256' },
      { name: 'schnorrZ', type: 'uint256' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'reactivate',
    stateMutability: 'nonpayable',
    inputs: [],
    outputs: [],
  },
  {
    type: 'function',
    name: 'heartbeat',
    stateMutability: 'nonpayable',
    inputs: [],
    outputs: [],
  },
  {
    type: 'function',
    name: 'reap',
    stateMutability: 'nonpayable',
    inputs: [{ name: 'operator', type: 'address' }],
    outputs: [],
  },
  {
    type: 'function',
    name: 'markActive',
    stateMutability: 'nonpayable',
    inputs: [{ name: 'operator', type: 'address' }],
    outputs: [],
  },

  // ── view functions ────────────────────────────────────────────────────────
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
    name: 'manager',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
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
          // Block of the first registerKey; the node only enters lotteries of
          // epochs created after it.
          { name: 'registeredAtBlock', type: 'uint64' },
        ],
      },
    ],
  },

  // ── events ────────────────────────────────────────────────────────────────
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
  {
    type: 'event',
    name: 'ManagerSet',
    inputs: [{ name: 'manager', type: 'address', indexed: true }],
  },

  // ── errors ────────────────────────────────────────────────────────────────
  { type: 'error', name: 'InvalidKey', inputs: [] },
  { type: 'error', name: 'AlreadyRegistered', inputs: [] },
  { type: 'error', name: 'NotRegistered', inputs: [] },
  { type: 'error', name: 'NotManager', inputs: [] },
  { type: 'error', name: 'ManagerAlreadySet', inputs: [] },
  { type: 'error', name: 'ManagerNotSet', inputs: [] },
  { type: 'error', name: 'NotActive', inputs: [] },
  { type: 'error', name: 'StillActive', inputs: [] },
  { type: 'error', name: 'NotInactive', inputs: [] },
  { type: 'error', name: 'InvalidAddress', inputs: [] },
  { type: 'error', name: 'Unauthorized', inputs: [] },
  { type: 'error', name: 'InvalidSchnorrProof', inputs: [] },
  { type: 'error', name: 'PointNotCanonical', inputs: [] },
  { type: 'error', name: 'PointNotOnCurve', inputs: [] },
  { type: 'error', name: 'PointIsIdentity', inputs: [] },
] as const;

export const dkgAppManagerAbi = [
  // ── write functions ───────────────────────────────────────────────────────
  {
    type: 'function',
    name: 'registerApplication',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      {
        name: 'policy',
        type: 'tuple',
        components: [
          // DKGTypes.AppMode: 0 = OrganizerLocked, 1 = Automatic.
          { name: 'mode', type: 'uint8' },
          { name: 'openSubmission', type: 'bool' },
          { name: 'submitters', type: 'address[]' },
          { name: 'maxCiphertexts', type: 'uint16' },
          { name: 'notBeforeBlock', type: 'uint64' },
          { name: 'notAfterBlock', type: 'uint64' },
          { name: 'decryptNotBefore', type: 'uint64' },
          { name: 'decryptNotAfter', type: 'uint64' },
        ],
      },
      // Automatic mode ignores these and stores the fixed identity (0, 1)
      // with a zero Schnorr proof.
      { name: 'pkOrgX', type: 'uint256' },
      { name: 'pkOrgY', type: 'uint256' },
      { name: 'schnorrAx', type: 'uint256' },
      { name: 'schnorrAy', type: 'uint256' },
      { name: 'schnorrZ', type: 'uint256' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'revealOrganizerSecret',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'organizerSecret', type: 'uint256' },
    ],
    outputs: [],
  },

  // ── view functions ────────────────────────────────────────────────────────
  {
    type: 'function',
    name: 'getApplication',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
    ],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
          { name: 'creator', type: 'address' },
          {
            name: 'organizerPK',
            type: 'tuple',
            components: [
              { name: 'x', type: 'uint256' },
              { name: 'y', type: 'uint256' },
            ],
          },
          { name: 'organizerSecret', type: 'uint256' },
          { name: 'poolIndex', type: 'uint8' },
          {
            name: 'policy',
            type: 'tuple',
            components: [
              { name: 'mode', type: 'uint8' },
              { name: 'openSubmission', type: 'bool' },
              { name: 'submitters', type: 'address[]' },
              { name: 'maxCiphertexts', type: 'uint16' },
              { name: 'notBeforeBlock', type: 'uint64' },
              { name: 'notAfterBlock', type: 'uint64' },
              { name: 'decryptNotBefore', type: 'uint64' },
              { name: 'decryptNotAfter', type: 'uint64' },
            ],
          },
          { name: 'createdAtBlock', type: 'uint64' },
          { name: 'exists', type: 'bool' },
        ],
      },
    ],
  },
  {
    type: 'function',
    name: 'getOrganizerPK',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
    ],
    outputs: [
      { name: '', type: 'uint256' },
      { name: '', type: 'uint256' },
    ],
  },
  // Reverts DecryptionNotOpen() before decryptNotBefore, DecryptionClosed()
  // past decryptNotAfter, OrganizerSecretNotRevealed() for an OrganizerLocked
  // application whose organizer has not revealed yet, and InvalidApplication()
  // for an unknown aid; returns nothing otherwise.
  {
    type: 'function',
    name: 'requireDecryptionOpen',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'requireCanSubmitCiphertext',
    stateMutability: 'view',
    inputs: [
      { name: 'epochId', type: 'bytes12' },
      { name: 'aid', type: 'bytes32' },
      { name: 'ciphertextIndex', type: 'uint16' },
      { name: 'sender', type: 'address' },
    ],
    outputs: [],
  },
  {
    type: 'function',
    name: 'getRegisteredAids',
    stateMutability: 'view',
    inputs: [{ name: 'epochId', type: 'bytes12' }],
    outputs: [{ name: '', type: 'bytes32[]' }],
  },
  {
    type: 'function',
    name: 'MANAGER',
    stateMutability: 'view',
    inputs: [],
    outputs: [{ name: '', type: 'address' }],
  },

  // ── events ────────────────────────────────────────────────────────────────
  {
    type: 'event',
    name: 'ApplicationRegistered',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'aid', type: 'bytes32', indexed: true },
      { name: 'creator', type: 'address', indexed: true },
      { name: 'organizerPKx', type: 'uint256', indexed: false },
      { name: 'organizerPKy', type: 'uint256', indexed: false },
      { name: 'mode', type: 'uint8', indexed: false },
      { name: 'poolIndex', type: 'uint8', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'OrganizerSecretRevealed',
    inputs: [
      { name: 'epochId', type: 'bytes12', indexed: true },
      { name: 'aid', type: 'bytes32', indexed: true },
      { name: 'organizerSecret', type: 'uint256', indexed: false },
    ],
  },

  // ── errors ────────────────────────────────────────────────────────────────
  { type: 'error', name: 'InvalidApplication', inputs: [] },
  { type: 'error', name: 'ApplicationAlreadyExists', inputs: [] },
  { type: 'error', name: 'InvalidSchnorrProof', inputs: [] },
  { type: 'error', name: 'InvalidEpoch', inputs: [] },
  { type: 'error', name: 'InvalidPhase', inputs: [] },
  { type: 'error', name: 'InvalidAddress', inputs: [] },
  { type: 'error', name: 'NotOwner', inputs: [] },
  { type: 'error', name: 'DecryptionNotYetAllowed', inputs: [] },
  { type: 'error', name: 'DecryptionExpired', inputs: [] },
  { type: 'error', name: 'DecryptionLimitReached', inputs: [] },
  // requireDecryptionOpen: before policy.decryptNotBefore.
  { type: 'error', name: 'DecryptionNotOpen', inputs: [] },
  // Past policy.decryptNotAfter: submitCiphertext, submitPartialDecryption
  // and combineDecryption all revert with it (the manager bubbles it up).
  { type: 'error', name: 'DecryptionClosed', inputs: [] },
  // requireDecryptionOpen: OrganizerLocked application whose organizerSecret
  // is still 0 — no partial and no combine exists on chain before the reveal.
  { type: 'error', name: 'OrganizerSecretNotRevealed', inputs: [] },
  // revealOrganizerSecret: sk·G != organizerPK (OrganizerLocked only —
  // Automatic applications have no organizer secret to reveal).
  { type: 'error', name: 'InvalidOrganizerSecret', inputs: [] },
  { type: 'error', name: 'AlreadyRevealed', inputs: [] },
  // registerApplication: no unclaimed pool key left this epoch.
  { type: 'error', name: 'PoolExhausted', inputs: [] },
  // registerApplication: the next unclaimed pool key hasn't been activated yet.
  { type: 'error', name: 'PoolKeyNotActive', inputs: [] },
  // Contradictory policy: open submission with a list, a zero address in the
  // list, more than 32 submitters, or a deadline not in the future.
  { type: 'error', name: 'InvalidPolicy', inputs: [] },
  { type: 'error', name: 'IsIdentity', inputs: [] },
  { type: 'error', name: 'NotCanonical', inputs: [] },
  { type: 'error', name: 'NotOnCurve', inputs: [] },
] as const;
