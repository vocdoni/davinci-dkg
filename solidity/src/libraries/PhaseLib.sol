// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "./DKGTypes.sol";

library PhaseLib {
    function inCommitteeSelection(DKGTypes.EpochPhase status, uint64 committeeSelectionDeadlineBlock) internal view returns (bool) {
        return status == DKGTypes.EpochPhase.CommitteeSelection && block.number <= committeeSelectionDeadlineBlock;
    }

    function inKeyAssembly(DKGTypes.EpochPhase status, uint64 keyAssemblyDeadlineBlock) internal view returns (bool) {
        return status == DKGTypes.EpochPhase.KeyAssembly && block.number <= keyAssemblyDeadlineBlock;
    }
}
