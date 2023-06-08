// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Other {
    uint256 balance;

    function getBalance() public view returns (uint256) {
        return balance;
    }

    function setBalance(uint256 amount) public {
        balance += amount;
    }
}
