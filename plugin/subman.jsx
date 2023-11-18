import React, { useState, useEffect } from "react";
import "./subman.css";

const Subman = ({
  showPrice,
  subscriptionId,
  chainId,
  backgroundColor,
  textColor,
}) => {
  const [visible, setVisible] = useState(false);

  console.log("plugin", showPrice, subscriptionId, chainId);

  showPrice = showPrice ?? false;

  let buttonDisabled = true;

  let userIsSubscribed = false;

  let buttonText = {
    true: "Subscribed",
    false: "Subscription " + (showPrice ? "| 0.1 USDT" : ""),
  };

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
        {buttonText[userIsSubscribed]}
      </button>

      {visible && (
        <div className="subman-modal">
          <div
            className="subman-modal-mask"
            onClick={() => setVisible(false)}
          ></div>
          <div className="subman-modal-content">
            <div className="subman-modal-title">
              Lorem ipsum dolor sit amet.
            </div>

            <div className="subman-modal-description">
              Lorem ipsum dolor sit amet consectetur, adipisicing elit.
              Asperiores recusandae molestias incidunt tempore? Rem, hic!
            </div>

            {userIsSubscribed && (
              <div className="subman-modal-endDate">
                Your subscription will end on 2021-08-30
              </div>
            )}

            <div className="subman-modal-week">
              <div className="subman-modal-week-title">
                {userIsSubscribed ? "Extend Subscribe" : "Subscribe"}
              </div>
              <div className="subman-modal-week-input">
                <input type="number" placeholder="Ex. 1" />
                <span>Week</span>
              </div>
            </div>

            {userIsSubscribed ? (
              <div className="subman-modal-buttons">
                <button className="subman-modal-button">Extend</button>
                <button className="subman-modal-cancel-button">Cancel</button>
              </div>
            ) : (
              <button className="subman-modal-button">
                Subscribe | 0.1 USDT
              </button>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

export default Subman;
