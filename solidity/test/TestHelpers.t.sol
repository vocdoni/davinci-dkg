// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../src/interfaces/IZKVerifier.sol";
import {BRLC} from "../src/libraries/BRLC.sol";
import {
    MAX_N,
    MAX_K,
    MERKLE_DEPTH,
    CONTRIB_TRANSCRIPT_WORDS,
    POOLKEY_TRANSCRIPT_WORDS,
    COMBINE_TRANSCRIPT_WORDS
} from "../src/libraries/Sizes.sol";
import {BabyJubJub} from "../src/libraries/BabyJubJub.sol";
import {DKGProtocol} from "../src/libraries/DKGProtocol.sol";
import {TestInputs} from "./TestInputs.t.sol";

/// @title  TestHelpers
/// @notice Canned transcripts, public-input vectors and Merkle fixtures for
///         the mock-verifier suites. Every builder here mirrors the layouts
///         documented in `src/libraries/Sizes.sol` word for word: the mocks
///         wave the pairing check through, so the transcript bindings the
///         contract *does* enforce (committee prefix hash, participant rows,
///         Fiat–Shamir challenge, BRLC commitment, Merkle paths) are exactly
///         what these tests exercise.
///
///         All builders derive the calldata bytes AND the BRLC commitment
///         from one flat `uint256[]` word vector, so the two can never drift.
abstract contract TestHelpers is TestInputs {
    /// @dev Canonical on-curve ciphertext used by the ZK-mock tests.
    ///   c1 = 1·G  (the gnark RTE generator)
    ///   c2 = 4096·G (the SCHNORR_THIS vector's pubkey, secret = 0x1000)
    uint256 internal constant TEST_CT_C1X = 9671717474070082183213120605117400219616337014328744928644933853176787189663;
    uint256 internal constant TEST_CT_C1Y = 16950150798460657717958625567821834550301663161624707787222815936182638968203;
    uint256 internal constant TEST_CT_C2X = 17765672829315743641357949553430354448961270408100494783209553303687184365803;
    uint256 internal constant TEST_CT_C2Y = 13591243454297365848719372676992908085762757043204242277513940025707896351954;

    /// @dev Organizer secret of every `OrganizerLocked` application the tests
    ///      register. The registration PoP is derived from it on the fly with
    ///      `BabyJubJub`, so the fixtures stay valid for the freshly-minted
    ///      epoch ids the lottery hands out.
    uint256 internal constant TEST_ORG_SK = 12345;
    /// @dev Fixed nonce: the tests never need unpredictability, only
    ///      reproducibility.
    uint256 internal constant TEST_ORG_POP_NONCE = 67890;

    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:decrypt-combine:v1");
    bytes32 internal constant POOLKEY_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:poolkey:v1");

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

    // ─── Committee fixture ────────────────────────────────────────────────────

    /// @dev The first two committee slots use real on-curve BabyJubJub keys
    ///      (THIS and BEEF Schnorr vectors from cmd/operator-schnorr-vectors).
    ///      Slots 3+ stay at zero and are padded with the identity (0,1).
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

    // ─── Contribution ─────────────────────────────────────────────────────────

    function contributionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(1), 2, 3, 4, 5, 6, 7, 8]);
    }

    /// @dev The `3KN + 5N` contribution transcript words, in the order the
    ///      contract reads them from calldata:
    ///        [0, 2KN)          commitments, key-major (identity padded)
    ///        [2KN, 2KN+N)      recipientIndexes
    ///        [2KN+N, 2KN+3N)   recipientPubKeys
    ///        [2KN+3N, 2KN+5N)  ephemerals
    ///        [2KN+5N, 3KN+5N)  maskedShares, key-major
    ///      Only the committee section is bound to on-chain state; the rest
    ///      just has to be internally consistent with the BRLC commitment.
    function contributionWords(uint16 committeeSize) internal pure returns (uint256[] memory v) {
        v = new uint256[](CONTRIB_TRANSCRIPT_WORDS);
        uint256 idxBase = 2 * MAX_K * MAX_N;
        uint256 keyBase = idxBase + MAX_N;
        uint256 ephBase = idxBase + 3 * MAX_N;
        uint256 msBase  = idxBase + 5 * MAX_N;
        // Identity padding: every commitment and every unused point is (0, 1).
        for (uint256 i = 0; i < MAX_K * MAX_N; i++) {
            v[i * 2 + 1] = 1;
        }
        for (uint256 i = 0; i < MAX_N; i++) {
            v[keyBase + i * 2 + 1] = 1;
            v[ephBase + i * 2 + 1] = 1;
        }
        for (uint256 i = 0; i < committeeSize; i++) {
            v[idxBase + i] = i + 1;
            (uint256 px, uint256 py) = _slotPubKey(i);
            v[keyBase + i * 2] = px;
            v[keyBase + i * 2 + 1] = py;
            v[ephBase + i * 2] = 300 + i + 1;
            v[ephBase + i * 2 + 1] = 400 + i + 1;
            for (uint256 j = 0; j < MAX_K; j++) {
                v[msBase + j * MAX_N + i] = 500 + j * 100 + i + 1;
            }
        }
    }

    function contributionTranscript(uint16 committeeSize) internal pure returns (bytes memory) {
        return abi.encodePacked(contributionWords(committeeSize));
    }

    function contributionInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    ) internal pure returns (bytes memory) {
        return contributionInputForWords(
            epochId, threshold, committeeSize, contributorIndex, commitmentsHash, encryptedSharesHash,
            contributionWords(committeeSize)
        );
    }

    /// @dev Same, over an arbitrary (possibly tampered) word vector: the
    ///      challenge and the BRLC commitment are recomputed over `words`, so
    ///      the contract gets past its Fiat–Shamir check and the test reaches
    ///      whichever binding it wants to exercise.
    function contributionInputForWords(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        uint256[] memory words
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            CONTRIBUTION_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(
                commitmentsHash, encryptedSharesHash, keccak256(abi.encodePacked(words))
            ))
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
                BRLC.commit(challenge, words)
            ]
        );
    }

    /// @dev The Poseidon `commitmentsHash` each test contributor submits.
    function contributorCommitmentsHash(uint16 contributorIndex) internal pure returns (bytes32) {
        return bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + contributorIndex - 1);
    }

    function contributorSharesHash(uint16 contributorIndex) internal pure returns (bytes32) {
        return bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + contributorIndex - 1);
    }

    // ─── Pool-key activation ──────────────────────────────────────────────────

    function poolKeyProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(21), 22, 23, 24, 25, 26, 27, 28]);
    }

    /// @dev The prover's transcript digest (public input 5). Any
    ///      deterministic value will do: the mock verifier accepts the proof
    ///      and the contract only checks it against the call argument and
    ///      feeds it into the challenge anchor.
    function testTranscriptDigest(uint8 keyIndex) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("test-poolkey-transcript-digest", keyIndex));
    }

    /// @dev `P_j` of the fixtures: a distinct value per key so a test can tell
    ///      the epoch's keys apart. `activatePoolKey` stores
    ///      `aggregateCommitments[0]` verbatim — proving it is a real group
    ///      element is the circuit's job, not the contract's — so the fixture
    ///      does not need an on-curve point, and staying `pure` keeps every
    ///      builder safe to evaluate as an argument after `vm.expectRevert`.
    function testPoolKey(uint8 keyIndex) internal pure returns (uint256 x, uint256 y) {
        return (0xB0071 + uint256(keyIndex), 0xB0072 + uint256(keyIndex));
    }

    /// @dev The committee member's share commitment `D_p` of pool key
    ///      `keyIndex`. Distinct per key: a partial decryption proving `D_p`
    ///      of one key can never produce a Merkle path into another key's
    ///      root.
    function testShareCommitment(uint8 keyIndex, uint16 participantIndex)
        internal
        pure
        returns (uint256 x, uint256 y)
    {
        uint256 offset = 1000 * uint256(keyIndex) + uint256(participantIndex);
        return (1000 + offset, 2000 + offset);
    }

    /// @dev The `6N` pool-key activation transcript words:
    ///        [0, N)    participantIndexes   (rows < acceptedCount, 0 beyond)
    ///        [N, 2N)   contributionHashes   (rows < acceptedCount, 0 beyond)
    ///        [2N, 4N)  aggregateCommitments (aggregate[0] = P_j)
    ///        [4N, 6N)  shareCommitments     (slot i = member i + 1 for
    ///                                        i < committeeSize, else (0, 1))
    ///      The accepted contributors are members 1..acceptedCount.
    function poolKeyWords(uint8 keyIndex, uint16 acceptedCount, uint16 committeeSize)
        internal
        pure
        returns (uint256[] memory v)
    {
        v = new uint256[](POOLKEY_TRANSCRIPT_WORDS);
        uint256 aggBase = 2 * MAX_N;
        uint256 scBase = 4 * MAX_N;
        for (uint256 i = 0; i < MAX_N; i++) {
            v[aggBase + i * 2 + 1] = 1; // identity padding for m >= t
            v[scBase + i * 2 + 1] = 1;  // identity padding beyond the committee
        }
        (uint256 pkx, uint256 pky) = testPoolKey(keyIndex);
        v[aggBase] = pkx;
        v[aggBase + 1] = pky;
        for (uint16 i = 0; i < acceptedCount; i++) {
            v[i] = uint256(i) + 1;
            v[MAX_N + i] = uint256(contributorCommitmentsHash(i + 1));
        }
        for (uint16 i = 0; i < committeeSize; i++) {
            (uint256 scx, uint256 scy) = testShareCommitment(keyIndex, i + 1);
            v[scBase + uint256(i) * 2] = scx;
            v[scBase + uint256(i) * 2 + 1] = scy;
        }
    }

    function poolKeyTranscript(uint8 keyIndex, uint16 acceptedCount, uint16 committeeSize)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(poolKeyWords(keyIndex, acceptedCount, committeeSize));
    }

    /// @dev 8-element public-input vector: eid, t, n, acceptedCount, keyIndex,
    ///      transcriptDigest, challenge, transcriptCommitment.
    function poolKeyInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 acceptedCount,
        uint8 keyIndex
    ) internal pure returns (bytes memory) {
        return poolKeyInputForWords(epochId, threshold, committeeSize, acceptedCount, keyIndex,
            poolKeyWords(keyIndex, acceptedCount, committeeSize));
    }

    /// @dev Same, over an arbitrary (possibly tampered) word vector so a test
    ///      can hand the contract a transcript whose challenge and BRLC
    ///      commitment are internally consistent and still watch the on-chain
    ///      row bindings reject it.
    function poolKeyInputForWords(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 acceptedCount,
        uint8 keyIndex,
        uint256[] memory words
    ) internal pure returns (bytes memory) {
        return poolKeyInputForDigest(
            epochId, threshold, committeeSize, acceptedCount, keyIndex, testTranscriptDigest(keyIndex), words
        );
    }

    /// @dev Same, with an explicit transcript digest. The challenge anchor is
    ///      `keccak(digest ‖ keccak(transcript))`, exactly as the contract
    ///      derives it.
    function poolKeyInputForDigest(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 acceptedCount,
        uint8 keyIndex,
        bytes32 transcriptDigest,
        uint256[] memory words
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            POOLKEY_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(transcriptDigest, keccak256(abi.encodePacked(words))))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(threshold),
                uint256(committeeSize),
                uint256(acceptedCount),
                uint256(keyIndex),
                uint256(transcriptDigest),
                challenge,
                BRLC.commit(challenge, words)
            ]
        );
    }

    // ─── Share-commitment Merkle tree ─────────────────────────────────────────

    /// @dev Tagged hashes, recomputed here rather than imported so the test
    ///      does not borrow the contract's own arithmetic to check the
    ///      contract. Leaves carry tag 0x00, internal nodes tag 0x01.
    bytes32 internal constant TEST_MERKLE_EMPTY_LEAF = keccak256("davinci-dkg:merkle-empty:v1");

    function leafHash(uint256 x, uint256 y) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint8(0), x, y));
    }

    function nodeHash(bytes32 left, bytes32 right) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint8(1), left, right));
    }

    /// @dev The `MAX_N` leaves of pool key `keyIndex`: the tagged leaf of
    ///      `D_{i+1}` for every committee slot `i < committeeSize`, the empty
    ///      leaf everywhere else. Mirrors `DKGManager._verifyPoolKeyRows`.
    function shareLeaves(uint8 keyIndex, uint16 committeeSize)
        internal
        pure
        returns (bytes32[MAX_N] memory leaves)
    {
        for (uint256 i = 0; i < MAX_N; i++) {
            if (i < committeeSize) {
                (uint256 x, uint256 y) = testShareCommitment(keyIndex, uint16(i + 1));
                leaves[i] = leafHash(x, y);
            } else {
                leaves[i] = TEST_MERKLE_EMPTY_LEAF;
            }
        }
    }

    /// @dev Independently recomputed root.
    function expectedShareRoot(uint8 keyIndex, uint16 committeeSize) internal pure returns (bytes32) {
        bytes32[MAX_N] memory leaves = shareLeaves(keyIndex, committeeSize);
        uint256 width = MAX_N;
        for (uint256 level = 0; level < MERKLE_DEPTH; level++) {
            width >>= 1;
            for (uint256 i = 0; i < width; i++) {
                leaves[i] = nodeHash(leaves[i * 2], leaves[i * 2 + 1]);
            }
        }
        return leaves[0];
    }

    /// @dev `MERKLE_DEPTH` siblings bottom-up for leaf `participantIndex - 1`.
    function shareProofFor(uint8 keyIndex, uint16 committeeSize, uint16 participantIndex)
        internal
        pure
        returns (bytes32[] memory)
    {
        return shareProofWithNodeTag(keyIndex, committeeSize, participantIndex, 1);
    }

    /// @dev Same path, but with the internal nodes folded under `nodeTag`
    ///      instead of the canonical 0x01. Only a tag of 1 yields a path the
    ///      contract accepts; a test hands it 0 to show that a leaf hash can
    ///      never stand in for an internal node.
    function shareProofWithNodeTag(uint8 keyIndex, uint16 committeeSize, uint16 participantIndex, uint8 nodeTag)
        internal
        pure
        returns (bytes32[] memory path)
    {
        bytes32[MAX_N] memory leaves = shareLeaves(keyIndex, committeeSize);
        path = new bytes32[](MERKLE_DEPTH);
        uint256 index = uint256(participantIndex) - 1;
        uint256 width = MAX_N;
        for (uint256 level = 0; level < MERKLE_DEPTH; level++) {
            path[level] = leaves[index ^ 1];
            width >>= 1;
            for (uint256 i = 0; i < width; i++) {
                leaves[i] = keccak256(abi.encodePacked(nodeTag, leaves[i * 2], leaves[i * 2 + 1]));
            }
            index >>= 1;
        }
    }

    // ─── Partial decryption ───────────────────────────────────────────────────

    function partialDecryptionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(11), 12, 13, 14, 15, 16, 17, 18]);
    }

    /// partialdecrypt layout (15 public inputs):
    ///   [0] eid, [1] aid, [2] ctIdx, [3] participantIndex,
    ///   [4..5] C1.x/y, [6..7] D_i.x/y, [8..9] delta.x/y,
    ///   [10..11] A1.x/y, [12..13] A2.x/y, [14] response.
    function partialDecryptionInput(
        bytes12 epochId,
        bytes32 aid,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        uint8 keyIndex
    ) internal pure returns (bytes memory) {
        // pi[6..7] = the member's own D_i under the application's pool key —
        // the leaf whose Merkle path the contract checks.
        (uint256 dx, uint256 dy) = testShareCommitment(keyIndex, participantIndex);
        return partialDecryptionInputWithShare(epochId, aid, participantIndex, ciphertextIndex, dx, dy);
    }

    /// @dev Same, with an explicit share commitment in pi[6..7] so a test can
    ///      submit somebody else's `D` (or one from another key) at the
    ///      caller's index.
    function partialDecryptionInputWithShare(
        bytes12 epochId,
        bytes32 aid,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        uint256 shareX,
        uint256 shareY
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
        inputs[6] = shareX;
        inputs[7] = shareY;
        inputs[8] = 7000 + participantIndex;    // delta.x
        inputs[9] = 8000 + participantIndex;    // delta.y
        return abi.encode(inputs);
    }

    function partialDecryptionHash(uint16 participantIndex) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint256(7000 + participantIndex), uint256(8000 + participantIndex)));
    }

    // ─── Combine ──────────────────────────────────────────────────────────────

    /// @notice Everything `combineDecryption` needs for one ciphertext. The
    ///         organizer contributes nothing but its public key now: the
    ///         secret is a private witness of the circuit.
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
    }

    /// @dev Combine fixture for an `OrganizerLocked` application.
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
            pkOrgY: pky
        });
    }

    /// @dev Combine fixture for an `Automatic` application: `PK_org` is the
    ///      identity, so the transcript's organizer words are `(0, 1)`.
    function combineFixtureAutomatic(
        bytes12 epochId,
        bytes32 aid,
        uint16 ctIdx,
        uint16 threshold,
        uint16 shareCount,
        bytes32 combineHash,
        uint256 plaintext
    ) internal pure returns (CombineFixture memory f) {
        f = CombineFixture({
            epochId: epochId,
            aid: aid,
            ctIdx: ctIdx,
            threshold: threshold,
            shareCount: shareCount,
            combineHash: combineHash,
            plaintext: plaintext,
            pkOrgX: 0,
            pkOrgY: 1
        });
    }

    function decryptCombineProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(31), 32, 33, 34, 35, 36, 37, 38]);
    }

    /// @dev The `6 + 3N` transcript words of `combineDecryption`, in the exact
    ///      order the contract reads them out of calldata:
    ///        w[0..3]          C1.x C1.y C2.x C2.y
    ///        w[4..5]          PK_org.x PK_org.y
    ///        w[6 .. 6+N)      participant indexes (0 when inactive)
    ///        w[6+N .. 6+3N)   partial decryptions (identity when inactive)
    function combineWords(CombineFixture memory f) internal pure returns (uint256[] memory v) {
        v = new uint256[](COMBINE_TRANSCRIPT_WORDS);
        v[0] = TEST_CT_C1X;
        v[1] = TEST_CT_C1Y;
        v[2] = TEST_CT_C2X;
        v[3] = TEST_CT_C2Y;
        v[4] = f.pkOrgX;
        v[5] = f.pkOrgY;
        uint256 partialBase = 6 + MAX_N;
        for (uint256 i = 0; i < MAX_N; i++) {
            v[partialBase + i * 2 + 1] = 1; // identity padding
        }
        for (uint256 i = 0; i < f.shareCount; i++) {
            v[6 + i] = i + 1;
            v[partialBase + i * 2] = 7000 + i + 1;
            v[partialBase + i * 2 + 1] = 8000 + i + 1;
        }
    }

    function combineTranscript(CombineFixture memory f) internal pure returns (bytes memory) {
        return abi.encodePacked(combineWords(f));
    }

    /// @dev 9-element public-input vector matching the combine circuit and
    ///      verifier: eid, aid, ctIdx, threshold, shareCount, combineHash,
    ///      plaintext, challenge, transcriptCommitment.
    function combineInput(CombineFixture memory f) internal pure returns (bytes memory) {
        return combineInputForWords(f, combineWords(f));
    }

    /// @dev Same, but over an arbitrary (possibly tampered) word vector.
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
        uint256[9] memory v;
        v[0] = uint256(uint96(f.epochId));
        v[1] = uint256(f.aid);
        v[2] = uint256(f.ctIdx);
        v[3] = uint256(f.threshold);
        v[4] = uint256(f.shareCount);
        v[5] = uint256(f.combineHash);
        v[6] = f.plaintext;
        v[7] = challenge;
        v[8] = BRLC.commit(challenge, words);
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

contract MockPoolKeyVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return POOLKEY_PROVING_KEY_HASH;
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
