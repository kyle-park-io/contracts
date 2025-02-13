// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "./interface/IE.sol";

contract Example is IE, ERC165 {
    bytes4 public constant _INTERFACE_ID_Example = type(IE).interfaceId;
    uint256 balance;

    // // IERC165,
    // /// @inheritdoc IERC165
    // function supportsInterface(
    //     bytes4 interfaceId
    // ) public view virtual override(ERC165) returns (bool) {
    //     return
    //         interfaceId == type(IE).interfaceId ||
    //         super.supportsInterface(interfaceId);
    // }

    function getIERC165Id() public pure returns (bytes4) {
        return type(IERC165).interfaceId;
    }

    function getERC165Id() public pure returns (bytes4) {
        return type(ERC165).interfaceId;
    }

    function getInterfaceId() public pure returns (bytes4) {
        return _INTERFACE_ID_Example;
    }

    function getBalance() public view returns (uint256) {
        return balance;
    }

    function setBalance(uint256 amount) public {
        balance += amount;
    }
}
