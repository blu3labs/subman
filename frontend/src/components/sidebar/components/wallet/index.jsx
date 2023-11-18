import React from "react";
import { useAccount, useEnsAvatar, useEnsName, useNetwork } from "wagmi";
import { useWeb3Modal } from "@web3modal/wagmi/react";
import { chainLogo } from "@/utils/chainLogo";
import { Button, LockSVG, Profile, Modal } from "@ensdomains/thorin";
import blockies from "ethereum-blockies";
import "./index.css";
import CreateSub from "../createSub";

function Wallet() {
  let { address, isConnected } = useAccount();
  const { chain } = useNetwork();
  const { open } = useWeb3Modal();

  // address = "0xb8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5"

  const ensName = useEnsName({ address, chainId: 1 });
  const ensAvatar = useEnsAvatar({ name: ensName?.data, chainId: 1 });

  let wrongNetwork = chain?.unsupported;

  let isEnsName =
    ensName?.data && !ensName.isError && !ensName.isLoading ? true : false;

  const addressOrName = {
    true: ensName?.data,
    false: address,
  };

  let userIcon = blockies.create({
    seed: address,
    color: "#dfe",
    bgcolor: "#aaafff",
    size: 15,
    scale: 3,
    spotcolor: "#00000050",
  });

  return (
    <div className="walletWrapper">
      {isConnected && !wrongNetwork && (
        <CreateSub />
      )}
      {wrongNetwork ? (
        <Button
          colorStyle="redSecondary"
          onClick={() => open({ view: "Networks" })}
        >
          Wrong Network
        </Button>
      ) : isConnected ? (
        <Button
          onClick={() => open({ view: "Networks" })}
          colorStyle="greenSecondary"
          prefix={
            chainLogo[chain?.id] && (
              <img
                src={chainLogo[chain?.id]}
                alt="chain"
                className="networkIcon"
              />
            )
          }
        >
          {chain?.name}
        </Button>
      ) : null}
      {isConnected ? (
        <Button
          colorStyle="blueSecondary"
          onClick={() => open({ view: "Account" })}
          className="profileBtn"
        >
          <Profile
            address={addressOrName[isEnsName]}
            avatar={ensAvatar?.data || userIcon?.toDataURL()}
          />
        </Button>
      ) : (
        <Button
          prefix={<LockSVG />}
          colorStyle="blueSecondary"
          onClick={() => open()}
        >
          Connect Wallet
        </Button>
      )}
    </div>
  );
}

export default Wallet;
