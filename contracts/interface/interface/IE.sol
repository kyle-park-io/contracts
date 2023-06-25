// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/interfaces/IERC165.sol";

// interface IE is IERC165 {
interface IE {
    function getBalance() external view returns (uint256);

    function setBalance(uint256 amount) external;
}
