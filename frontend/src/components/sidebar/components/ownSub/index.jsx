import React, { useState } from "react";
import {
  Button,
  Modal,
  Card,
  Typography,
  Heading,
  Input,
} from "@ensdomains/thorin";
import "./index.css";

function OwnSub({ item }) {
  const [visible, setVisible] = useState(false);

  return (
    <>
      <button onClick={() => setVisible(true)}>
        <span title={item.title}>{item.title}</span>
        <span>
          {item.price} {item.symbol}
        </span>
      </button>

      {visible && (
        <Modal open={visible} onDismiss={() => setVisible(false)}>
          <Card style={{ width: "100%" }}>
            <Heading align="center">{item.title}</Heading>

            <Typography className="subDescription">
              {item.description}
            </Typography>

            <Typography>
              150 Person has subscribed to your subscription
            </Typography>

            <Input
              label="Update Price"
              placeholder="Ex. 0.1"
              type="number"
              suffix="USDT"
            />
            <Button colorStyle="blueSecondary">Update Price</Button>

            <div className="subButtons">
              <Button colorStyle="blueSecondary">Activate</Button>
              <Button colorStyle="redSecondary">Deactivate</Button>
            </div>
          </Card>
        </Modal>
      )}
    </>
  );
}

export default OwnSub;
