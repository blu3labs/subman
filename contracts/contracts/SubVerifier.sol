// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {IERC1271} from "@openzeppelin/contracts/interfaces/IERC1271.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {LibSub} from "./libraries/LibSub.sol";

abstract contract SubVerifier is EIP712 {
    using ECDSA for bytes32;
    using LibSub for LibSub.SubPayment;

    bytes4 constant internal MAGICVALUE = 0x1626ba7e;

    mapping(bytes32 => bool) private _canceledSubPayments;

    constructor(string memory _name, string memory _version) EIP712(_name, _version) {}

    function isCanceled(LibSub.SubPayment memory _subPayment) public view returns (bool) {
        return _canceledSubPayments[_subPayment.hash()];
    }

    function _cancelSubPayment(LibSub.SubPayment memory _subPayment) internal {
        require(!_canceledSubPayments[_subPayment.hash()], "SubVerifier: subPayment is already canceled");
        _canceledSubPayments[_subPayment.hash()] = true;
    }

    function verify(LibSub.SubPayment memory _subPayment, bytes memory _signature) internal view {
        require(!isCanceled(_subPayment), "SubVerifier: subPayment is canceled");
        bytes32 hash = _subPayment.hash();
        address signer = _subPayment.subscriber;
        if (_isContract(signer)) {
            require(
                IERC1271(signer).isValidSignature(
                    _hashTypedDataV4(hash),
                    _signature
                ) == MAGICVALUE,
                "SubVerifier: ERC1271 ticket signature verification error"
            );
        } else {
            if (_hashTypedDataV4(hash).recover(_signature) != signer) {
                revert("SubVerifier: ECDSA ticket signature verification error");
            } else {
                require(signer != address(0), "SubVerifier: Invalid owner");
            }
        }
    }

    function _isContract(address account) internal view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(account)
        }
        return size > 0;
    }
}