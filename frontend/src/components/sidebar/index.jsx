import React, { useEffect, useState } from "react";
import { Heading } from "@ensdomains/thorin";
import Wallet from "./components/wallet";
import "./index.css";
import Sub from "./components/sub";
import OwnSub from "./components/ownSub";
import { getAddress } from "../../utils/getAddress";
import { submanAbi, submanAddress } from "../../contract";
import { readContract } from "../../utils/readContract";

function Sidebar() {
  const address = getAddress();

  const [subPlansByOwner, setSubPlansByOwner] = useState(null);

  const getSubPlansByOwner = async () => {
    if (address === "0x0000000000000000000000000000000000000000") {
      setSubPlansByOwner(false);
      return;
    }
    try {
      let context = {
        chain: 84531,
        address: submanAddress[84531],
        abi: submanAbi,
        method: "getSubPlansByOwner",
        args: [address],
      };

      let context1 = {
        chain: 534_351,
        address: submanAddress[534_351],
        abi: submanAbi,
        method: "getSubPlansByOwner",
        args: [address],
      };

      let promiseArr = [readContract(context), readContract(context1)];

      let res = await Promise.all(promiseArr);

      let baseArr = res[0];
      let scrollArr = res[1];

      let allArr = baseArr.concat(scrollArr);

      setSubPlansByOwner(allArr);
    } catch (e) {
      console.log(e);
      setSubPlansByOwner(false);
    }
  };


  const [activeSubs, setActiveSubs] = useState(null);

  const getActiveSubs = async () => {
    if (address === "0x0000000000000000000000000000000000000000") {
      setActiveSubs(false);
      return;
    }

    try {
      let context = {
        chain: 84531,
        address: submanAddress[84531],
        abi: submanAbi,
        method: "getActiveSubscriptions",
        args: [address],
      };

      let context1 = {
        chain: 534_351,
        address: submanAddress[534_351],
        abi: submanAbi,
        method: "getActiveSubscriptions",
        args: [address],
      };

      let promiseArr = [readContract(context), readContract(context1)];

 
      let res = await Promise.all(promiseArr);

      let baseArr = res[0];
      let scrollArr = res[1];

      let allArr = baseArr.concat(scrollArr);

      setActiveSubs(allArr);


  
      
    } catch (e) {
      console.log(e);
      setActiveSubs(false);
    }
  };

  useEffect(() => {
    getSubPlansByOwner();
    getActiveSubs();

    let interval = setInterval(() => {
      getSubPlansByOwner();
      getActiveSubs();
    }, 10_000);

    return () => {
      clearInterval(interval);
    };
  }, [address]);

  return (
    <div className="sidebarWrapper">
      <div className="sidebar">
        <div className="sidebarTop">
          <Heading>Subman</Heading>

          <div className="sidebarBody">
            {activeSubs && activeSubs?.length > 0 && (
              <div className="sidebarSubsWrapper">
                <div className="sidebarSubsTitle">Subscribed</div>
                <div className="sidebarSubsItems">
                  {activeSubs?.map((item, index) => (
                    <Sub key={index} item={item?.subPlan} deadline={item?.deadline} />
                  ))}
                </div>
              </div>
            )}

            {subPlansByOwner && subPlansByOwner?.length > 0 && (
              <div className="sidebarSubsWrapper">
                <div className="sidebarSubsTitle">Your Subscribed Plans</div>
                <div className="sidebarSubsItems">
                  {subPlansByOwner?.map((item, index) => (
                    <OwnSub key={index} item={item} />
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
        <Wallet />
      </div>
    </div>
  );
}

export default Sidebar;
