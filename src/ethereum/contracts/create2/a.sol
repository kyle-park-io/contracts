// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./interface/ia.sol";

contract A is IA {
    function Q() public view returns (bool) {}

    function W() public {}
}
