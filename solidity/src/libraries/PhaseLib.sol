// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {DKGTypes} from "./DKGTypes.sol";

library PhaseLib {
    function inRegistration(DKGTypes.RoundStatus status, uint64 registrationDeadlineBlock) internal view returns (bool) {
        return status == DKGTypes.RoundStatus.Registration && block.number <= registrationDeadlineBlock;
    }

    function inContribution(DKGTypes.RoundStatus status, uint64 contributionDeadlineBlock) internal view returns (bool) {
        return status == DKGTypes.RoundStatus.Contribution && block.number <= contributionDeadlineBlock;
    }
}
