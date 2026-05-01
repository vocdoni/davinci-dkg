// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {DKGManager} from "../src/DKGManager.sol";
import {DKGAppManager} from "../src/DKGAppManager.sol";
import {ContributionVerifier} from "../src/verifiers/ContributionVerifier.sol";
import {FinalizeVerifier} from "../src/verifiers/FinalizeVerifier.sol";
import {PartialDecryptVerifier} from "../src/verifiers/PartialDecryptVerifier.sol";
import {DecryptCombineVerifier} from "../src/verifiers/DecryptCombineVerifier.sol";

contract DeployAllScript is Script {
    /// Default inactivity window if `INACTIVITY_WINDOW` is not set in the
    /// environment: 50 400 blocks ≈ 7 days at 12-second block time. Local
    /// testnets with 2-second blocks will run 6× faster than real time,
    /// which is usually fine — override with the env var when it matters.
    uint256 internal constant DEFAULT_INACTIVITY_WINDOW = 50_400;

    /// Default epoch duration if `EPOCH_DURATION_BLOCKS` is not set in the
    /// environment: 100 blocks ≈ 20 minutes at 12-second block time. Pass 0
    /// to fall back to the contract's compiled-in default (also 100). For
    /// faster chains (e.g. Polygon zkEVM at ~2s) bump this to ≈600 to
    /// preserve the 20-minute target.
    uint256 internal constant DEFAULT_EPOCH_DURATION_BLOCKS = 100;

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        uint32 chainId = uint32(vm.envUint("CHAIN_ID"));
        uint64 inactivityWindow =
            uint64(vm.envOr("INACTIVITY_WINDOW", DEFAULT_INACTIVITY_WINDOW));
        uint256 epochDurationBlocks =
            vm.envOr("EPOCH_DURATION_BLOCKS", DEFAULT_EPOCH_DURATION_BLOCKS);

        vm.startBroadcast(deployerPrivateKey);

        ContributionVerifier contributionVerifier = new ContributionVerifier();
        console.log("ContributionVerifier deployed at:", address(contributionVerifier));

        FinalizeVerifier finalizeVerifier = new FinalizeVerifier();
        console.log("FinalizeVerifier deployed at:", address(finalizeVerifier));

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
            address(finalizeVerifier),
            address(decryptCombineVerifier),
            epochDurationBlocks
        );
        console.log("DKGManager deployed at:", address(manager));
        console.log("DKGManager epochDurationBlocks:", manager.EPOCH_DURATION_BLOCKS());

        // Wire the one-shot link from registry → manager so the latter can
        // call registry.markActive(...) from submitContribution.
        registry.setManager(address(manager));
        console.log("DKGRegistry.setManager:", address(manager));

        // Deploy the sibling app manager (per-application surface). It needs
        // the manager address (cyclic dependency resolved by setAppManager
        // afterwards) and the same partial-decrypt verifier the manager uses
        // for the organizer DLEQ proof.
        DKGAppManager appManager = new DKGAppManager(
            address(manager),
            address(partialDecryptVerifier)
        );
        console.log("DKGAppManager deployed at:", address(appManager));

        manager.setAppManager(address(appManager));
        console.log("DKGManager.setAppManager:", address(appManager));

        vm.stopBroadcast();
    }
}
