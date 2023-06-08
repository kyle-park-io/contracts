// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract A {
    uint256 balance;
    address other;

    constructor(address _other) {
        other = _other;
    }

    function getBalance() public view returns (uint256) {
        return balance;
    }

    function setBalance(uint256 amount) public {
        balance += amount;
    }

    function getOtherBalance() public view returns (uint256) {
        (bool success, bytes memory data) = other.staticcall(
            abi.encodeWithSignature("getBalance()")
        );
        require(success, "Failed to get other balance");
        uint256 otherBalance = abi.decode(data, (uint256));

        return otherBalance;
    }

    function setOtherBalance(uint256 amount) public {
        (bool success, ) = other.call(
            abi.encodeWithSignature("setBalance(uint256)", amount)
        );
        require(success, "Failed to set other balance");
    }

    function setBalanceDelegate(uint256 amount) public {
        (bool success, ) = other.delegatecall(
            abi.encodeWithSignature("setBalance(uint256)", amount)
        );
        require(success, "Failed to set balance by delegate call");
    }
}
