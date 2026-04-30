// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

library DKGTypes {
    struct Point {
        uint256 x;
        uint256 y;
    }

    enum EpochPhase {
        None,
        Registration,    // accepting claimSlot calls (replaces Readiness)
        Contribution,
        Finalized,
        Aborted,
        Completed
    }

    struct EpochPolicy {
        uint16 threshold;
        uint16 committeeSize;
        uint16 minValidContributions;
        uint16 lotteryAlphaBps;            // candidate-pool size = α × committeeSize, α expressed in basis points (10000 = 1.0)
        uint16 seedDelay;                  // blocks between createEpoch and the block whose hash becomes the seed
        uint64 registrationDeadlineBlock;  // last block in which claimSlot is accepted
        uint64 contributionDeadlineBlock;  // last block in which submitContribution is accepted
        uint64 finalizeNotBeforeBlock;     // earliest block at which finalizeEpoch can succeed; must be > contributionDeadlineBlock
    }

    /// @notice Gates `submitCiphertext` for a epoch. All checks AND together; an
    ///         unset (zero) field is a no-op for that check.
    ///         The policy only gates SUBMISSION; once a ciphertext is on-chain,
    ///         decryption by the committee proceeds regardless of these fields.
    struct DecryptionPolicy {
        bool   ownerOnly;           // if true, only the epoch organizer can submitCiphertext
        uint16 maxDecryptions;      // max ciphertexts accepted per epoch; 0 = unlimited (up to MAX_CIPHERTEXT_INDEX)
        uint64 notBeforeBlock;      // submitCiphertext reverts if block.number < this; 0 = no lock
        uint64 notBeforeTimestamp;  // submitCiphertext reverts if block.timestamp < this; 0 = no lock
        uint64 notAfterBlock;       // submitCiphertext reverts if block.number > this; 0 = no deadline
        uint64 notAfterTimestamp;   // submitCiphertext reverts if block.timestamp > this; 0 = no deadline
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
        address participant;
        uint16 participantIndex;
        uint16 ciphertextIndex;
        bytes32 deltaHash;
        Point delta;
        bool accepted;
    }

    struct CombinedDecryptionRecord {
        uint16 ciphertextIndex;
        bool completed;
        uint256 plaintext;
    }

    // ─── Application surface (paper §4.3, PLAN.md §4.3) ─────────────────────
    //
    // An Application is registered against a finalized Epoch and obtains a
    // unique encryption key derived from `PK_ep` plus a per-application
    // correction term selected by `mode`. See `solidity/src/libraries/DKGProtocol.sol`
    // for the canonical mode and role constants.

    /// @notice Application registration mode.
    /// @dev Values must match `DKGProtocol.MODE_*`. Stored as `uint8` rather
    ///      than as an enum so the on-chain layout matches the value the
    ///      combine circuit consumes as a `frontend.Variable`.
    enum AppMode {
        PublicDerivation, // = 0, see paper §4.3
        OrganizerCoDec    // = 1, see paper §6
    }

    /// @notice DLEQ role for Chaum-Pedersen partial decryptions.
    /// @dev Values 1 (committee) and 2 (organizer); see paper §4.4 / §6.3.
    ///      The enum starts at None=0 so an uninitialized record cannot be
    ///      mistaken for a valid committee proof.
    enum Role {
        None,      // = 0  (uninitialized)
        Committee, // = 1
        Organizer  // = 2
    }

    /// @notice Per-application access policy. Mirrors DecryptionPolicy
    ///         semantics but is scoped per application rather than per epoch.
    ///         All checks AND together; a zero-valued field is a no-op.
    struct AppPolicy {
        address authorizedSubmitter;  // 0 = open (anyone can submitCiphertext)
        uint16  maxCiphertexts;       // 0 = unlimited (capped by MAX_CIPHERTEXT_INDEX)
        uint64  notBeforeBlock;
        uint64  notAfterBlock;
    }

    /// @notice On-chain application record. `aid` is keyed in the manager's
    ///         per-epoch mapping; it is not duplicated here.
    struct Application {
        address creator;       // who called registerApplication / registerApplicationCoDec
        AppMode mode;          // 0 = public derivation; 1 = organizer co-decryption
        uint256 derivationS;   // S = keccak256(eid || PK_ep || aid) % q  (mode 0 only; 0 in mode 1)
        Point   organizerPK;   // PK_org (mode 1 only; identity in mode 0)
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

    /// @notice Organizer's Δ_org submission (mode 1 only). Verified once via
    ///         the DLEQ proof at submitOrganizerShare time, then consumed by
    ///         combineDecryption as the correction point.
    struct OrganizerShareRecord {
        Point   deltaOrg;
        bytes32 dleqHash;       // commitment to (A, B, z)
        bool    accepted;
    }
}
