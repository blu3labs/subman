import React, { useState } from "react";
import {
  Button,
  Modal,
  Card,
  Typography,
  Heading,
  Input,
  Textarea,
  Select,
  CurrencyToggle,
  FieldSet,
  RadioButton,
  RadioButtonGroup,
  Tooltip,
} from "@ensdomains/thorin";
import { useSigner } from "@/utils/useSigner";
import { writeContract } from "@/utils/writeContract";
import { useAccount, useNetwork } from "wagmi";
import { ethers } from "ethers";
import { currencyAddress } from "@/utils/currencyAddress";
import { toast } from "react-hot-toast";
import { submanAddress, submanAbi } from "@/contract";  
import "./index.css";

function CreateSub() {
  const [visible, setVisible] = useState(false);
  const [title, setTitle] = useState(null);
  const [description, setDescription] = useState(null);
  const [cycleDuration, setCycleDuration] = useState(null);
  const [currency, setCurrency] = useState("DAI");
  const [tokenAddress, setTokenAddress] = useState(null);
  const [tokenName, setTokenName] = useState(null);
  const [subscriptionPrice, setSubscriptionPrice] = useState(null);
  const [serviceFee, setServiceFee] = useState(null);

  const { chain } = useNetwork();

  const signer = useSigner();
  const { address } = useAccount();


  const [loading, setLoading] = useState(false);
  const handleCreate = async () => {
    if (!signer) {
      toast.error("Please connect your wallet");
      return;
    }

    setLoading(true);
    try {
      //todo other
      let currency_ = currencyAddress[chain?.id]?.filter(
        (item) => item?.symbol === currency
      )?.[0];
   
      let price_ = ethers.utils.parseUnits(
        subscriptionPrice?.toString(),
        currency_?.decimals
      );

   

      let serviceFee_ = ethers.utils.parseUnits(
        serviceFee?.toString(),
        currency_?.decimals
      );

      let context = {
        signer: signer,
        address: submanAddress[chain?.id],
        abi: submanAbi,
        method: "createSubPlan",
        message: "Successfully created subscription plan",
        args: [
          title,
          description,
          address,
          currency_?.address,
          price_,
          cycleDuration * 60,
          serviceFee_,
        ],
      };

 



      //todo modalı kapat

      let res = await writeContract(context);


    } catch (e) {
      console.log(e);
    }
    setLoading(false);
  };

  return (
    <>
      <Button colorStyle="blueSecondary" onClick={() => setVisible(true)}>
        Create a Subscription
      </Button>

      {visible && (
        <Modal open={visible} onDismiss={() => setVisible(false)}>
          <Card style={{ width: "100%" }}>
            <Heading align="center">Create a Subscription</Heading>
            <Input
              label="Title"
              placeholder="Enter a title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />

            <Textarea
              label="Description"
              placeholder="Enter a description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />

            <Input
              label="Cycle Duration"
              placeholder="Enter cycle duration"
              type="number"
              suffix="Week"
              value={cycleDuration}
              onChange={(e) => setCycleDuration(e.target.value)}
            />

            <RadioButtonGroup
              label="Currency"
              name="RadioButtonGroup"
              value={currency}
              onChange={(e) => setCurrency(e.target.value)}
            >
              <RadioButton
                label="WETH"
                name="RadioButtonGroup"
                value="WETH"
                disabled
              />
              <RadioButton label="DAI" name="RadioButtonGroup" value="DAI" />
              <RadioButton label="USDC" name="RadioButtonGroup" value="USDC" />
              <RadioButton
                label="OTHER"
                name="RadioButtonGroup"
                value="OTHER"
              />
            </RadioButtonGroup>

            {currency === "OTHER" && (
              <Input
                label="Token Address"
                placeholder="0xA0Cf…251e"
                suffix={true && "symbol"}
                value={tokenAddress}
                onChange={(e) => setTokenAddress(e.target.value)}
              />
            )}

            <Input
              label="Subscription Price"
              placeholder="Enter subscription price"
              type="number"
              suffix={currency}
              value={subscriptionPrice}
              onChange={(e) => setSubscriptionPrice(e.target.value)}
            />

            <Input
              label={
                <span title="Fee you are willing to pay for automation on each subscription. Make sure it is worth more than the execution gas fee!">
                  Service Fee *
                </span>
              }
              placeholder="Enter service fee"
              type="number"
              suffix={currency}
              value={serviceFee}
              onChange={(e) => setServiceFee(e.target.value)}
            />

            <Button
              colorStyle="blueSecondary"
              style={{ width: "100%" }}
              onClick={() => handleCreate()}
              loading={loading}
              disabled={loading}
              
            >
              Create
            </Button>
          </Card>
        </Modal>
      )}
    </>
  );
}

export default CreateSub;
