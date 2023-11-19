import React, { useState, useEffect } from "react";
import { useSigner } from "./utils/useSigner";
import "./subman.css";
import { submanAbi, submanAddress } from "./contract";
import { useAccount, useSwitchNetwork, erc20ABI } from "wagmi";
import { readContract } from "./utils/readContract";
import { currencyAddress } from "./utils/currencyAddress";
import { writeContract } from "./utils/writeContract";
import { ethers } from "ethers";
import { toast } from "react-hot-toast";
import { api } from "./utils/api";
import moment from "moment";
import axios from "axios";

const Subman = ({ showPrice, planId, chainId, backgroundColor, textColor }) => {
  const [visible, setVisible] = useState(false);

  const signer = useSigner();

  const { address } = useAccount();

  const { switchNetworkAsync } = useSwitchNetwork();

  showPrice = showPrice ?? false;

  backgroundColor = backgroundColor ?? "#00a9ff";
  textColor = textColor ?? "#fff";

  useEffect(() => {
    document.documentElement.style.setProperty(
      "--subman-button-bg",
      backgroundColor
    );
    document.documentElement.style.setProperty(
      "--subman-button-text",
      textColor
    );
  }, [backgroundColor, textColor]);

  const [week, setWeek] = useState(null);

  const [data, setData] = useState(null);
  const [amount_, setAmount_] = useState(0);
  const getData = async () => {
    try {
      let context = {
        chain: Number(chainId),
        address: submanAddress[chainId],
        abi: submanAbi,
        method: "getSubPlan",
        args: [planId],
      };

      let res = await readContract(context);

      setData(res);
    } catch (e) {
      console.log(e);
      setData(false);
    }
  };

  const [userData, setUserData] = useState(null);
  const [userIsSubscribed, setUserIsSubscribed] = useState(false);
  const getUserData = async () => {
    if (!address) {
      setUserData(false);
      setUserIsSubscribed(false);
      return;
    }
    try {
      let context = {
        chain: Number(chainId),
        address: submanAddress[chainId],
        abi: submanAbi,
        method: "getUserSubDeadline",
        args: [address, planId],
      };

      let res = await readContract(context);

      let timestamp = res?.toString(10);
      let now = Date.now() / 1000;

      if (parseFloat(timestamp) > now) {
        setUserIsSubscribed(true);
      } else {
        setUserIsSubscribed(false);
      }

      setUserData(parseFloat(res?.toString(10)));
    } catch (e) {
      console.log(e);
      setUserData(false);
      setUserIsSubscribed(false);
    }
  };

  useEffect(() => {
    getData();

    let interval = setInterval(() => {
      getData();
    }, 10_000);

    return () => clearInterval(interval);
  }, [planId, chainId]);

  useEffect(() => {
    getUserData();

    let interval = setInterval(() => {
      getUserData();
    }, 10_000);

    return () => clearInterval(interval);
  }, [address, planId, chainId]);

  let buttonDisabled = data === null || data === false;

  let currency = currencyAddress?.[Number(chainId)]?.filter(
    (item_) => item_?.address === data?.paymentToken
  )?.[0];

  let currencyDecimals = currency?.decimals;
  let currencySymbol = currency?.symbol;

  let subPrice = data?.price?.toString(10) / 10 ** currencyDecimals;

  let buttonText = {
    true: "Subscribed",
    false:
      "Subscription " +
      "| " +
      (showPrice
        ? parseFloat(subPrice)?.toLocaleString("en-US") + " " + currencySymbol
        : ""),
    undefined: "Subscribe",
  };

  const [allowance, setAllowance] = useState(null);

  const getAllowance = async () => {
    if (!address) {
      setAllowance(false);
      return;
    }

    try {
      let context = {
        chain: Number(chainId),
        address: currency?.address,
        abi: erc20ABI,
        method: "allowance",
        args: [address, submanAddress[Number(chainId)]],
      };

      let res = await readContract(context);

      setAllowance(res?.toString(10));
    } catch (e) {
      console.log(e);
      setAllowance(false);
    }
  };

  useEffect(() => {
    getAllowance();
  }, [address, chainId]);

  const approveVisible = (value) => {
    if (
      allowance == 0 ||
      allowance == "0" ||
      allowance == null ||
      allowance == undefined ||
      parseFloat(value) > parseFloat(allowance)
    ) {
      return true;
    } else {
      return false;
    }
  };

  const types = {
    SubPayment: [
      { name: "subPlanId", type: "uint256" },
      { name: "subscriber", type: "address" },
      { name: "startTime", type: "uint256" },
      { name: "endTime", type: "uint256" },
      { name: "duration", type: "uint256" },
      { name: "paymentToken", type: "address" },
      { name: "price", type: "uint256" },
    ],
  };

  const getDomain = () => {
    return {
      name: "SubMan",
      version: "1",
      chainId: chainId,
      verifyingContract: submanAddress[chainId],
    };
  };

  const getValues = () => {
    return {
      subPlanId: planId,
      subscriber: address,
      startTime: Math.floor(Date.now() / 1000),
      endTime: Math.floor(Date.now() / 1000) + week * 60,
      duration: parseFloat(data?.duration?.toString(10)),
      paymentToken: currency?.address,
      price: ethers.utils.parseUnits(subPrice?.toString(10), currencyDecimals),
    };
  };

  const [subscribeLoading, setSubscribeLoading] = useState(false);
  const [buttonName, setButtonName] = useState(null);

  useEffect(() => {
    if (
      week === null ||
      week === undefined ||
      week === "" ||
      isNaN(week) ||
      parseFloat(week) <= 0
    ) {
      setAmount_(subPrice);
    } else {
      setAmount_(subPrice * week);
    }
  }, [week, subPrice]);

  const handleSubscribe = async () => {
    if (!signer) {
      toast.error("Please connect your wallet");
      return;
    }

    if (
      week === null ||
      week === undefined ||
      week === "" ||
      isNaN(week) ||
      parseFloat(week) <= 0
    ) {
      toast.error("Please enter a valid number");
      return;
    }
    setSubscribeLoading(true);

    let amountNew = ethers.utils.parseUnits(
      amount_.toString(),
      currencyDecimals
    );

    if (approveVisible(amountNew)) {
      setButtonName(`Approving | ${amount_} ${currencySymbol}`);

      try {
        let contextApprove = {
          signer: signer,
          address: currency?.address,
          abi: erc20ABI,
          method: "approve",
          message: "Approved Successfully",
          args: [submanAddress[chainId], amountNew], // ethers.constants.MaxUint256
        };

        let res2 = await writeContract(contextApprove);

        if (res2 == "err") {
          setButtonName(null);
          setSubscribeLoading(false);
          return;
        }
      } catch (err) {
        console.log(err);
      }
    }

    setButtonName(`Subscribing ${amount_} ${currencySymbol}`);
    try {
      let domain_ = getDomain();
      let values_ = getValues();

      let signRes = await signer._signTypedData(domain_, types, values_);

      let backRes = await axios.post(api + "subscription", {
        subPayment: {
          ...values_,
          price: values_?.price?.toString(10),
        },
        chainId: parseFloat(chainId),
        subDeadline: 0,
        planActive: true,
        canceled: false,
        signature: signRes,
      });

      if (backRes?.status == 200) {
        toast.success("Subscribed Successfully");
      } else {
        toast.error("Something went wrong");
      }
    } catch (e) {
      console.log(e);
    }

    setButtonName(null);
    setSubscribeLoading(false);
  };

  return (
    <div className="subman-wrapper">
      <button
        onClick={() => setVisible(true)}
        className="subman-button"
        disabled={buttonDisabled}
        style={{
          background: backgroundColor ?? "#00a9ff;",
          color: textColor ?? "#fff",
        }}
      >
        {buttonText[buttonDisabled ? "undefined" : userIsSubscribed]}
      </button>

      {visible && (
        <div className="subman-modal">
          <div
            className="subman-modal-mask"
            onClick={() => setVisible(false)}
          ></div>
          <div className="subman-modal-content">
            <div className="subman-modal-title">{data?.title}</div>

            <div className="subman-modal-description">{data?.description}</div>

            {userIsSubscribed && (
              <div className="subman-modal-endDate">
                Your subscription will end on{" "}
                {moment(userData * 1000).format("DD MMM YYYY, hh:mm:ss A")}
              </div>
            )}

            <div className="subman-modal-week">
              <div className="subman-modal-week-title">
                {userIsSubscribed ? "Extend Subscribe" : "Subscribe"}
              </div>
              <div className="subman-modal-week-input">
                <input
                  type="number"
                  placeholder="Ex. 1"
                  value={week}
                  onChange={(e) => setWeek(e.target.value)}
                />
                <span>Week</span>
              </div>
            </div>

            {userIsSubscribed ? (
              <div className="subman-modal-buttons">
                <button
                  className="subman-modal-button"
                  disabled
                  style={{
                    cursor: "not-allowed",
                    opacity: "0.5",
                  }}
                >
                  Extend
                </button>
                <button className="subman-modal-cancel-button">Cancel</button>
              </div>
            ) : (
              <button
                className="subman-modal-button"
                onClick={() => handleSubscribe()}
                disabled={subscribeLoading}
                style={{
                  cursor: subscribeLoading ? "not-allowed" : "pointer",
                  opacity: subscribeLoading ? "0.7" : "1",
                }}
              >
                {buttonName
                  ? buttonName
                  : "Subscribe | " + amount_ + " " + currencySymbol}
              </button>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

export default Subman;
