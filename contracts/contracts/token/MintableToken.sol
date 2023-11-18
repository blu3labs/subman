// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

contract Token is ERC20 {
    uint8 private immutable _decimals;

    constructor(string memory name_, string memory symbol_, uint8 decimals_) ERC20(name_, symbol_) {
        _decimals = decimals_;
        _mint(msg.sender, 1e6 * (10 ** _decimals));
    }

    function decimals() view public override returns(uint8) {
        return _decimals;
    }

    function mint(address recepient_, uint256 amount_) external {
        _mint(recepient_, amount_);
    }

    function mint(uint256 amount_) external {
        _mint(msg.sender, amount_);
    }
}