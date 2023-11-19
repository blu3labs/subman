import React, { useState } from "react";
import {
  Button,
  Modal,
  Card,
  Typography,
  Heading,
  Input,
} from "@ensdomains/thorin";
import { currencyAddress } from "@/utils/currencyAddress";
import "./index.css";
import { submanAbi, submanAddress } from "../../../../contract";
import { useSwitchNetwork } from "wagmi";
import { writeContract } from "../../../../utils/writeContract";
import { useSigner } from "../../../../utils/useSigner";

function OwnSub({ item }) {
  const [visible, setVisible] = useState(false);

  const { switchNetworkAsync } = useSwitchNetwork();

  const signer = useSigner();



  //todo chainId
  let currency = currencyAddress?.[534_351]?.filter(
    (item_) => item_?.address === item?.paymentToken
  )?.[0];



  let currencyDecimals = currency?.decimals;
  let currencySymbol = currency?.symbol;

  let duration = item?.duration?.toString(10) / 60;




  const [activateLoading, setActivateLoading] = useState(false);
  const handleActivate = async () => {

    setActivateLoading(true)
    try{

      let context = {
          signer: signer,
          address: submanAddress[item?.chainId],
          abi: submanAbi,
          method: "activateSubPlan",
          args: [item?.subPlanId],
          message: "Activate Subscription",
          chainId: item?.chainId,
          switchNetworkAsync: switchNetworkAsync,
      }

      await writeContract(context)

    }catch(e){
      console.log(e)
    }
    setActivateLoading(false)
  }

  const [deactivateLoading, setDeactivateLoading] = useState(false);
  const handleDeactivate = async () => {
      
      setDeactivateLoading(true)
      try{
  
        let context = {
            signer: signer,
            address: submanAddress[item?.chainId?.toString(10)],
            abi: submanAbi,
            method: "deactivateSubPlan",
            args: [item?.subPlanId],
            message: "Deactivate Subscription",
            chainId: item?.chainId?.toString(10),
            switchNetworkAsync: switchNetworkAsync,
        }
  
        await writeContract(context)
  
      }catch(e){
        console.log(e)
      }
      setDeactivateLoading(false)
    }

  return (
    <>
      <button onClick={() => setVisible(true)}>
        <span title={item?.title}>{item?.title}</span>
        <span>
          {parseFloat(
            item?.price?.toString(10) / 10 ** currencyDecimals
          ).toLocaleString("en-US")}{" "}
          {currencySymbol}
        </span>
      </button>

      {visible && (
        <Modal open={visible} onDismiss={() => setVisible(false)}>
          <Card style={{ width: "100%" }}>
            <Heading align="center">{item?.title}</Heading>

            <Typography className="subDescription">
              {item?.description}
            </Typography>

            {/* <Typography>
              150 Person has subscribed to your subscription
            </Typography> */}

            {/* <div className="subUpdatePriceWrapper">
              <Input
                label="Update Price"
                placeholder="Ex. 0.1"
                type="number"
                suffix={currencySymbol}
              />
              <Button colorStyle="blueSecondary">Update</Button>
            </div> */}

            <div className="ownSubModalInfo">
              <Typography>Price</Typography>

              <Typography>
                <b>
                  {parseFloat(
                    item?.price?.toString(10) / 10 ** currencyDecimals
                  ).toLocaleString("en-US")}{" "}
                  {currencySymbol}
                </b>
              </Typography>
            </div>

            <div className="ownSubModalInfo">
              <Typography>Cycle Duration</Typography>

              <Typography>
                <b>{duration} Weeks</b>
              </Typography>
            </div>

            <div className="ownSubModalInfo">
              <Typography>
                Plan Id
              </Typography>

              <Typography>
                <b>{item?.subPlanId?.toString(10)}</b>
              </Typography>
            </div>

            <div className="ownSubModalInfo">
              <Typography>Chain Id</Typography>

              <Typography>
                <b>{item?.chainId?.toString(10)}</b> 
              </Typography>
            </div>
                {/* todo */}

            <div className="subButtons">
              <Button colorStyle="blueSecondary" 
              loading={activateLoading}
              onClick={handleActivate}
              disabled={activateLoading || item.active}
              >
                Activate
              </Button>
              <Button colorStyle="redSecondary" 
              disabled={!item.active || deactivateLoading}
              loading={deactivateLoading}
              onClick={handleDeactivate}
              >
                Deactivate
              </Button>
            </div>
          </Card>
        </Modal>
      )}
    </>
  );
}

export default OwnSub;
