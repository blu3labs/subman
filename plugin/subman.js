import { useState } from "react";
import "./subman.css";

const Subman = () => {
  const [showPrice, setShowPrice] = useState(false);

  return (
    <div className="subman-wrapper">
      <button onClick={() => setShowPrice(!showPrice)} className="subman-button">
        Subscription {showPrice && "0.1 USDT"}
      </button>
    </div>
  );
};

export default Subman;
