// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

library DKGTypes {
    struct Point {
        uint256 x;
        uint256 y;
    }

    /// @notice Epoch state machine. The first three values group into the
    ///         "Preparation" phase (committee assembly + key generation); `Live`
    ///         is the "Service" phase in which apps register their derived keys
    ///         and ciphertexts get decrypted under the collective key.
    enum EpochPhase {
        None,
        CommitteeSelection,  // accepting claimSlot calls; lottery picks the committee
        KeyAssembly,         // committee submits Feldman VSS contributions
        Live,                // collective key PK_ep is live; apps + decryption are open
        Aborted,
        Completed
    }

    /// @notice Per-epoch DKG policy. The phase deadline blocks are derived at
    ///         `createEpoch` time from `DKGManager.EPOCH_DURATION_BLOCKS` and
    ///         the BPS constants in `Sizes.sol` — they are stored on the struct
    ///         so downstream phase checks remain a single SLOAD. Layout: 4 ×
    ///         uint16 + 3 × uint64 = 256 bits = exactly 1 storage slot.
    struct EpochPolicy {
        uint16 threshold;
        uint16 committeeSize;
        uint16 minValidContributions;
        uint16 lotteryAlphaBps;                  // candidate-pool size = α × committeeSize, α expressed in basis points (10000 = 1.0)
        uint64 committeeSelectionDeadlineBlock;  // last block in which claimSlot is accepted
        uint64 keyAssemblyDeadlineBlock;         // last block in which submitContribution is accepted
        uint64 liveNotBeforeBlock;               // earliest block at which finalizeEpoch can flip the epoch to Live
    }

    struct ContributionRecord {
        address contributor;
        uint16 contributorIndex;
        bytes32 commitmentsHash;
        bytes32 encryptedSharesHash;
        bytes32 commitmentVectorDigest;
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
    // An Application is registered against a Live Epoch by its organizer and
    // obtains the encryption key `PK_aid = PK_ep + PK_org`. Decryption needs
    // both the committee's threshold partial decryptions and the organizer's
    // share `Δ = sk_org · C_1`.

    /// @notice Per-application access policy. Gates submitCiphertext with the same
    ///         semantics but is scoped per application rather than per epoch.
    ///         All checks AND together; a zero-valued field is a no-op — except
    ///         `authorizedSubmitter`, which is resolved to the registering
    ///         address at registration time (there is no open submission).
    struct AppPolicy {
        address authorizedSubmitter;  // resolved to the registrant when passed as 0
        uint16  maxCiphertexts;       // 0 = unlimited (capped by MAX_CIPHERTEXT_INDEX)
        uint64  notBeforeBlock;
        uint64  notAfterBlock;
    }

    /// @notice On-chain application record. `aid` is keyed in the app manager's
    ///         per-epoch mapping; it is not duplicated here.
    struct Application {
        address creator;       // who called registerApplication
        Point   organizerPK;   // PK_org, proven at registration
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
