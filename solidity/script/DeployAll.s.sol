// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {DKGManager} from "../src/DKGManager.sol";
import {DKGAppManager} from "../src/DKGAppManager.sol";
import {ContributionVerifier} from "../src/verifiers/ContributionVerifier.sol";
import {PoolKeyVerifier} from "../src/verifiers/PoolKeyVerifier.sol";
import {PartialDecryptVerifier} from "../src/verifiers/PartialDecryptVerifier.sol";
import {DecryptCombineVerifier} from "../src/verifiers/DecryptCombineVerifier.sol";

contract DeployAllScript is Script {
    /// Default inactivity window if `INACTIVITY_WINDOW` is not set in the
    /// environment: 50 400 blocks ≈ 7 days at 12-second block time. Local
    /// testnets with 2-second blocks will run 6× faster than real time,
    /// which is usually fine — override with the env var when it matters.
    uint256 internal constant DEFAULT_INACTIVITY_WINDOW = 50_400;

    /// Default phase budgets when the matching env vars are not set. Pass 0
    /// for any of them to fall back to the contract's compiled-in defaults
    /// (`Sizes.sol`). All values are in BLOCKS — scale per chain block time.
    /// At ~12s blocks: 100 / 25 / 25 / 5 ≈ 20 min epoch with 5/5/1 min
    /// preparation budget and ~9 min Service window.
    uint256 internal constant DEFAULT_EPOCH_DURATION_BLOCKS      = 100;
    uint256 internal constant DEFAULT_COMMITTEE_SELECTION_BLOCKS = 25;
    uint256 internal constant DEFAULT_KEY_ASSEMBLY_BLOCKS        = 25;
    uint256 internal constant DEFAULT_FINALIZE_GAP_BLOCKS        = 5;

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        uint32 chainId = uint32(vm.envUint("CHAIN_ID"));
        uint64 inactivityWindow =
            uint64(vm.envOr("INACTIVITY_WINDOW", DEFAULT_INACTIVITY_WINDOW));
        uint256 epochDurationBlocks =
            vm.envOr("EPOCH_DURATION_BLOCKS", DEFAULT_EPOCH_DURATION_BLOCKS);
        uint256 committeeSelectionBlocks =
            vm.envOr("COMMITTEE_SELECTION_BLOCKS", DEFAULT_COMMITTEE_SELECTION_BLOCKS);
        uint256 keyAssemblyBlocks =
            vm.envOr("KEY_ASSEMBLY_BLOCKS", DEFAULT_KEY_ASSEMBLY_BLOCKS);
        uint256 finalizeGapBlocks =
            vm.envOr("FINALIZE_GAP_BLOCKS", DEFAULT_FINALIZE_GAP_BLOCKS);

        vm.startBroadcast(deployerPrivateKey);

        ContributionVerifier contributionVerifier = new ContributionVerifier();
        console.log("ContributionVerifier deployed at:", address(contributionVerifier));

        PoolKeyVerifier poolKeyVerifier = new PoolKeyVerifier();
        console.log("PoolKeyVerifier deployed at:", address(poolKeyVerifier));

        PartialDecryptVerifier partialDecryptVerifier = new PartialDecryptVerifier();
        console.log("PartialDecryptVerifier deployed at:", address(partialDecryptVerifier));

        DecryptCombineVerifier decryptCombineVerifier = new DecryptCombineVerifier();
        console.log("DecryptCombineVerifier deployed at:", address(decryptCombineVerifier));

        DKGRegistry registry = new DKGRegistry(inactivityWindow);
        console.log("DKGRegistry deployed at:", address(registry));
        console.log("DKGRegistry inactivityWindow:", inactivityWindow);

        DKGManager manager = new DKGManager(
            chainId,
            address(registry),
            address(contributionVerifier),
            address(partialDecryptVerifier),
            address(poolKeyVerifier),
            address(decryptCombineVerifier),
            epochDurationBlocks,
            committeeSelectionBlocks,
            keyAssemblyBlocks,
            finalizeGapBlocks,
            uint16(vm.envOr("MIN_THRESHOLD", uint256(1))),
            uint16(vm.envOr("MIN_COMMITTEE_SIZE", uint256(1))),
            uint16(vm.envOr("MAX_LOTTERY_ALPHA_BPS", uint256(65535)))
        );
        console.log("DKGManager deployed at:", address(manager));
        console.log("DKGManager epochDurationBlocks:", manager.EPOCH_DURATION_BLOCKS());
        console.log("DKGManager committeeSelectionBlocks:", committeeSelectionBlocks);
        console.log("DKGManager keyAssemblyBlocks:       ", keyAssemblyBlocks);
        console.log("DKGManager finalizeGapBlocks:       ", finalizeGapBlocks);

        // Wire the one-shot link from registry → manager so the latter can
        // call registry.markActive(...) from submitContribution.
        registry.setManager(address(manager));
        console.log("DKGRegistry.setManager:", address(manager));

        // Deploy the sibling app manager (per-application surface). It only
        // needs the manager address (cyclic dependency resolved by
        // setAppManager afterwards); the organizer secret is a private
        // witness of the combine circuit, so no verifier is wired here.
        DKGAppManager appManager = new DKGAppManager(address(manager));
        console.log("DKGAppManager deployed at:", address(appManager));

        manager.setAppManager(address(appManager));
        console.log("DKGManager.setAppManager:", address(appManager));

        vm.stopBroadcast();
    }
}
