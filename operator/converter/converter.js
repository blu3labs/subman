"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConvertToTxArgs = exports.ConvertToSubscription = void 0;
const ethers_1 = require("ethers");
function ConvertToSubscription(data) {
    const sub = {
        SubPayment: {
            SubPlanID: data.subPayment.subPlanId,
            Subscriber: data.subPayment.subscriber,
            StartTime: data.subPayment.startTime,
            EndTime: data.subPayment.endTime,
            Duration: data.subPayment.duration,
            PaymentToken: data.subPayment.paymentToken,
            Price: data.subPayment.price,
        },
        ChainID: data.chainId,
        SubDeadline: data.subDeadline,
        PlanActive: data.planActive,
        Canceled: data.canceled,
        Signature: data.signature,
    };
    return sub;
}
exports.ConvertToSubscription = ConvertToSubscription;
function ConvertToTxArgs(sub) {
    const txArgs = {
        subPayment: {
            subPlanId: sub.SubPayment.SubPlanID,
            subscriber: sub.SubPayment.Subscriber,
            startTime: sub.SubPayment.StartTime,
            endTime: sub.SubPayment.EndTime,
            duration: sub.SubPayment.Duration,
            paymentToken: sub.SubPayment.PaymentToken,
            price: (0, ethers_1.parseUnits)(sub.SubPayment.Price, 0),
        },
        signature: sub.Signature,
    };
    return txArgs;
}
exports.ConvertToTxArgs = ConvertToTxArgs;
module.exports = {
    ConvertToSubscription,
    ConvertToTxArgs,
};
