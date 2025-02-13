// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Output {
    function A() public view returns (address) {
        return msg.sender;
    }

    function B() public view returns (address, address) {
        return (msg.sender, msg.sender);
    }

    function C() public view returns (address[] memory) {
        address[] memory result = new address[](2);
        return result;
    }
}
