// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract A {
    uint256 balance;
    uint256 balance2;
    uint256 balance_test;
    uint256 balance_test2;
    address other;

    event Balance(uint256 balance);

    constructor(address _other) {
        other = _other;
    }

    function getBalance() public view returns (uint256) {
        return balance;
    }

    function getBalances()
        public
        view
        returns (uint256, uint256, uint256, uint256)
    {
        return (balance, balance2, balance_test, balance_test2);
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

    function setOtherBalance2(uint256 amount) public {
        (bool success, bytes memory result) = other.call(
            abi.encodeWithSignature("setBalance2(uint256)", amount)
        );
        require(success, "Failed to set other balance");

        bytes32 test;
        assembly {
            test := mload(add(result, 32))
        }
        emit Balance(uint256(test));
    }

    function setBalanceDelegate(uint256 amount) public {
        (bool success, ) = other.delegatecall(
            abi.encodeWithSignature("setBalance(uint256)", amount)
        );
        require(success, "Failed to set balance by delegate call");
    }

    function setBalance2Delegate(uint256 amount) public {
        (bool success, ) = other.delegatecall(
            abi.encodeWithSignature("setBalance(uint256)", amount)
        );
        require(success, "Failed to set balance by delegate call");
    }

    function setBalanceDelegateNoVar(uint256 amount) public {
        (bool success, ) = other.delegatecall(
            abi.encodeWithSignature("noVar(uint256)", amount)
        );
        require(success, "Failed to set balance by delegate call");
    }

    function setBalanceDelegateNoVar2(uint256 amount) public {
        (bool success, ) = other.delegatecall(
            abi.encodeWithSignature("noVar2(uint256)", amount)
        );
        require(success, "Failed to set balance by delegate call");
    }
}
