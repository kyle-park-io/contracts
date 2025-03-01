// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Proxy {
  address public logic; // ✅ Address of the logic (implementation) contract

  constructor(address _logic) {
    logic = _logic;
  }

  // ✅ This function allows the contract to receive Ether when sent directly without data
  receive() external payable {}

  fallback() external payable {
    address _logic = logic;
    require(_logic != address(0), 'Logic contract not set');

    assembly {
      let ptr := mload(0x40) // Allocate memory space
      calldatacopy(ptr, 0, calldatasize()) // Copy input data

      let result := delegatecall(gas(), _logic, ptr, calldatasize(), 0, 0) // Execute delegatecall
      let size := returndatasize()

      returndatacopy(ptr, 0, size) // Copy return data

      switch result
      case 0 {
        revert(ptr, size)
      } // Revert on failure
      default {
        return(ptr, size)
      } // Return on success
    }
  }
}

contract Logic {
  uint256 public value; // ✅ Stored in the storage of the proxy contract

  function setValue(uint256 _value) public {
    value = _value; // ✅ Stored in the proxy's storage through delegatecall
  }

  function getValue() public view returns (uint256) {
    return value; // ✅ Returns the value stored in the proxy's storage through delegatecall
  }
}
