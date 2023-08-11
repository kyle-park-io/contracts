// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Address {
    function A() public view returns (address) {
        return msg.sender;
    }

    function B(address addr) public view returns (bool) {
        if (addr == msg.sender) return true;
        else return false;
    }

    function C(
        bytes32 message,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public pure returns (address) {
        address pubAddress = ecrecover(message, v, r, s);
        return pubAddress;
    }
}
