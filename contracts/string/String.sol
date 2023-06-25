// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract String {
    string save;

    event Test(string a, string b, uint256 c);

    function A(string calldata req) public pure returns (string calldata) {
        return req;
    }

    function A1(
        string calldata req,
        string calldata req2
    ) public pure returns (string calldata) {
        return req;
    }

    function A2(address req, address req2) public pure returns (address) {
        return req;
    }

    function B(string memory req) public pure returns (string memory) {
        return req;
    }

    function C(string calldata req) public {
        emit Test("hi", "kyle", 100);
        save = req;
    }

    function D(string calldata req, string calldata req2) public {
        save = req;
    }

    function E(string calldata req, address req2) public {
        save = req;
    }

    function F(
        string calldata req,
        string calldata req2,
        string calldata req3
    ) public {
        save = req;
    }
}
