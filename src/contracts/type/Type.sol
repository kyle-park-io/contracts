// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Type {
    struct a {
        address name;
        string nick;
    }

    function A(int[] calldata arr) public {}

    function B(a calldata arr) public {}

    function C(a[] calldata arr) public {}

    function D(a[] calldata arr, a[] calldata arr2) public {}

    function E(a calldata arr, a calldata arr2) public {}
}
