// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./interface/IERC165.sol";

contract InterfaceTest is IERC165Test {
    bytes4 public constant _INTERFACE_ID_Example =
        type(IERC165Test).interfaceId;

    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual returns (bool) {
        return interfaceId == type(IERC165Test).interfaceId;
    }

    function getInterfaceId() public pure returns (bytes4) {
        return _INTERFACE_ID_Example;
    }
}
