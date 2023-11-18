import React from "react";
import { Heading } from "@ensdomains/thorin";
import Wallet from "./components/wallet";
import "./index.css";
import Sub from "./components/sub";
import OwnSub from "./components/ownSub";

function Sidebar() {
  let subs = [
    {
      title: "Sub asjkdhs asdjkhska askldjhklasd1",
      price: "0.1",
      symbol: "ETH",
      description: "lorem ajksdhaskl askldjklsad kjlasjdhkldsa",
      endDate: "12/12/2021",
    },
    {
      title: "Sub 2",
      price: "0.2",
      symbol: "ETH",
    },
    {
      title: "Sub 3",
      price: "0.3",
      symbol: "ETH",
    },
    {
      title: "Sub 1",
      price: "0.1",
      symbol: "ETH",
    },
    {
      title: "Sub 2",
      price: "0.2",
      symbol: "ETH",
    },
  ];

  return (
    <div className="sidebarWrapper">
      <div className="sidebar">
        <div className="sidebarTop">
       
          <Heading>Subman</Heading>

          <div className="sidebarBody">
            {true && (
              <div className="sidebarSubsWrapper">
                <div className="sidebarSubsTitle">Subscribed</div>
                <div className="sidebarSubsItems">
                  {subs.map((item, index) => (
                    <Sub key={index} item={item} />
                  ))}
                </div>
              </div>
            )}

            {true && (
              <div className="sidebarSubsWrapper">
                <div className="sidebarSubsTitle">Subscribed You</div>
                <div className="sidebarSubsItems">
                  {subs.map((item, index) => (
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
