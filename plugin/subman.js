import { useState } from "react";

const Subman = () => {
  const [showPrice, setShowPrice] = useState(false);

  return (
    <>
      <button onClick={() => setShowPrice(!showPrice)}>
        Subscription {showPrice && "0.1 USDT"}
      </button>
    </>
  );
};

export default Subman;
