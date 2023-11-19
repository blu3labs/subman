import React from "react";
import Wallet from "./wallet";
import { Subman } from "../../plugin";
import "./index.css";

function App() {
  return (
    <div className="app">
      <div className="appTitle">SubMan Plugin Demo Dapp</div>
      <Wallet />

      <div className="appCards">
        <Subman planId={1} chainId={84_531} />

        <Subman planId={1} chainId={84_531}
          backgroundColor={"red"}
          textColor={"#ffffff"}
          showPrice={true}
        />

        <Subman planId={1} chainId={534351} 
          backgroundColor={"#f0f0f0"}
          textColor={"#000000"}
        />
      </div>
    </div>
  );
}

export default App;
