// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {LibSub} from "./libraries/LibSub.sol";
import {SubVerifier} from "./SubVerifier.sol";

contract SubMan is SubVerifier, ReentrancyGuard {
    using EnumerableSet for EnumerableSet.UintSet;
    using LibSub for LibSub.SubPayment;

    uint256 public subPlanCount;

    mapping(uint256 => LibSub.SubPlan) private _subPlans;
    mapping(address => uint256[]) private _subPlansByOwner;
    mapping(address => EnumerableSet.UintSet) private _subedPlans; //add getters

    event SubPlanCreated(uint256 subPlanId, LibSub.SubPlan subPlan);

    constructor(string memory _name, string memory _version) SubVerifier(_name, _version) {}

    function getSubPlan(uint256 _subPlanId) public view returns (LibSub.SubPlan memory) {
        return _subPlans[_subPlanId];
    }

    function getSubPlansByOwner(address _owner) public view returns (uint256[] memory) {
        return _subPlansByOwner[_owner];
    }

    function createSubPlan(
        string memory _title,
        string memory _description,
        address _paymentReceiver,
        address _paymentToken,
        uint256 _price,
        uint256 _duration,
        uint256 _serviceFee
    ) external nonReentrant {
        require(_duration > 0, "SubMan: duration must be greater than 0");
        require(_price > 0, "SubMan: price must be greater than 0");
        require(_serviceFee > 0, "SubMan: serviceFee must be greater than 0");
        require(_serviceFee < _price, "SubMan: serviceFee must be less than price");
        require(_paymentReceiver != address(0), "SubMan: paymentReceiver is the zero address");
        require(_paymentToken != address(0), "SubMan: paymentToken is the zero address");

        subPlanCount++;

        _subPlans[subPlanCount] = LibSub.SubPlan({
            title: _title,
            description: _description,
            owner: msg.sender,
            paymentReceiver: _paymentReceiver,
            paymentToken: _paymentToken,
            price: _price,
            duration: _duration,
            deadline: block.timestamp + _duration,
            active: true,
            serviceFee: _serviceFee
        });

        _subPlansByOwner[msg.sender].push(subPlanCount);

        emit SubPlanCreated(subPlanCount, _subPlans[subPlanCount]);
    }

    function processPayment(
        LibSub.SubPayment memory _subPayment,
        bytes memory _signature
    ) external nonReentrant {
        verify(_subPayment, _signature);
        require(_subPayment.startTime > block.timestamp, "SubMan: startTime must be greater than current time");
        LibSub.SubPlan memory _subPlan = _subPlans[_subPayment.subPlanId];
        require(_subPlan.active, "SubMan: subPlan is not active");
        require(_subPlan.deadline >= block.timestamp + _subPlan.duration, "SubMan: subPlan is expired");
        require(_subPlan.price == _subPayment.price, "SubMan: subPlan price changed");
        require(_subPayment.endTime >= block.timestamp + _subPlan.duration, "SubMan: signature is expired");
    }
}
