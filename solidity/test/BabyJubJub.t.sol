// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {BabyJubJub} from "../src/libraries/BabyJubJub.sol";

/// @title BabyJubJubTest
/// @notice Tests scalar multiplication and validation against canonical
///         test vectors generated from gnark's
///         `bn254/twistededwards.PointAffine.ScalarMultiplication`.
///
///         All test vectors below were produced by the Go script in
///         `/tmp/bjj_vectors.go` (see PR description); each vector is the
///         output of a known-good implementation, so a passing test
///         certifies that our reduced-form (a = -1) Solidity scalar mul
///         agrees with gnark byte-for-byte.
contract BabyJubJubTest is Test {
    // Canonical generator (also exposed by the library)
    uint256 internal constant GX =
        9671717474070082183213120605117400219616337014328744928644933853176787189663;
    uint256 internal constant GY =
        16950150798460657717958625567821834550301663161624707787222815936182638968203;
    // Subgroup order
    uint256 internal constant L =
        2736030358979909402780800718157159386076813972158567259200215660948447373041;

    // ─── scalarMul ────────────────────────────────────────────────────────────

    function test_ScalarMulBy0_ReturnsIdentity() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(0);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    function test_ScalarMulBy1_ReturnsGenerator() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(1);
        assertEq(x, GX);
        assertEq(y, GY);
    }

    function test_ScalarMulBy2_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(2);
        assertEq(x, 6181589843805936102166733432625702983249926793164243794170119559257043191516);
        assertEq(y, 633281375905621697187330766174974863687049529291089048651929454608812697683);
    }

    function test_ScalarMulBy7_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(7);
        assertEq(x, 19649055493706707054227543952293108083531865249967651845006107589639068627074);
        assertEq(y, 12112450042127193446189577552007703839818242727902437791835414514847797088033);
    }

    function test_ScalarMul_LargeScalar_MatchesGnark() public view {
        uint256 k =
            123456789012345678901234567890123456789012345678901234567890;
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(k);
        assertEq(x, 20191150976955993420137836425697324752060073403133934639955319916335856022496);
        assertEq(y, 428693552930091294818271783784035400470921813234790925968369531890244111764);
    }

    function test_ScalarMulByL_ReturnsIdentity() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(L);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    function test_ScalarMulByLPlus1_EqualsGenerator() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(L + 1);
        assertEq(x, GX);
        assertEq(y, GY);
    }

    /// @notice Scalar mul of an arbitrary-base point: 7 * (2*G) == 14 * G.
    function test_ScalarMul_NonBasePoint() public view {
        (uint256 px, uint256 py) = BabyJubJub.scalarMulBase(2);
        (uint256 x1, uint256 y1) = BabyJubJub.scalarMul(7, px, py);
        (uint256 x2, uint256 y2) = BabyJubJub.scalarMulBase(14);
        assertEq(x1, x2);
        assertEq(y1, y2);
    }

    // ─── pointAdd / pointDouble agreement ─────────────────────────────────────

    function test_PointDouble_AgreesWithPointAdd() public view {
        (uint256 x1, uint256 y1) = BabyJubJub.pointDouble(GX, GY);
        (uint256 x2, uint256 y2) = BabyJubJub.pointAdd(GX, GY, GX, GY);
        assertEq(x1, x2);
        assertEq(y1, y2);
    }

    // ─── isOnCurve ────────────────────────────────────────────────────────────

    function test_IsOnCurve_Generator() public pure {
        assertTrue(BabyJubJub.isOnCurve(GX, GY));
    }

    function test_IsOnCurve_Identity() public pure {
        assertTrue(BabyJubJub.isOnCurve(0, 1));
    }

    function test_IsOnCurve_Random() public pure {
        // Off-curve point: change one coordinate of G by 1.
        assertFalse(BabyJubJub.isOnCurve(GX + 1, GY));
        assertFalse(BabyJubJub.isOnCurve(GX, GY + 1));
    }

    function test_IsOnCurve_RejectsOutOfRange() public pure {
        // Coordinates >= Q must be rejected as not canonical (and trivially not on curve).
        assertFalse(BabyJubJub.isOnCurve(BabyJubJub.Q, GY));
        assertFalse(BabyJubJub.isOnCurve(GX, BabyJubJub.Q));
    }

    // ─── isCanonical / isIdentity ─────────────────────────────────────────────

    function test_IsCanonical() public pure {
        assertTrue(BabyJubJub.isCanonical(0, 1));
        assertTrue(BabyJubJub.isCanonical(GX, GY));
        assertFalse(BabyJubJub.isCanonical(BabyJubJub.Q, 0));
        assertFalse(BabyJubJub.isCanonical(0, BabyJubJub.Q));
    }

    function test_IsIdentity() public pure {
        assertTrue(BabyJubJub.isIdentity(0, 1));
        assertFalse(BabyJubJub.isIdentity(GX, GY));
        assertFalse(BabyJubJub.isIdentity(0, 2));
    }

    // ─── requireValidPoint ────────────────────────────────────────────────────

    function test_RequireValidPoint_AcceptsGenerator() public pure {
        BabyJubJub.requireValidPoint(GX, GY);
    }

    function test_RequireValidPoint_RejectsIdentity() public {
        vm.expectRevert(BabyJubJub.IsIdentity.selector);
        this.callRequireValidPoint(0, 1);
    }

    function test_RequireValidPoint_RejectsNonCanonical() public {
        vm.expectRevert(BabyJubJub.NotCanonical.selector);
        this.callRequireValidPoint(BabyJubJub.Q, GY);
    }

    function test_RequireValidPoint_RejectsOffCurve() public {
        vm.expectRevert(BabyJubJub.NotOnCurve.selector);
        this.callRequireValidPoint(GX + 1, GY);
    }

    /// @dev External wrapper for `vm.expectRevert` (must be an external call).
    function callRequireValidPoint(uint256 x, uint256 y) external pure {
        BabyJubJub.requireValidPoint(x, y);
    }

    // ─── isInPrimeSubgroup ────────────────────────────────────────────────────

    function test_IsInPrimeSubgroup_GeneratorIsInside() public view {
        assertTrue(BabyJubJub.isInPrimeSubgroup(GX, GY));
    }

    function test_IsInPrimeSubgroup_IdentityIsInside() public view {
        // [L]·O = O, trivially.
        assertTrue(BabyJubJub.isInPrimeSubgroup(0, 1));
    }

    function test_IsInPrimeSubgroup_KGIsInside() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(123456789);
        assertTrue(BabyJubJub.isInPrimeSubgroup(x, y));
    }
}
