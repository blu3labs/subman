import { parseUnits } from "ethers";

type SubPayment = {
    SubPlanID: number;
    Subscriber: string;
    StartTime: number;
    EndTime: number;
    Duration: number;
    PaymentToken: string;
    Price: string;
};

export type Subscription = {
    SubPayment: SubPayment;
    ChainID: number;
    SubDeadline: number;
    PlanActive: boolean;
    Canceled: boolean;
    Signature: string;
};

export function ConvertToSubscription(data: any) {
  const sub: Subscription = {
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

export function ConvertToTxArgs(sub: Subscription) {
  const txArgs = {
    subPayment: {
      subPlanId: sub.SubPayment.SubPlanID,
      subscriber: sub.SubPayment.Subscriber,
      startTime: sub.SubPayment.StartTime,
      endTime: sub.SubPayment.EndTime,
      duration: sub.SubPayment.Duration,
      paymentToken: sub.SubPayment.PaymentToken,
      price: parseUnits(sub.SubPayment.Price, 0),
    },
    signature: sub.Signature,
  };
  return txArgs;
}

module.exports = {
  ConvertToSubscription,
  ConvertToTxArgs,
};
