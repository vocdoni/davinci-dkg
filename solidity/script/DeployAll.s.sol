// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {DKGManager} from "../src/DKGManager.sol";
import {ContributionVerifier} from "../src/verifiers/ContributionVerifier.sol";
import {FinalizeVerifier} from "../src/verifiers/FinalizeVerifier.sol";
import {PartialDecryptVerifier} from "../src/verifiers/PartialDecryptVerifier.sol";
import {DecryptCombineVerifier} from "../src/verifiers/DecryptCombineVerifier.sol";
import {RevealSubmitVerifier} from "../src/verifiers/RevealSubmitVerifier.sol";
import {RevealShareVerifier} from "../src/verifiers/RevealShareVerifier.sol";

contract DeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        uint32 chainId = uint32(vm.envUint("CHAIN_ID"));
        vm.startBroadcast(deployerPrivateKey);

        ContributionVerifier contributionVerifier = new ContributionVerifier();
        console.log("ContributionVerifier deployed at:", address(contributionVerifier));

        FinalizeVerifier finalizeVerifier = new FinalizeVerifier();
        console.log("FinalizeVerifier deployed at:", address(finalizeVerifier));

        PartialDecryptVerifier partialDecryptVerifier = new PartialDecryptVerifier();
        console.log("PartialDecryptVerifier deployed at:", address(partialDecryptVerifier));

        DecryptCombineVerifier decryptCombineVerifier = new DecryptCombineVerifier();
        console.log("DecryptCombineVerifier deployed at:", address(decryptCombineVerifier));

        RevealSubmitVerifier revealSubmitVerifier = new RevealSubmitVerifier();
        console.log("RevealSubmitVerifier deployed at:", address(revealSubmitVerifier));

        RevealShareVerifier revealShareVerifier = new RevealShareVerifier();
        console.log("RevealShareVerifier deployed at:", address(revealShareVerifier));

        DKGRegistry registry = new DKGRegistry();
        console.log("DKGRegistry deployed at:", address(registry));

        DKGManager manager = new DKGManager(
            chainId,
            address(registry),
            address(contributionVerifier),
            address(partialDecryptVerifier),
            address(finalizeVerifier),
            address(decryptCombineVerifier),
            address(revealSubmitVerifier),
            address(revealShareVerifier)
        );
        console.log("DKGManager deployed at:", address(manager));

        vm.stopBroadcast();
    }
}
