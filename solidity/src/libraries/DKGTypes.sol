// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

library DKGTypes {
    struct Point {
        uint256 x;
        uint256 y;
    }

    /// @notice Epoch state machine. The first three values group into the
    ///         "Preparation" phase (committee assembly + key generation); `Live`
    ///         is the "Service" phase in which pool keys activate, apps claim
    ///         them and ciphertexts get decrypted.
    enum EpochPhase {
        None,
        CommitteeSelection,  // accepting claimSlot calls; lottery picks the committee
        KeyAssembly,         // committee submits Feldman VSS contributions
        Live,                // pool keys may activate; apps + decryption are open
        Aborted,
        Completed
    }

    /// @notice Per-epoch DKG policy. The phase deadline blocks are derived at
    ///         `createEpoch` time from `DKGManager.EPOCH_DURATION_BLOCKS` and
    ///         the per-phase constants in `Sizes.sol` — they are stored on the
    ///         struct so downstream phase checks remain a single SLOAD.
    ///         Layout: 4 × uint16 + 3 × uint64 = 256 bits = exactly 1 storage slot.
    struct EpochPolicy {
        uint16 threshold;
        uint16 committeeSize;
        uint16 minValidContributions;
        uint16 lotteryAlphaBps;                  // candidate-pool size = α × committeeSize, α expressed in basis points (10000 = 1.0)
        uint64 committeeSelectionDeadlineBlock;  // last block in which claimSlot is accepted
        uint64 keyAssemblyDeadlineBlock;         // last block in which submitContribution is accepted
        uint64 liveNotBeforeBlock;               // earliest block at which finalizeEpoch can flip the epoch to Live
    }

    /// @notice One committee member's accepted contribution.
    /// @dev    `commitmentsHash` is the Poseidon digest of public input 4 of
    ///         the contribution proof — it commits to the member's `MAX_K`
    ///         commitment vectors at once. `activatePoolKey` re-checks each
    ///         participant row against it, which is what binds an activation
    ///         to the contributions actually accepted on chain.
    struct ContributionRecord {
        address contributor;
        uint16 contributorIndex;
        bytes32 commitmentsHash;
        bytes32 encryptedSharesHash;
        bool accepted;
    }

    struct PartialDecryptionRecord {
        uint16 participantIndex;
        uint16 ciphertextIndex;
        bytes32 deltaHash;
        bool accepted;
    }

    struct CombinedDecryptionRecord {
        uint16 ciphertextIndex;
        bool completed;
        uint256 plaintext;
    }

    // ─── Application surface ────────────────────────────────────────────────
    //
    // An Application is registered against a Live Epoch and claims one of the
    // epoch's `MAX_K` committee-held pool keys `P_j`. Its encryption key is
    // `PK_aid = P_j` (Automatic) or `PK_aid = P_j + PK_org`
    // (OrganizerLocked). Because every application holds a *different*
    // committee key, a ciphertext copied into another application decrypts
    // under a key that has nothing to do with it.

    /// @notice How an application's ciphertexts get decrypted.
    ///
    ///         `OrganizerLocked`: `PK_org` is registered with a Schnorr proof
    ///         of possession and `PK_aid = P_j + PK_org`. The committee alone
    ///         only ever recovers the `P_j` half; the organizer later calls
    ///         `revealOrganizerSecret` once (or never), after which the
    ///         committee combines by itself.
    ///
    ///         `Automatic`: there is no organizer key at all. `organizerPK` is
    ///         stored as the identity `(0, 1)` and `organizerSecret` stays 0,
    ///         so `PK_aid = P_j` and decryption needs nobody but the
    ///         committee. Confidentiality rests on the committee threshold and
    ///         on `P_j` being this application's own key.
    enum AppMode {
        OrganizerLocked,
        Automatic
    }

    /// @notice Per-application policy, fixed at registration. Submission is
    ///         open to anyone when `openSubmission` is set, else to the
    ///         listed `submitters`, else to the registrant only. Windows AND
    ///         together; a zero value is a no-op.
    /// @dev    The submission window is measured in BLOCKS, the decryption
    ///         window in unix SECONDS: submission races the chain's own
    ///         clock, while the decryption window is what an organizer
    ///         advertises to voters and must therefore read as wall time.
    struct AppPolicy {
        AppMode   mode;
        bool      openSubmission;    // anyone may submitCiphertext
        address[] submitters;        // allow-list; empty = the registrant only
        uint16    maxCiphertexts;    // 0 = unlimited (capped by MAX_CIPHERTEXT_INDEX)
        uint64    notBeforeBlock;    // submitCiphertext window (blocks)
        uint64    notAfterBlock;
        uint64    decryptNotBefore;  // unix seconds; partials and combines revert before it (0 = none)
        uint64    decryptNotAfter;   // unix seconds; partials and combines revert after it  (0 = none)
    }

    /// @notice On-chain application record. `aid` is keyed in the app manager's
    ///         per-epoch mapping; it is not duplicated here.
    struct Application {
        address creator;          // who called registerApplication
        Point   organizerPK;      // PK_org, proven at registration; (0, 1) when Automatic
        uint256 organizerSecret;  // sk_org once revealed; 0 while sealed and always when Automatic
        uint8   poolIndex;        // which of the epoch's MAX_K pool keys this app claimed
        AppPolicy policy;
        uint64  createdAtBlock;
        bool    exists;
    }

    /// @notice Per-(epoch, app, ctIdx) ciphertext record.
    struct CiphertextRecord {
        Point   c1;
        Point   c2;
        address submitter;
        uint64  submittedAtBlock;
        bool    exists;
    }
}
