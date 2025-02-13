// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Proxy {
    address implement;

    function upgradeLogic(address new_implement) public {
        implement = new_implement;
    }

    function invokeContract(uint256 amount) public {
        (bool success, ) = implement.delegatecall(
            abi.encodeWithSignature("function(uint256)", amount)
        );
        require(success, "Failed to set balance by delegate call");
    }
}
