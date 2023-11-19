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
import { useSwitchNetwork } from "wagmi";
import { useSigner } from "@/utils/useSigner";
import moment from "moment";

import "./index.css";

function Sub({ item, deadline }) {
  const [visible, setVisible] = useState(false);

  const { switchNetworkAsync } = useSwitchNetwork();
  const signer = useSigner();

  //todo chainId
  let currency = currencyAddress?.[item?.chainId?.toString(10)]?.filter(
    (item_) => item_?.address === item?.paymentToken
  )?.[0];

  let currencyDecimals = currency?.decimals;
  let currencySymbol = currency?.symbol;

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

            <Typography className="subEndDate">
              Your subscription will end on {" "}
             
              {
                moment(
                  deadline?.toString(10) * 1000
                ).format(
                  "DD MMM YYYY, hh:mm:ss A"
                )
              }
            </Typography>

            <Input
              label="Extend Subscription"
              placeholder="Ex. 1"
              type="number"
              suffix="Week"
              disabled
            />

            <div className="subButtons">
              <Button colorStyle="blueSecondary" disabled>
                Extend
              </Button>
              <Button colorStyle="redSecondary">Cancel</Button>
            </div>
          </Card>
        </Modal>
      )}
    </>
  );
}

export default Sub;
