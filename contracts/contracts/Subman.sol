// SPDX-License-Identifier: MIT

pragma solidity 0.8.23;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {LibSub} from "./libraries/LibSub.sol";
import {SubVerifier} from "./SubVerifier.sol";

contract SubMan is SubVerifier, ReentrancyGuard {
    using EnumerableSet for EnumerableSet.UintSet;
    using SafeERC20 for IERC20;

    uint256 public subPlanCount;

    mapping(uint256 => LibSub.SubPlan) private _subPlans;
    mapping(address => uint256[]) private _subPlansByOwner;
    mapping(address => EnumerableSet.UintSet) private _subedPlans;
    mapping(address => mapping(uint256 => uint256)) private _userSubDeadline;

    event SubPaymentCanceled(LibSub.SubPayment subPayment);
    event SubPaymentProcessed(LibSub.SubPayment subPayment, uint256 newDeadline);
    event SubPlanActivated(uint256 subPlanId, LibSub.SubPlan subPlan);
    event SubPlanCreated(uint256 subPlanId, LibSub.SubPlan subPlan);
    event SubPlanDeactivated(uint256 subPlanId, LibSub.SubPlan subPlan);

    constructor(string memory _name, string memory _version) SubVerifier(_name, _version) {}

    function getSubPlan(uint256 _subPlanId) public view returns (LibSub.SubPlan memory) {
        return _subPlans[_subPlanId];
    }

    function getSubPlanIdsByOwner(address _owner) public view returns (uint256[] memory) {
        return _subPlansByOwner[_owner];
    }

    function getSubPlansByOwner(address _owner) public view returns (LibSub.SubPlan[] memory) {
        uint256[] memory _subPlanIds = getSubPlanIdsByOwner(_owner);
        LibSub.SubPlan[] memory subPlansByOwner = new LibSub.SubPlan[](_subPlanIds.length);
        for (uint256 i = 0; i < _subPlanIds.length; i++) {
            subPlansByOwner[i] = getSubPlan(_subPlanIds[i]);
        }
        return subPlansByOwner;
    }

    function getSubedPlans(address _subscriber) public view returns (uint256[] memory) {
        return _subedPlans[_subscriber].values();
    }

    function getUserSubDeadline(address _subscriber, uint256 _subPlanId) public view returns (uint256) {
        return _userSubDeadline[_subscriber][_subPlanId];
    }

    function getActiveSubPlans(address _subscriber) public view returns (uint256[] memory) {
        uint256[] memory subedPlans = getSubedPlans(_subscriber);
        uint256[] memory activeSubPlans = new uint256[](subedPlans.length);
        uint256 _activeSubPlansCount = 0;
        for (uint256 i = 0; i < subedPlans.length; i++) {
            uint256 _subPlanId = subedPlans[i];
            if (_userSubDeadline[_subscriber][_subPlanId] >= block.timestamp) {
                activeSubPlans[_activeSubPlansCount] = _subPlanId;
                _activeSubPlansCount++;
            }
        }
        uint256[] memory _activeSubscriptions = new uint256[](_activeSubPlansCount);
        for (uint256 i = 0; i < _activeSubPlansCount; i++) {
            _activeSubscriptions[i] = activeSubPlans[i];
        }
        return _activeSubscriptions;
    }

    function getActiveSubscriptions(address _subscriber) public view returns (LibSub.Subscription[] memory) {
        uint256[] memory _activeSubPlans = getActiveSubPlans(_subscriber);
        LibSub.Subscription[] memory _activeSubscriptions = new LibSub.Subscription[](_activeSubPlans.length);
        for (uint256 i = 0; i < _activeSubPlans.length; i++) {
            _activeSubscriptions[i] = LibSub.Subscription({
                subPlan: _subPlans[_activeSubPlans[i]],
                deadline: _userSubDeadline[_subscriber][_activeSubPlans[i]]
            });
        }
        return _activeSubscriptions;
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
            subPlanId: subPlanCount,
            serviceFee: _serviceFee,
            chainId: block.chainid,
            active: true
        });

        _subPlansByOwner[msg.sender].push(subPlanCount);

        emit SubPlanCreated(subPlanCount, _subPlans[subPlanCount]);
    }

    function activateSubPlan(uint256 _subPlanId) external nonReentrant {
        LibSub.SubPlan storage _subPlan = _subPlans[_subPlanId];
        require(_subPlan.owner == msg.sender, "SubMan: caller is not the owner");
        require(!_subPlan.active, "SubMan: subPlan is already active");
        
        _subPlan.active = true;

        emit SubPlanActivated(_subPlanId, _subPlan);
    }

    function deactivateSubPlan(uint256 _subPlanId) external nonReentrant {
        LibSub.SubPlan storage _subPlan = _subPlans[_subPlanId];
        require(_subPlan.owner == msg.sender, "SubMan: caller is not the owner");
        require(_subPlan.active, "SubMan: subPlan is not active");

        _subPlan.active = false;

        emit SubPlanDeactivated(_subPlanId, _subPlan);
    }

    function processPayment(
        LibSub.SubPayment memory _subPayment,
        bytes memory _signature,
        address _serviceFeeReceiver
    ) external nonReentrant {
        verify(_subPayment, _signature);
        require(_subPayment.startTime <= block.timestamp, "SubMan: startTime is in the future");
        LibSub.SubPlan memory _subPlan = _subPlans[_subPayment.subPlanId];
        uint256 newDeadline = block.timestamp + _subPlan.duration;
        require(_subPlan.active, "SubMan: subPlan is not active");
        require(_subPayment.endTime >= newDeadline, "SubMan: signature is expired");
        require(_subPlan.paymentToken == _subPayment.paymentToken, "SubMan: invalid signed data (paymentToken)");
        require(_subPlan.duration == _subPayment.duration, "SubMan: invalid signed data (duration)");
        require(_subPlan.price == _subPayment.price, "SubMan: subPlan price different");
        require(_userSubDeadline[_subPayment.subscriber][_subPayment.subPlanId] < block.timestamp, "SubMan: user already subscribed");

        _subedPlans[_subPayment.subscriber].add(_subPayment.subPlanId);
        _userSubDeadline[_subPayment.subscriber][_subPayment.subPlanId] = newDeadline;

        IERC20(_subPlan.paymentToken).safeTransferFrom(
            _subPayment.subscriber,
            _subPlan.paymentReceiver,
            _subPlan.price - _subPlan.serviceFee
        );

        IERC20(_subPlan.paymentToken).safeTransferFrom(
            _subPayment.subscriber,
            _serviceFeeReceiver,
            _subPlan.serviceFee
        );

        emit SubPaymentProcessed(_subPayment, newDeadline);
    }

    function cancel(LibSub.SubPayment memory _subPayment) external {
        require(_subPayment.subscriber == msg.sender, "SubMan: caller is not the subscriber");
        require(_subPayment.endTime > block.timestamp, "SubMan: subPayment is expired");
        
        _cancelSubPayment(_subPayment);
        
        emit SubPaymentCanceled(_subPayment);
    }
}
