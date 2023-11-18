/**
 * #Subman Specs for Certora Prover
*/

// The methods block below gives various declarations regarding solidity methods.
methods
{
    // envfree and view functions that we will use in the rules.
    function subPlanCount() external returns (uint256) envfree;
    function getSubPlan(uint) external returns (LibSub.SubPlan memory) envfree;
    function getSubPlansByOwner(address) external returns (LibSub.SubPlan[] memory) envfree;
    function getSubedPlans(address) external returns (uint256[] memory) envfree;
    function getUserSubDeadline(address,uint256) external returns (uint256) envfree;
    function getActiveSubPlans(address) external returns (uint256[] memory);
    function getActiveSubscriptions(address) external returns (LibSub.Subscription[] memory);
}
// for the check first test

// subPlanCount should increase with created new subPlan.
rule checkCreateSubPlan(
        string _title,
        string _description,
        address _paymentReceiver,
        address _paymentToken,
        uint256 _price,
        uint256 _duration,
        uint256 _serviceFee
    ) {
        env e;

        // subPlanCount before createSubPlan call.
        mathint subPlanCountBefore = subPlanCount();

        // Subscription duration must be greater than zero.
        require _duration > 0;

        // Subscription price must be greater than zero.
        require _price > 0;

        // Service fee must be greater than zero.
        require _serviceFee > 0;

        // And also _serviceFee should be less than _price
        require _serviceFee < _price;

        // Payment receiver and payment token address cannot be address(0)
        require _paymentReceiver != 0;
        require _paymentToken != 0;

        // msg.value will be used for the service fee.
        assert e.msg.value > 0;

        // call createSubPlan for the create a subscription plan.
        createSubPlan(e,
         _title,
         _description,
         _paymentReceiver,
         _paymentToken,
         _price,
         _duration,
         _serviceFee);

        // subPlanCount after createSubPlan call.
        mathint subPlanCountAfter = subPlanCount();

        // subPlanCount must increase after call createSubPlan successfully as expected.
        assert subPlanCountAfter > subPlanCountBefore, "subPlanCount must increase after call createSubPlan successfully";

}