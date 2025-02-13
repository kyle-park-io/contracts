// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";

contract IC {
    using ERC165Checker for address;

    function isERC165(address contractAddress) external view returns (bool) {
        return contractAddress.supportsERC165();
    }
}
