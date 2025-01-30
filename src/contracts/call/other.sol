// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Other {
    uint256 balance;
    uint256 balance_test;
    uint256 balance_test2;

    function getBalance() public view returns (uint256) {
        return balance;
    }

    function setBalance(uint256 amount) public {
        balance += amount;
    }

    function setBalance2(uint256 amount) public returns (uint256) {
        balance += amount;
        return balance;
    }

    function noVar(uint256 amount) public returns (uint256) {
        balance_test += amount;
        return balance_test;
    }

    function noVar2(uint256 amount) public returns (uint256) {
        balance_test2 += amount;
        return balance_test2;
    }
}
