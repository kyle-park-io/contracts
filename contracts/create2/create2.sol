// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./a.sol";
import "./interface/ia.sol";
import "hardhat/console.sol";

contract Create2 {
    bytes32 a;
    uint256 b;

    constructor() {
        a = 0x2542e5e2c816ee4d0e0459559e57951cbeedd3ec5b22a8cc3fed1065fcb76ea0;
        b = 1000;
    }

    function contractBytecode() public pure returns (bytes memory code) {
        return (type(A).creationCode);
    }

    function saltResult() public view returns (bytes32 salt) {
        return keccak256(abi.encodePacked(a, b));
    }

    function Calc() public returns (address) {
        // contract
        bytes memory bytecode = type(A).creationCode;

        // constructor arguments
        bytes32 salt = keccak256(abi.encodePacked(a, b));

        // calculate contract address
        address calc_address;
        assembly {
            calc_address := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }

        console.log(calc_address);
        return calc_address;
    }
}
