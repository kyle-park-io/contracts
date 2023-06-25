// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Address {
    function A() public view returns (address) {
        return msg.sender;
    }
}
