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
        return _canceledSubPayments[_subPayment.hashSubPayment()];
    }

    function _cancelSubPayment(LibSub.SubPayment memory _subPayment) internal {
        require(!_canceledSubPayments[_subPayment.hashSubPayment()], "SubVerifier: subPayment is already canceled");
        _canceledSubPayments[_subPayment.hashSubPayment()] = true;
    }

    function verify(LibSub.SubPayment memory _subPayment, bytes memory _signature) internal view {
        require(!isCanceled(_subPayment), "SubVerifier: subPayment is canceled");
        bytes32 _hash = _hashTypedDataV4(_subPayment.hashSubPayment());
        _isValidSigner(_subPayment.subscriber, _hash, _signature);
    }

    function _isValidSigner(address _signer, bytes32 _hash, bytes memory _signature) internal view {
        if (_isContract(_signer)) {
            require(
                IERC1271(_signer).isValidSignature(
                    _hashTypedDataV4(_hash),
                    _signature
                ) == MAGICVALUE,
                "SubVerifier: ERC1271 ticket signature verification error"
            );
        } else {
            if (_hashTypedDataV4(_hash).recover(_signature) != _signer) {
                revert("SubVerifier: ECDSA ticket signature verification error");
            } else {
                require(_signer != address(0), "SubVerifier: Invalid owner");
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