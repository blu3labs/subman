import React from "react";
import { useAccount, useNetwork } from "wagmi";
import { useWeb3Modal } from "@web3modal/wagmi/react";
import "./index.css";

function Wallet() {
  let { address, isConnected } = useAccount();
  const { chain } = useNetwork();
  const { open } = useWeb3Modal();

  let wrongNetwork = chain?.unsupported;

  return (
    <div className="appWallets">
      {wrongNetwork ? (
        <button onClick={() => open({ view: "Networks" })}>
          Wrong Network
        </button>
      ) : isConnected ? (
        <button onClick={() => open({ view: "Networks" })}>
          {chain?.name}
        </button>
      ) : null}
      {isConnected ? (
        <button onClick={() => open({ view: "Account" })}>
          {address?.slice(0, 6) + "..." + address?.slice(-4)}
        </button>
      ) : (
        <button onClick={() => open()}>Connect Wallet</button>
      )}
    </div>
  );
}

export default Wallet;
