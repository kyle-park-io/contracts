// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Uncheck {
    function add(uint256 a, uint256 b) public pure returns (uint256) {
        uint256 result = a + b;
        return result;
    }

    function tryAdd(uint256 a, uint256 b) public pure returns (bool, uint256) {
        unchecked {
            uint256 c = a + b;
            if (c < a) return (false, 0);
            return (true, c);
        }
    }
}
