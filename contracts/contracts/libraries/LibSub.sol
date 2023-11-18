// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

library LibSub {
    bytes32 constant SUB_PAYMENT_TYPEHASH = keccak256(
        "SubPayment(uint256 subPlanId,address subscriber,uint256 startTime,uint256 endTime,uint256 duration,address paymentToken,uint256 price)"
    );

    struct SubPlan {
        string  title;
        string  description;
        address owner;
        address paymentReceiver;
        address paymentToken;
        uint256 price;
        uint256 duration;
        uint256 deadline;
        uint256 subPlanId;
        uint256 serviceFee;
        bool active;
    }

    struct SubPayment {
        uint256 subPlanId;
        address subscriber;
        uint256 startTime;
        uint256 endTime;
        uint256 duration;
        address paymentToken;
        uint256 price;
    }

    struct Subscription {
        SubPlan subPlan;
        uint256 deadline;
    }

    function hashSubPayment(SubPayment memory _subPayment) internal pure returns (bytes32) {
        return keccak256(abi.encode(
            SUB_PAYMENT_TYPEHASH,
            _subPayment.subPlanId,
            _subPayment.subscriber,
            _subPayment.startTime,
            _subPayment.endTime,
            _subPayment.duration,
            _subPayment.paymentToken,
            _subPayment.price
        ));
    }
}