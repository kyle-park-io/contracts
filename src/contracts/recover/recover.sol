// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Recover {
    address test;
    event Addr(address addr);

    function recoverAddress(
        bytes32 message,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public returns (address) {
        address addr = ecrecover(message, v, r, s);
        test = addr;
        emit Addr(addr);
        return addr;
    }

    function getAddr() public view returns (address) {
        return test;
    }

    function check() public view returns (bool) {
        if (msg.sender == test) return true;
        else return false;
    }
}
