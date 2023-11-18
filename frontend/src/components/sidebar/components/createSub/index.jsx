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
} from "@ensdomains/thorin";
import "./index.css";

function CreateSub() {
  const [visible, setVisible] = useState(false);

  const [currency, setCurrency] = useState("USDT");

  return (
    <>
      <Button colorStyle="blueSecondary" onClick={() => setVisible(true)}>
        Create a Subscription
      </Button>

      {visible && (
        <Modal open={visible} onDismiss={() => setVisible(false)}>
          <Card style={{ width: "100%" }}>
            <Heading align="center">Create a Subscription</Heading>
            <Input label="Title" placeholder="Enter a title" />

            <Textarea label="Description" placeholder="Enter a description" />
            <Input
              label="Week a Price"
              placeholder="Enter week a price"
              type="number"
            />

            <RadioButtonGroup
              label="Currency"
              name="RadioButtonGroup"
              value={currency}
              onChange={(e) => setCurrency(e.target.value)}
              inline={false}
            >
              <RadioButton label="USDT" name="RadioButtonGroup" value="USDT" />
              <RadioButton label="USDC" name="RadioButtonGroup" value="USDC" />
              <RadioButton label="BUSD" name="RadioButtonGroup" value="BUSD" />
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
                suffix={true && "michel"}
              />
            )}

            <Button colorStyle="blueSecondary" style={{ width: "100%" }}>
              Create
            </Button>
          </Card>
        </Modal>
      )}
    </>
  );
}

export default CreateSub;
