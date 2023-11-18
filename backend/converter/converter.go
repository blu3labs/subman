package converter

import (
	"strconv"
	contract "subman/contract"
	types "subman/types"
)

func ConvertToType(ctrSubPay contract.LibSubSubPayment) (types.SubPayment, error) {
	subPlanId, err := strconv.ParseUint(ctrSubPay.SubPlanId.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	startTime, err := strconv.ParseUint(ctrSubPay.StartTime.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	endTime, err := strconv.ParseUint(ctrSubPay.EndTime.String(), 10, 64)
	if err != nil {
		return types.SubPayment{}, err
	}
	duration, err := strconv.ParseUint(ctrSubPay.Duration.String(), 10, 64)
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
