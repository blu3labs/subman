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

function Sub({ item }) {
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

            <Typography className="subEndDate"
          
            >
              Your subscription will end on {item.endDate}
            </Typography>

            <Input  
              label="Extend Subscription"
              placeholder="Ex. 1"
              type="number"
              suffix="Week"
            />


          <div className="subButtons">

            <Button colorStyle="blueSecondary" >
              Extend
            </Button>
            <Button colorStyle="redSecondary" >
              Cancel
            </Button>
          </div>
          </Card>
        </Modal>
      )}
    </>
  );
}

export default Sub;
