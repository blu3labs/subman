package converter

import (
	"strconv"
	types "subman/common"
	contract "subman/contract"
)

// ConvertToType converts contract.SubPayment to types.
func ConvertToType(ctrSubPay contract.LibSubSubPayment) (types.SubPayment, error) {
	subPlanId, err := strconv.ParseInt(ctrSubPay.SubPlanId.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	startTime, err := strconv.ParseInt(ctrSubPay.StartTime.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	endTime, err := strconv.ParseInt(ctrSubPay.EndTime.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	duration, err := strconv.ParseInt(ctrSubPay.Duration.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	return types.SubPayment{
		SubPlanID:    subPlanId,
		Subscriber:   ctrSubPay.Subscriber.String(),
		StartTime:    startTime,
		EndTime:      endTime,
		Duration:     duration,
		PaymentToken: ctrSubPay.PaymentToken.String(),
		Price:        ctrSubPay.Price.String(),
	}, nil
}
