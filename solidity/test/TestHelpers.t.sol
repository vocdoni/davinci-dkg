// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../src/interfaces/IZKVerifier.sol";
import {BRLC} from "../src/libraries/BRLC.sol";
import {MAX_N, COMBINE_TRANSCRIPT_WORDS} from "../src/libraries/Sizes.sol";
import {BabyJubJub} from "../src/libraries/BabyJubJub.sol";
import {DKGProtocol} from "../src/libraries/DKGProtocol.sol";
import {TestInputs} from "./TestInputs.t.sol";

abstract contract TestHelpers is TestInputs {
    /// @dev Canonical on-curve ciphertext used by the ZK-mock tests.
    ///   c1 = 1·G  (the gnark RTE generator)
    ///   c2 = 4096·G (the SCHNORR_THIS vector's pubkey, secret = 0x1000)
    uint256 internal constant TEST_CT_C1X = 9671717474070082183213120605117400219616337014328744928644933853176787189663;
    uint256 internal constant TEST_CT_C1Y = 16950150798460657717958625567821834550301663161624707787222815936182638968203;
    uint256 internal constant TEST_CT_C2X = 17765672829315743641357949553430354448961270408100494783209553303687184365803;
    uint256 internal constant TEST_CT_C2Y = 13591243454297365848719372676992908085762757043204242277513940025707896351954;

    /// @dev Organizer secret of every application the tests register. Every
    ///      organizer artefact (registration PoP, decryption share, DLEQ) is
    ///      derived from it on the fly with `BabyJubJub`, so the fixtures stay
    ///      valid for the freshly-minted epoch ids the lottery hands out.
    uint256 internal constant TEST_ORG_SK = 12345;
    /// @dev Fixed nonces: the tests never need unpredictability, only
    ///      reproducibility.
    uint256 internal constant TEST_ORG_POP_NONCE = 67890;
    uint256 internal constant TEST_ORG_SHARE_NONCE = 424242;

    /// @notice The organizer's Chaum–Pedersen share of one ciphertext:
    ///         Δ = sk_org·C_1 with the DLEQ (A1 = w·G, A2 = w·C_1, z), plus
    ///         the challenge `e` the contract recomputes.
    struct OrgShare {
        uint256 deltaX;
        uint256 deltaY;
        uint256 a1x;
        uint256 a1y;
        uint256 a2x;
        uint256 a2y;
        uint256 z;
        uint256 e;
    }

    /// @notice Everything `combineDecryption` needs for one ciphertext.
    struct CombineFixture {
        bytes12 epochId;
        bytes32 aid;
        uint16 ctIdx;
        uint16 threshold;
        uint16 shareCount;
        bytes32 combineHash;
        uint256 plaintext;
        uint256 pkOrgX;
        uint256 pkOrgY;
        OrgShare share;
    }

    function testOrganizerPK() internal view returns (uint256 x, uint256 y) {
        return BabyJubJub.scalarMulBase(TEST_ORG_SK);
    }

    /// @dev Schnorr proof of possession of `sk_org` bound to (epochId, aid),
    ///      exactly as `DKGAppManager._organizerSchnorrChallenge` recomputes it.
    function organizerPoP(bytes12 epochId, bytes32 aid)
        internal
        view
        returns (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z)
    {
        (pkx, pky) = testOrganizerPK();
        (ax, ay) = BabyJubJub.scalarMulBase(TEST_ORG_POP_NONCE);
        uint256 c = uint256(keccak256(abi.encodePacked(
            DKGProtocol.DOMAIN_ORGANIZER_REGISTER_V1, epochId, aid, pkx, pky, ax, ay
        ))) % BabyJubJub.SUBGROUP_ORDER;
        z = addmod(
            TEST_ORG_POP_NONCE,
            mulmod(c, TEST_ORG_SK, BabyJubJub.SUBGROUP_ORDER),
            BabyJubJub.SUBGROUP_ORDER
        );
    }

    /// @dev Organizer decryption share of the canonical test ciphertext,
    ///      with the keccak challenge of protocol §3:
    ///
    ///        e = keccak(DOMAIN ‖ eid ‖ aid ‖ ctIdx ‖ PK_org ‖ C_1 ‖ Δ ‖ A1 ‖ A2) mod L
    ///        z = w + e·sk_org mod L
    function organizerShare(bytes12 epochId, bytes32 aid, uint16 ctIdx, uint256 c1x, uint256 c1y)
        internal
        view
        returns (OrgShare memory s)
    {
        (uint256 pkx, uint256 pky) = testOrganizerPK();
        (s.deltaX, s.deltaY) = BabyJubJub.scalarMul(TEST_ORG_SK, c1x, c1y);
        (s.a1x, s.a1y) = BabyJubJub.scalarMulBase(TEST_ORG_SHARE_NONCE);
        (s.a2x, s.a2y) = BabyJubJub.scalarMul(TEST_ORG_SHARE_NONCE, c1x, c1y);
        s.e = uint256(keccak256(abi.encodePacked(
            DKGProtocol.DOMAIN_ORGANIZER_SHARE_V1,
            epochId,
            aid,
            uint256(ctIdx),
            pkx, pky,
            c1x, c1y,
            s.deltaX, s.deltaY,
            s.a1x, s.a1y,
            s.a2x, s.a2y
        ))) % BabyJubJub.SUBGROUP_ORDER;
        s.z = addmod(
            TEST_ORG_SHARE_NONCE,
            mulmod(s.e, TEST_ORG_SK, BabyJubJub.SUBGROUP_ORDER),
            BabyJubJub.SUBGROUP_ORDER
        );
    }

    /// @dev Organizer share of the canonical test ciphertext (c1 = 1·G).
    function testOrganizerShare(bytes12 epochId, bytes32 aid, uint16 ctIdx)
        internal
        view
        returns (OrgShare memory)
    {
        return organizerShare(epochId, aid, ctIdx, TEST_CT_C1X, TEST_CT_C1Y);
    }

    /// @dev The combine fixture for the canonical test ciphertext.
    function combineFixture(
        bytes12 epochId,
        bytes32 aid,
        uint16 ctIdx,
        uint16 threshold,
        uint16 shareCount,
        bytes32 combineHash,
        uint256 plaintext
    ) internal view returns (CombineFixture memory f) {
        (uint256 pkx, uint256 pky) = testOrganizerPK();
        f = CombineFixture({
            epochId: epochId,
            aid: aid,
            ctIdx: ctIdx,
            threshold: threshold,
            shareCount: shareCount,
            combineHash: combineHash,
            plaintext: plaintext,
            pkOrgX: pkx,
            pkOrgY: pky,
            share: testOrganizerShare(epochId, aid, ctIdx)
        });
    }

    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:decrypt-combine:v1");
    bytes32 internal constant FINALIZE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:finalize:v1");

    function contributionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(1), 2, 3, 4, 5, 6, 7, 8]);
    }

    function contributionInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            CONTRIBUTION_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(commitmentsHash, encryptedSharesHash, keccak256(contributionTranscript(committeeSize))))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(threshold),
                uint256(committeeSize),
                uint256(contributorIndex),
                uint256(commitmentsHash),
                uint256(encryptedSharesHash),
                challenge,
                contributionTranscriptCommitment(challenge, committeeSize)
            ]
        );
    }

    // The test fixtures below mirror the on-chain transcript layouts,
    // parameterized in `MAX_N` (imported from src/libraries/Sizes.sol):
    //   contribution: 8N words = (2N + N + 2N + 2N + N) layout.
    //   finalize:     2N²+5N words.
    //   combine:      4 + 3N words.
    //   reconstruct:  2N words.
    /// @dev The first two committee slots use real on-curve BabyJubJub keys
    ///      (THIS and BEEF Schnorr vectors from cmd/operator-schnorr-vectors).
    ///      Slots 3+ stay at zero and are padded with the identity (0,1) below.
    ///      This matches what the DKGManager test fixtures register at setUp,
    ///      so the committee snapshot hash and the contribution transcript
    ///      stay byte-identical. All tests fix committeeSize = 2.
    function _slotPubKey(uint256 i) internal pure returns (uint256 x, uint256 y) {
        if (i == 0) {
            return (
                17765672829315743641357949553430354448961270408100494783209553303687184365803,
                13591243454297365848719372676992908085762757043204242277513940025707896351954
            );
        }
        if (i == 1) {
            return (
                10228722604559478181013548940833210623190136968531440936190496170400150013980,
                13886497050333420293068628977630539070604271411621054562122682889313139677221
            );
        }
        revert("test fixture: only 2 committee members supported");
    }

    function contributionTranscript(uint16 committeeSize) internal pure returns (bytes memory) {
        uint256[2 * MAX_N] memory commitments;
        uint256[MAX_N] memory recipientIndexes;
        uint256[2 * MAX_N] memory recipientPubKeys;
        uint256[2 * MAX_N] memory ephemerals;
        uint256[MAX_N] memory maskedShares;
        for (uint256 i = 0; i < MAX_N; i++) {
            commitments[i * 2 + 1] = 1;
            recipientPubKeys[i * 2 + 1] = 1;
            ephemerals[i * 2 + 1] = 1;
        }
        for (uint256 i = 0; i < committeeSize; i++) {
            commitments[i * 2 + 1] = 0;
            recipientIndexes[i] = i + 1;
            (uint256 px, uint256 py) = _slotPubKey(i);
            recipientPubKeys[i * 2] = px;
            recipientPubKeys[i * 2 + 1] = py;
            ephemerals[i * 2] = 300 + i + 1;
            ephemerals[i * 2 + 1] = 400 + i + 1;
            maskedShares[i] = 500 + i + 1;
        }
        return abi.encode(commitments, recipientIndexes, recipientPubKeys, ephemerals, maskedShares);
    }

    function contributionTranscriptCommitment(uint256 challenge, uint16 committeeSize) internal pure returns (uint256) {
        uint256[] memory values = new uint256[](8 * MAX_N);
        for (uint256 i = 0; i < MAX_N; i++) {
            values[i * 2 + 1] = 1;                       // commitments y pad
            values[3 * MAX_N + i * 2 + 1] = 1;           // recipientPubKeys y pad (offset 2N+N)
            values[5 * MAX_N + i * 2 + 1] = 1;           // ephemerals y pad (offset 2N+N+2N)
        }
        uint256 cursor = 2 * MAX_N; // recipientIndexes start (after 2N commitments)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[i * 2 + 1] = 0;
            values[cursor++] = i + 1;
        }
        cursor = 3 * MAX_N; // recipientPubKeys start (2N+N)
        for (uint256 i = 0; i < committeeSize; i++) {
            (uint256 px, uint256 py) = _slotPubKey(i);
            values[cursor++] = px;
            values[cursor++] = py;
        }
        cursor = 5 * MAX_N; // ephemerals start (2N+N+2N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 300 + i + 1;
            values[cursor++] = 400 + i + 1;
        }
        cursor = 7 * MAX_N; // maskedShares start (2N+N+2N+2N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 500 + i + 1;
        }
        return BRLC.commit(challenge, values);
    }

    function partialDecryptionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(11), 12, 13, 14, 15, 16, 17, 18]);
    }

    /// partialdecrypt layout (15 public inputs):
    ///   [0] eid, [1] aid, [2] ctIdx, [3] participantIndex,
    ///   [4..5] C1.x/y, [6..7] D_i.x/y, [8..9] delta.x/y,
    ///   [10..11] A1.x/y, [12..13] A2.x/y, [14] response.
    /// 3-arg overload defaults ciphertextIndex to 1 for the most common case.
    function partialDecryptionInput(bytes12 epochId, bytes32 aid, uint16 participantIndex)
        internal
        pure
        returns (bytes memory)
    {
        return partialDecryptionInputCt(epochId, aid, participantIndex, 1);
    }

    function partialDecryptionInputCt(
        bytes12 epochId,
        bytes32 aid,
        uint16 participantIndex,
        uint16 ciphertextIndex
    ) internal pure returns (bytes memory) {
        uint256[15] memory inputs;
        inputs[0] = uint256(uint96(epochId));
        inputs[1] = uint256(aid);
        inputs[2] = ciphertextIndex;
        inputs[3] = participantIndex;
        // pi[4..5] = C1, must match the test ciphertext fixture so
        // submitPartialDecryption's ciphertext binding accepts.
        inputs[4] = TEST_CT_C1X;
        inputs[5] = TEST_CT_C1Y;
        inputs[6] = 1000 + participantIndex;    // D_i.x
        inputs[7] = 2000 + participantIndex;    // D_i.y
        inputs[8] = 7000 + participantIndex;    // delta.x
        inputs[9] = 8000 + participantIndex;    // delta.y
        return abi.encode(inputs);
    }

    function partialDecryptionHash(uint16 participantIndex) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint256(7000 + participantIndex), uint256(8000 + participantIndex)));
    }

    function finalizeProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(21), 22, 23, 24, 25, 26, 27, 28]);
    }

    function finalizeInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 acceptedCount,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            FINALIZE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(
                aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash,
                FINALIZED_ROWS_HASH, keccak256(finalizeTranscript(acceptedCount))
            ))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(threshold),
                uint256(committeeSize),
                uint256(acceptedCount),
                uint256(aggregateCommitmentsHash),
                uint256(collectivePublicKeyHash),
                uint256(shareCommitmentHash),
                uint256(FINALIZED_ROWS_HASH),
                challenge,
                finalizeTranscriptCommitment(challenge, acceptedCount)
            ]
        );
    }

    function finalizeTranscript(uint16 acceptedCount) internal pure returns (bytes memory) {
        uint256[MAX_N] memory participantIndexes;
        uint256[2 * MAX_N * MAX_N] memory contributionCommitments;
        uint256[2 * MAX_N] memory aggregateCommitments;
        uint256[2 * MAX_N] memory shareCommitments;
        // Per participant, mirror contributionTranscript(committeeSize=2):
        // pt[0]=pt[1]=(0,0); pt[2..N-1]=(0,1) → odd indices 5,7,...,2N-1 = 1
        // (each participant's commitments occupy 2N words).
        for (uint256 i = 0; i < acceptedCount; i++) {
            participantIndexes[i] = i + 1;
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                contributionCommitments[i * (2 * MAX_N) + k] = 1;
            }
            shareCommitments[i * 2] = 1000 + i + 1;
            shareCommitments[i * 2 + 1] = 2000 + i + 1;
        }
        // The contract requires aggregate[0] to equal
        // the accumulated _collectiveKey. The fixtures all use commitment0 =
        // (0, 1) (identity), so the running sum is identity (0, 1).
        aggregateCommitments[1] = 1;
        return abi.encode(participantIndexes, contributionCommitments, aggregateCommitments, shareCommitments);
    }

    function finalizeTranscriptCommitment(uint256 challenge, uint16 acceptedCount) internal pure returns (uint256) {
        // Layout (2N²+5N words): [0..N) participantIndexes,
        //                        [N..N+2N²) contributionCommitments,
        //                        [N+2N²..N+2N²+2N) aggregateCommitments,
        //                        [N+2N²+2N..2N²+5N) shareCommitments.
        uint256[] memory values = new uint256[](2 * MAX_N * MAX_N + 5 * MAX_N);
        for (uint256 i = 0; i < acceptedCount; i++) {
            values[i] = i + 1;
            uint256 offset = MAX_N + i * (2 * MAX_N); // N + i*2N
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                values[offset + k] = 1;
            }
        }
        // shareCommitments start: N + 2N² + 2N
        uint256 shareOffsetBase = MAX_N + 2 * MAX_N * MAX_N + 2 * MAX_N;
        for (uint256 i = 0; i < acceptedCount; i++) {
            uint256 offset = shareOffsetBase + i * 2;
            values[offset] = 1000 + i + 1;
            values[offset + 1] = 2000 + i + 1;
        }
        // aggregateCommitments[0] = identity (0, 1) — see finalizeTranscript.
        // Slot offset is N + 2N²; aggregate[0].y is at offset+1.
        values[MAX_N + 2 * MAX_N * MAX_N + 1] = 1;
        return BRLC.commit(challenge, values);
    }

    /// @dev Variant of finalizeTranscript used by the duplicate-row regression
    /// test. Sets participantIndexes = [1, 1] for an
    /// acceptedCount of 2 instead of [1, 2].
    function finalizeTranscriptWithDuplicateRows() internal pure returns (bytes memory) {
        uint256[MAX_N] memory participantIndexes;
        uint256[2 * MAX_N * MAX_N] memory contributionCommitments;
        uint256[2 * MAX_N] memory aggregateCommitments;
        uint256[2 * MAX_N] memory shareCommitments;
        participantIndexes[0] = 1;
        participantIndexes[1] = 1;
        for (uint256 i = 0; i < 2; i++) {
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                contributionCommitments[i * (2 * MAX_N) + k] = 1;
            }
            shareCommitments[i * 2] = 1000 + i + 1;
            shareCommitments[i * 2 + 1] = 2000 + i + 1;
        }
        aggregateCommitments[1] = 1;
        return abi.encode(participantIndexes, contributionCommitments, aggregateCommitments, shareCommitments);
    }

    function finalizeInputWithDuplicateRows(
        bytes12 epochId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            FINALIZE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(
                aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash,
                FINALIZED_ROWS_HASH, keccak256(finalizeTranscriptWithDuplicateRows())
            ))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(2),
                uint256(2),
                uint256(2),
                uint256(aggregateCommitmentsHash),
                uint256(collectivePublicKeyHash),
                uint256(shareCommitmentHash),
                uint256(FINALIZED_ROWS_HASH),
                challenge,
                _finalizeTranscriptCommitmentWithDuplicateRows(challenge)
            ]
        );
    }

    function _finalizeTranscriptCommitmentWithDuplicateRows(uint256 challenge) private pure returns (uint256) {
        uint256[] memory values = new uint256[](2 * MAX_N * MAX_N + 5 * MAX_N);
        values[0] = 1;
        values[1] = 1;
        for (uint256 i = 0; i < 2; i++) {
            uint256 offset = MAX_N + i * (2 * MAX_N);
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                values[offset + k] = 1;
            }
        }
        uint256 shareOffsetBase = MAX_N + 2 * MAX_N * MAX_N + 2 * MAX_N;
        for (uint256 i = 0; i < 2; i++) {
            uint256 offset = shareOffsetBase + i * 2;
            values[offset] = 1000 + i + 1;
            values[offset + 1] = 2000 + i + 1;
        }
        values[MAX_N + 2 * MAX_N * MAX_N + 1] = 1;
        return BRLC.commit(challenge, values);
    }

    function decryptCombineProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(31), 32, 33, 34, 35, 36, 37, 38]);
    }

    /// @dev The 12 + 3N transcript words of `combineDecryption`, in the exact
    ///      order the contract reads them out of calldata:
    ///        w[0..3]   C1.x C1.y C2.x C2.y
    ///        w[4..5]   PK_org.x PK_org.y
    ///        w[6..7]   A1.x A1.y
    ///        w[8..9]   A2.x A2.y
    ///        w[10]     z          w[11] e
    ///        w[12 .. 12+N)        participant indexes (0 when inactive)
    ///        w[12+N .. 12+3N)     partial decryptions (identity when inactive)
    function combineWords(CombineFixture memory f) internal pure returns (uint256[] memory v) {
        v = new uint256[](COMBINE_TRANSCRIPT_WORDS);
        v[0] = TEST_CT_C1X;
        v[1] = TEST_CT_C1Y;
        v[2] = TEST_CT_C2X;
        v[3] = TEST_CT_C2Y;
        v[4] = f.pkOrgX;
        v[5] = f.pkOrgY;
        v[6] = f.share.a1x;
        v[7] = f.share.a1y;
        v[8] = f.share.a2x;
        v[9] = f.share.a2y;
        v[10] = f.share.z;
        v[11] = f.share.e;
        uint256 partialBase = 12 + MAX_N;
        for (uint256 i = 0; i < MAX_N; i++) {
            v[partialBase + i * 2 + 1] = 1; // identity padding
        }
        for (uint256 i = 0; i < f.shareCount; i++) {
            v[12 + i] = i + 1;
            v[partialBase + i * 2] = 7000 + i + 1;
            v[partialBase + i * 2 + 1] = 8000 + i + 1;
        }
    }

    function combineTranscript(CombineFixture memory f) internal pure returns (bytes memory) {
        return abi.encodePacked(combineWords(f));
    }

    /// @dev 11-element public-input vector matching the combine circuit and
    ///      verifier: eid, aid, ctIdx, Δ.x, Δ.y, threshold, shareCount,
    ///      combineHash, plaintext, challenge, transcriptCommitment.
    function combineInput(CombineFixture memory f) internal pure returns (bytes memory) {
        return combineInputForWords(f, combineWords(f));
    }

    /// @dev Same, but over an arbitrary (possibly tampered) word vector, so a
    ///      test can hand the contract a transcript whose ρ and BRLC
    ///      commitment are internally consistent and still see the on-chain
    ///      bindings reject it.
    function combineInputForWords(CombineFixture memory f, uint256[] memory words)
        internal
        pure
        returns (bytes memory)
    {
        uint256 challenge = BRLC.deriveChallenge(
            f.epochId,
            DECRYPT_COMBINE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(
                f.combineHash, bytes32(f.plaintext), keccak256(abi.encodePacked(words))
            ))
        );
        uint256[11] memory v;
        v[0] = uint256(uint96(f.epochId));
        v[1] = uint256(f.aid);
        v[2] = uint256(f.ctIdx);
        v[3] = f.share.deltaX;
        v[4] = f.share.deltaY;
        v[5] = uint256(f.threshold);
        v[6] = uint256(f.shareCount);
        v[7] = uint256(f.combineHash);
        v[8] = f.plaintext;
        v[9] = challenge;
        v[10] = BRLC.commit(challenge, words);
        return abi.encode(v);
    }

}

contract MockContributionVerifier is IZKVerifier, TestInputs {
    error InvalidProofInput();

    function verifyProof(bytes calldata proof, bytes calldata input) external pure override {
        if (proof.length == 0 || keccak256(input) == keccak256(CONTRIBUTION_INPUT_BAD)) {
            revert InvalidProofInput();
        }
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return CONTRIBUTION_PROVING_KEY_HASH;
    }
}

contract MockPartialDecryptVerifier is IZKVerifier, TestInputs {
    error InvalidProofInput();

    function verifyProof(bytes calldata proof, bytes calldata input) external pure override {
        if (proof.length == 0 || keccak256(input) == keccak256(PARTIAL_DECRYPTION_INPUT_BAD)) revert InvalidProofInput();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return PARTIAL_DECRYPTION_PROVING_KEY_HASH;
    }
}

contract MockFinalizeVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return FINALIZE_PROVING_KEY_HASH;
    }
}

contract MockDecryptCombineVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return DECRYPT_COMBINE_PROVING_KEY_HASH;
    }
}
