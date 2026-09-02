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
///
///         The vectors in the "extended-coordinate edge cases" section were
///         generated with the same gnark-crypto curve (`GetEdwardsCurve()`,
///         v0.19.x) and cross-checked against the previous affine
///         implementation of this library before the switch to
///         inversion-free extended coordinates, so they pin both the
///         gnark encoding and the pre-existing on-chain behaviour.
contract BabyJubJubTest is Test {
    /// @dev DSTest-compatible log event (decoded by forge by signature).
    event log_named_uint(string key, uint256 val);

    // Canonical generator (also exposed by the library)
    uint256 internal constant GX =
        9671717474070082183213120605117400219616337014328744928644933853176787189663;
    uint256 internal constant GY =
        16950150798460657717958625567821834550301663161624707787222815936182638968203;
    // Subgroup order
    uint256 internal constant L =
        2736030358979909402780800718157159386076813972158567259200215660948447373041;
    uint256 internal constant Q =
        21888242871839275222246405745257275088548364400416034343698204186575808495617;

    // ─── Known multiples of G (gnark) ─────────────────────────────────────────
    uint256 internal constant G2X =
        6181589843805936102166733432625702983249926793164243794170119559257043191516;
    uint256 internal constant G2Y =
        633281375905621697187330766174974863687049529291089048651929454608812697683;
    uint256 internal constant G3X =
        8787197234602503388844295997017788867392848147420327174860413988187367696532;
    uint256 internal constant G3Y =
        15305195750036305661220525648961313310481046260814497672243197092298550508693;
    uint256 internal constant G5X =
        3733495933403540068220568802612039083442073215153629701861059777108116389351;
    uint256 internal constant G5Y =
        15148236048131954717802795400425086368006776860859772698778589175317365693546;
    uint256 internal constant G11X =
        20116430324039290829419172776146199614469768477688294126172321412365539407642;
    uint256 internal constant G11Y =
        18856460861531942120859708048677603751294231190189224157283439874962410808705;
    uint256 internal constant G16X =
        5261822793729097469124322713944452436263585332274847136083146132068833612219;
    uint256 internal constant G16Y =
        21459189231378695508316163458360356529222201254620325044724979975334648070151;
    // -5G = (Q - G5X, G5Y)
    uint256 internal constant NEG_G5X =
        18154746938435735154025836942645236005106291185262404641837144409467692106266;
    // (L-1)·G = -G = (Q - GX, GY)
    uint256 internal constant NEG_GX =
        12216525397769193039033285140139874868932027386087289415053270333399021305954;

    // ─── Small-order (non prime-subgroup) points ──────────────────────────────
    /// @dev T2 = (0, -1): the unique point of order 2.
    uint256 internal constant T2X = 0;
    uint256 internal constant T2Y = Q - 1;
    /// @dev T4 = (sqrt(-1), 0): a point of order 4, 2·T4 = T2.
    uint256 internal constant T4X =
        21888242871839275217838484774961031246007050428528088939761107053157389710902;
    uint256 internal constant T4Y = 0;
    /// @dev T8: a point of order 8 (4·T8 = T2), obtained as [L]·R for the
    ///      on-curve point R = (4, y).
    uint256 internal constant T8X =
        21449313544428428285896623807777999089618640778178766447151795324822652860828;
    uint256 internal constant T8Y =
        17061719626832259898845741003733890968968767993363194771977168648564009544074;
    /// @dev G + T2 = (-GX, -GY): order 2L, on-curve, not in the prime subgroup.
    uint256 internal constant GT2X = NEG_GX;
    uint256 internal constant GT2Y =
        4938092073378617504287780177435440538246701238791326556475388250393169527414;
    /// @dev G + T8: order 8L, on-curve, not in the prime subgroup.
    uint256 internal constant GT8X =
        6427738952372247604518126340853389786697880434634009228104512850016410032069;
    uint256 internal constant GT8Y =
        15621640772646418583778590390256985303553439950620095229470902472999933216128;

    // ─── Synthetic Schnorr tuples: PK = s·G, A = r·G, z = r + c·s mod L ──────
    uint256 internal constant SCH_C =
        1234567890123456789012345678901234567890123456789012345678901234567890123;
    // s = 1 (PK = G — exercises the P + (-P) table entries), r = 987654322
    uint256 internal constant SCH1_Z =
        1234567890123456789012345678901234567890123456789012345678901235555544445;
    uint256 internal constant SCH1_AX =
        15929830136830197529779019716088774779100493081381449331652892757991648883651;
    uint256 internal constant SCH1_AY =
        10802615503362450110568798039436313326653683188035516274479660209240505613958;
    // s = 2 (PK = 2G), r = 987654323
    uint256 internal constant SCH2_Z =
        2469135780246913578024691357802469135780246913578024691357802470123434569;
    uint256 internal constant SCH2_AX =
        4850750824785483425776636665937266092688952658211784375515132518680319764268;
    uint256 internal constant SCH2_AY =
        11960148944620212842292215788462436617981251286745497574114562733939416465274;
    // s = 3 (PK = 3G), r = 987654324
    uint256 internal constant SCH3_Z =
        3703703670370370367037037036703703703670370370367037037036703704691324693;
    uint256 internal constant SCH3_AX =
        1317593406099134823760069169909273255367731763836547883377336344002192201875;
    uint256 internal constant SCH3_AY =
        14424631530722569806116988231285080084451322989477432120953043703021667006897;
    // s = 123456789, r = 1111111110
    uint256 internal constant SCH4_Z =
        2480339811955560407394069508747126722147655686566937867957859093308838884211;
    uint256 internal constant SCH4_PKX =
        16206784187338351803948532248314919174094404850951278235115538319641171317744;
    uint256 internal constant SCH4_PKY =
        1645780246786685895560641778865228215443840970280597910012614014295481144366;
    uint256 internal constant SCH4_AX =
        1541713900920503958512688235665803645371063041892924275585117477124544604846;
    uint256 internal constant SCH4_AY =
        19437540052186786462856119590384329059774438818151597515788174287588703945599;

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

    function test_IsInPrimeSubgroup_GeneratorIsInside() public pure {
        assertTrue(BabyJubJub.isInPrimeSubgroup(GX, GY));
    }

    function test_IsInPrimeSubgroup_IdentityIsInside() public pure {
        // [L]·O = O, trivially.
        assertTrue(BabyJubJub.isInPrimeSubgroup(0, 1));
    }

    function test_IsInPrimeSubgroup_KGIsInside() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(123456789);
        assertTrue(BabyJubJub.isInPrimeSubgroup(x, y));
    }

    /// @dev (0, Q-1) is the order-2 point on this twisted-Edwards curve
    ///      (a=-1, identity = (0, 1)). It satisfies the curve equation
    ///      but lies outside the prime-order subgroup, so a correct
    ///      `isInPrimeSubgroup` must reject it.
    function test_IsInPrimeSubgroup_RejectsOrderTwoPoint() public pure {
        uint256 q = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
        uint256 x = 0;
        uint256 y = q - 1;
        assertTrue(BabyJubJub.isOnCurve(x, y));
        assertTrue(!BabyJubJub.isInPrimeSubgroup(x, y));
    }

    // ═══════════════════════════════════════════════════════════════════════════
    // Extended-coordinate edge cases (gnark vectors, cross-checked against the
    // previous affine implementation).
    // ═══════════════════════════════════════════════════════════════════════════

    // ─── scalarMul: more known multiples ──────────────────────────────────────

    function test_ScalarMulBy3_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(3);
        assertEq(x, G3X);
        assertEq(y, G3Y);
    }

    function test_ScalarMulBy5_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(5);
        assertEq(x, G5X);
        assertEq(y, G5Y);
    }

    function test_ScalarMulBy11_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(11);
        assertEq(x, G11X);
        assertEq(y, G11Y);
    }

    function test_ScalarMulBy16_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(16);
        assertEq(x, G16X);
        assertEq(y, G16Y);
    }

    /// @notice [L-1]·G = -G = (Q - GX, GY): exercises every 2-bit window of a
    ///         full-length scalar.
    function test_ScalarMulByLMinus1_IsNegG() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(L - 1);
        assertEq(x, NEG_GX);
        assertEq(y, GY);
        assertEq(x, Q - GX);
    }

    /// @notice [L-1]·(5G) = -5G on a non-generator base.
    function test_ScalarMul_LMinus1_NonBasePoint() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMul(L - 1, G5X, G5Y);
        assertEq(x, NEG_G5X);
        assertEq(y, G5Y);
    }

    /// @notice [L]·P = O for a non-generator subgroup point.
    function test_ScalarMul_ByL_NonBasePoint_IsIdentity() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMul(L, G5X, G5Y);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    /// @notice Scalars are reduced mod L: [2L + 3]·G == 3G.
    function test_ScalarMul_ReducesModL() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMulBase(2 * L + 3);
        assertEq(x, G3X);
        assertEq(y, G3Y);
    }

    /// @notice Every 2-bit window value (0, 1, 2, 3) with a mix of leading
    ///         zeros: [0b11_10_01_00_11]·G == 0x3C3·G == 963·G, checked against
    ///         the additive decomposition 963 = 16·60 + 3.
    function test_ScalarMul_AllWindowValues() public view {
        (uint256 x1, uint256 y1) = BabyJubJub.scalarMulBase(963);
        (uint256 ax, uint256 ay) = BabyJubJub.scalarMul(60, G16X, G16Y);
        (uint256 x2, uint256 y2) = BabyJubJub.pointAdd(ax, ay, G3X, G3Y);
        assertEq(x1, x2);
        assertEq(y1, y2);
    }

    /// @notice Scalar multiplication of the identity is the identity.
    function test_ScalarMul_OfIdentity_IsIdentity() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMul(12345, 0, 1);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    // ─── pointAdd: completeness edge cases ────────────────────────────────────

    function test_PointAdd_GPlusG_Is2G() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(GX, GY, GX, GY);
        assertEq(x, G2X);
        assertEq(y, G2Y);
    }

    function test_PointAdd_2GPlusG_Is3G() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(G2X, G2Y, GX, GY);
        assertEq(x, G3X);
        assertEq(y, G3Y);
    }

    function test_PointAdd_5GPlus11G_Is16G() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(G5X, G5Y, G11X, G11Y);
        assertEq(x, G16X);
        assertEq(y, G16Y);
    }

    function test_PointAdd_IsCommutative() public view {
        (uint256 x1, uint256 y1) = BabyJubJub.pointAdd(G5X, G5Y, G11X, G11Y);
        (uint256 x2, uint256 y2) = BabyJubJub.pointAdd(G11X, G11Y, G5X, G5Y);
        assertEq(x1, x2);
        assertEq(y1, y2);
    }

    function test_PointAdd_IsAssociative() public view {
        (uint256 lx, uint256 ly) = BabyJubJub.pointAdd(GX, GY, G2X, G2Y); // 3G
        (lx, ly) = BabyJubJub.pointAdd(lx, ly, G5X, G5Y); // 8G
        (uint256 rx, uint256 ry) = BabyJubJub.pointAdd(G2X, G2Y, G5X, G5Y); // 7G
        (rx, ry) = BabyJubJub.pointAdd(GX, GY, rx, ry); // 8G
        (uint256 ex, uint256 ey) = BabyJubJub.scalarMulBase(8);
        assertEq(lx, rx);
        assertEq(ly, ry);
        assertEq(lx, ex);
        assertEq(ly, ey);
    }

    /// @notice P + (-P) = O — the case that breaks non-complete formulas.
    function test_PointAdd_PPlusNegP_IsIdentity() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(G5X, G5Y, NEG_G5X, G5Y);
        assertEq(x, 0);
        assertEq(y, 1);
        (x, y) = BabyJubJub.pointAdd(GX, GY, NEG_GX, GY);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    function test_PointAdd_IdentityIsNeutral() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(0, 1, G5X, G5Y);
        assertEq(x, G5X);
        assertEq(y, G5Y);
        (x, y) = BabyJubJub.pointAdd(G5X, G5Y, 0, 1);
        assertEq(x, G5X);
        assertEq(y, G5Y);
        (x, y) = BabyJubJub.pointAdd(0, 1, 0, 1);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    function test_PointDouble_Identity_IsIdentity() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointDouble(0, 1);
        assertEq(x, 0);
        assertEq(y, 1);
    }

    function test_PointDouble_2G_Is4G_MatchesScalarMul() public view {
        (uint256 x1, uint256 y1) = BabyJubJub.pointDouble(G2X, G2Y);
        (uint256 x2, uint256 y2) = BabyJubJub.scalarMulBase(4);
        assertEq(x1, x2);
        assertEq(y1, y2);
    }

    /// @notice Doubling the order-2 point gives the identity, doubling an
    ///         order-4 point gives the order-2 point: the formulas must be
    ///         complete on the full curve, not just the prime subgroup.
    function test_PointDouble_SmallOrderPoints() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointDouble(T2X, T2Y);
        assertEq(x, 0);
        assertEq(y, 1);
        (x, y) = BabyJubJub.pointAdd(T2X, T2Y, T2X, T2Y);
        assertEq(x, 0);
        assertEq(y, 1);
        (x, y) = BabyJubJub.pointDouble(T4X, T4Y);
        assertEq(x, T2X);
        assertEq(y, T2Y);
        (x, y) = BabyJubJub.pointDouble(T8X, T8Y);
        (x, y) = BabyJubJub.pointDouble(x, y);
        assertEq(x, T2X);
        assertEq(y, T2Y);
    }

    function test_PointAdd_GPlusT2_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(GX, GY, T2X, T2Y);
        assertEq(x, GT2X);
        assertEq(y, GT2Y);
        // G + T2 = (-GX, -GY)
        assertEq(x, Q - GX);
        assertEq(y, Q - GY);
    }

    function test_PointAdd_GPlusT8_MatchesGnark() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(GX, GY, T8X, T8Y);
        assertEq(x, GT8X);
        assertEq(y, GT8Y);
    }

    /// @notice Results are always canonical and on-curve.
    function test_PointAdd_ResultIsCanonicalOnCurve() public view {
        (uint256 x, uint256 y) = BabyJubJub.pointAdd(G5X, G5Y, G11X, G11Y);
        assertTrue(BabyJubJub.isCanonical(x, y));
        assertTrue(BabyJubJub.isOnCurve(x, y));
        (x, y) = BabyJubJub.scalarMulBase(L - 2);
        assertTrue(BabyJubJub.isCanonical(x, y));
        assertTrue(BabyJubJub.isOnCurve(x, y));
    }

    // ─── Constants ────────────────────────────────────────────────────────────

    /// @notice The a = -1 extended addition formula uses k = 2·d.
    function test_TwoDConstant() public pure {
        assertEq(BabyJubJub.TWO_D, mulmod(2, BabyJubJub.D, BabyJubJub.Q));
        assertEq(BabyJubJub.A, BabyJubJub.Q - 1);
    }

    /// @notice Completeness precondition of the unified addition law
    ///         (Bernstein et al. 2008, Thm 3.3): a must be a square and d a
    ///         non-square in F_Q. Checked via Euler's criterion.
    function test_CompletenessPrecondition_DIsNonSquare() public view {
        uint256 e = (BabyJubJub.Q - 1) / 2;
        assertEq(_powmod(BabyJubJub.D, e), BabyJubJub.Q - 1); // d^((Q-1)/2) = -1
        assertEq(_powmod(BabyJubJub.A, e), 1); // a = -1 is a square (Q = 1 mod 4)
    }

    function _powmod(uint256 base, uint256 exp) internal view returns (uint256 out) {
        bytes memory input = abi.encode(uint256(32), uint256(32), uint256(32), base, exp, BabyJubJub.Q);
        (bool ok, bytes memory ret) = address(0x05).staticcall(input);
        assertTrue(ok);
        out = abi.decode(ret, (uint256));
    }

    // ─── isInPrimeSubgroup: full-curve torsion ────────────────────────────────

    function test_IsInPrimeSubgroup_NegGIsInside() public pure {
        assertTrue(BabyJubJub.isInPrimeSubgroup(NEG_GX, GY));
    }

    function test_IsInPrimeSubgroup_RejectsOrderFourPoint() public pure {
        assertTrue(BabyJubJub.isOnCurve(T4X, T4Y));
        assertFalse(BabyJubJub.isInPrimeSubgroup(T4X, T4Y));
    }

    function test_IsInPrimeSubgroup_RejectsOrderEightPoint() public pure {
        assertTrue(BabyJubJub.isOnCurve(T8X, T8Y));
        assertFalse(BabyJubJub.isInPrimeSubgroup(T8X, T8Y));
    }

    /// @notice G + T2 has order 2L: it is on-curve and [8]·(G + T2) is in the
    ///         subgroup, but the point itself is not.
    function test_IsInPrimeSubgroup_RejectsGPlusT2() public pure {
        assertTrue(BabyJubJub.isOnCurve(GT2X, GT2Y));
        assertFalse(BabyJubJub.isInPrimeSubgroup(GT2X, GT2Y));
    }

    function test_IsInPrimeSubgroup_RejectsGPlusT8() public pure {
        assertTrue(BabyJubJub.isOnCurve(GT8X, GT8Y));
        assertFalse(BabyJubJub.isInPrimeSubgroup(GT8X, GT8Y));
    }

    /// @notice Clearing the cofactor of G + T8 lands back in the subgroup.
    function test_IsInPrimeSubgroup_CofactorClearedIsInside() public view {
        (uint256 x, uint256 y) = BabyJubJub.scalarMul(8, GT8X, GT8Y);
        assertTrue(BabyJubJub.isInPrimeSubgroup(x, y));
    }

    /// @notice Off-curve and non-canonical inputs are never "in the subgroup".
    function test_IsInPrimeSubgroup_RejectsOffCurveAndNonCanonical() public pure {
        assertFalse(BabyJubJub.isInPrimeSubgroup(GX + 1, GY));
        assertFalse(BabyJubJub.isInPrimeSubgroup(GX, GY + 1));
        assertFalse(BabyJubJub.isInPrimeSubgroup(GX + Q, GY));
        assertFalse(BabyJubJub.isInPrimeSubgroup(GX, GY + Q));
    }

    // ─── verifySchnorrEquation ────────────────────────────────────────────────

    function test_VerifySchnorr_AcceptsValid_PKIsG() public pure {
        // PK = G: the table entries i·G + j·(-PK) hit P + (-P) = O and
        // its multiples, so every degenerate addition case is exercised.
        assertTrue(BabyJubJub.verifySchnorrEquation(SCH1_Z, SCH_C, SCH1_AX, SCH1_AY, GX, GY));
    }

    function test_VerifySchnorr_AcceptsValid_PKIs2G() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(SCH2_Z, SCH_C, SCH2_AX, SCH2_AY, G2X, G2Y));
    }

    function test_VerifySchnorr_AcceptsValid_PKIs3G() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(SCH3_Z, SCH_C, SCH3_AX, SCH3_AY, G3X, G3Y));
    }

    function test_VerifySchnorr_AcceptsValid_LargeSecret() public pure {
        assertTrue(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY)
        );
    }

    function test_VerifySchnorr_ScalarsReducedModL() public pure {
        assertTrue(
            BabyJubJub.verifySchnorrEquation(SCH4_Z + L, SCH_C + L, SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY)
        );
    }

    function test_VerifySchnorr_RejectsTamperedZ() public pure {
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z + 1, SCH_C, SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY)
        );
    }

    function test_VerifySchnorr_RejectsTamperedC() public pure {
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C + 1, SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY)
        );
    }

    function test_VerifySchnorr_RejectsWrongA() public pure {
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH3_AX, SCH3_AY, SCH4_PKX, SCH4_PKY)
        );
    }

    function test_VerifySchnorr_RejectsWrongPK() public pure {
        assertFalse(BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX, SCH4_AY, G2X, G2Y));
    }

    function test_VerifySchnorr_RejectsNegatedA() public pure {
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, Q - SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY)
        );
    }

    /// @notice z = c = 0: 0·G == A + 0·PK holds iff A is the identity.
    function test_VerifySchnorr_ZeroScalars() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(0, 0, 0, 1, SCH4_PKX, SCH4_PKY));
        assertFalse(BabyJubJub.verifySchnorrEquation(0, 0, GX, GY, SCH4_PKX, SCH4_PKY));
        assertTrue(BabyJubJub.verifySchnorrEquation(L, L, 0, 1, SCH4_PKX, SCH4_PKY));
    }

    /// @notice c = 0: z·G == A, independent of PK.
    function test_VerifySchnorr_ZeroChallenge() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(5, 0, G5X, G5Y, SCH4_PKX, SCH4_PKY));
        assertTrue(BabyJubJub.verifySchnorrEquation(16, 0, G16X, G16Y, GX, GY));
        assertFalse(BabyJubJub.verifySchnorrEquation(5, 0, G11X, G11Y, SCH4_PKX, SCH4_PKY));
    }

    /// @notice z = 0: O == A + c·PK holds iff A = -c·PK.
    function test_VerifySchnorr_ZeroResponse() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(0, 5, NEG_GX, GY, G5X, G5Y) == false);
        // A = -(1·5G) = -5G, c = 1, PK = 5G.
        assertTrue(BabyJubJub.verifySchnorrEquation(0, 1, NEG_G5X, G5Y, G5X, G5Y));
    }

    /// @notice PK = identity: the equation degenerates to z·G == A.
    function test_VerifySchnorr_IdentityPK() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(5, SCH_C, G5X, G5Y, 0, 1));
        assertFalse(BabyJubJub.verifySchnorrEquation(5, SCH_C, G11X, G11Y, 0, 1));
    }

    /// @notice A = identity with non-trivial scalars: z·G == c·PK.
    ///         With PK = G, z = c: true. With z = c + 1: false.
    function test_VerifySchnorr_IdentityA() public pure {
        assertTrue(BabyJubJub.verifySchnorrEquation(SCH_C, SCH_C, 0, 1, GX, GY));
        assertFalse(BabyJubJub.verifySchnorrEquation(SCH_C + 1, SCH_C, 0, 1, GX, GY));
    }

    /// @notice Non-canonical A (coordinate >= Q) never verifies, even when the
    ///         reduced coordinates would match.
    function test_VerifySchnorr_RejectsNonCanonicalA() public pure {
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX + Q, SCH4_AY, SCH4_PKX, SCH4_PKY)
        );
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX, SCH4_AY + Q, SCH4_PKX, SCH4_PKY)
        );
        assertFalse(BabyJubJub.verifySchnorrEquation(0, 0, Q, 1, SCH4_PKX, SCH4_PKY));
        assertFalse(BabyJubJub.verifySchnorrEquation(0, 0, 0, Q + 1, SCH4_PKX, SCH4_PKY));
    }

    /// @notice Non-canonical PK (coordinate >= Q) never verifies: (Q, 1) must
    ///         not be treated as the identity, and (Q + x, y) must not be
    ///         treated as (x, y).
    function test_VerifySchnorr_RejectsNonCanonicalPK() public pure {
        assertFalse(BabyJubJub.verifySchnorrEquation(5, SCH_C, G5X, G5Y, Q, 1));
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX, SCH4_AY, SCH4_PKX + Q, SCH4_PKY)
        );
        assertFalse(
            BabyJubJub.verifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY + Q)
        );
    }

    /// @notice Fuzz: for random secret s and nonce r, the honestly built
    ///         tuple verifies and the tuple with a tampered response does not.
    function testFuzz_VerifySchnorr_Roundtrip(uint256 s, uint256 r, uint256 c) public view {
        s = 1 + (s % (L - 1));
        r = 1 + (r % (L - 1));
        c = c % L;
        (uint256 pkx, uint256 pky) = BabyJubJub.scalarMulBase(s);
        (uint256 ax, uint256 ay) = BabyJubJub.scalarMulBase(r);
        uint256 z = addmod(r, mulmod(c, s, L), L);
        assertTrue(BabyJubJub.verifySchnorrEquation(z, c, ax, ay, pkx, pky));
        assertFalse(BabyJubJub.verifySchnorrEquation(addmod(z, 1, L), c, ax, ay, pkx, pky));
    }

    /// @notice Fuzz: [a]·G + [b]·G == [a + b]·G and [a]·([b]·G) == [a·b]·G.
    function testFuzz_ScalarMul_Homomorphic(uint256 a, uint256 b) public view {
        a = a % L;
        b = b % L;
        (uint256 ax, uint256 ay) = BabyJubJub.scalarMulBase(a);
        (uint256 bx, uint256 by) = BabyJubJub.scalarMulBase(b);
        (uint256 sx, uint256 sy) = BabyJubJub.pointAdd(ax, ay, bx, by);
        (uint256 ex, uint256 ey) = BabyJubJub.scalarMulBase(addmod(a, b, L));
        assertEq(sx, ex);
        assertEq(sy, ey);
        (uint256 px, uint256 py) = BabyJubJub.scalarMul(a, bx, by);
        (uint256 qx, uint256 qy) = BabyJubJub.scalarMulBase(mulmod(a, b, L));
        assertEq(px, qx);
        assertEq(py, qy);
    }

    // ─── Gas measurements (reported with -vv and in --gas-report) ─────────────

    function callVerifySchnorrEquation(
        uint256 z,
        uint256 c,
        uint256 ax,
        uint256 ay,
        uint256 pkx,
        uint256 pky
    ) external pure returns (bool) {
        return BabyJubJub.verifySchnorrEquation(z, c, ax, ay, pkx, pky);
    }

    function callIsInPrimeSubgroup(uint256 x, uint256 y) external pure returns (bool) {
        return BabyJubJub.isInPrimeSubgroup(x, y);
    }

    function callScalarMulBase(uint256 s) external view returns (uint256, uint256) {
        return BabyJubJub.scalarMulBase(s);
    }


    function test_Gas_VerifySchnorrEquation() public {
        uint256 g0 = gasleft();
        bool ok = this.callVerifySchnorrEquation(SCH4_Z, SCH_C, SCH4_AX, SCH4_AY, SCH4_PKX, SCH4_PKY);
        uint256 used = g0 - gasleft();
        assertTrue(ok);
        emit log_named_uint("gas verifySchnorrEquation (external wrapper)", used);
    }

    function test_Gas_IsInPrimeSubgroup() public {
        uint256 g0 = gasleft();
        bool ok = this.callIsInPrimeSubgroup(SCH4_PKX, SCH4_PKY);
        uint256 used = g0 - gasleft();
        assertTrue(ok);
        emit log_named_uint("gas isInPrimeSubgroup (external wrapper)", used);
    }

    function test_Gas_ScalarMulBase() public {
        uint256 g0 = gasleft();
        (uint256 x,) = this.callScalarMulBase(L - 1);
        uint256 used = g0 - gasleft();
        assertEq(x, NEG_GX);
        emit log_named_uint("gas scalarMulBase(L-1) (external wrapper)", used);
    }
}
