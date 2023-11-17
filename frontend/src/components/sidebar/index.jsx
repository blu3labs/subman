import React from "react";
import "./index.css";

function Sidebar() {
  let subs = [
    {
      title: "Sub asjkdhs asdjkhska askldjhklasd1",
      price: "0.1",
      symbol: "ETH",
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
          <div className="sidebarHeader">
            <span>Subman</span>
          </div>
          <div className="sidebarBody">
            {true && (
              <div className="sidebarSubsWrapper">
                <div className="sidebarSubsTitle">Subscribed</div>
                <div className="sidebarSubsItems">
                  {subs.map((item, index) => (
                    <button key={index}>
                      <span title={item.title}>{item.title}</span>
                      <span>
                        {item.price} {item.symbol}
                      </span>
                    </button>
                  ))}
                </div>
              </div>
            )}

            {true && (
              <div className="sidebarSubsWrapper">
                <div className="sidebarSubsTitle">Subscribed You</div>
                <div className="sidebarSubsItems">
                  {subs.map((item, index) => (
                    <button key={index}>
                      <span title={item.title}>{item.title}</span>
                      <span>
                        {item.price} {item.symbol}
                      </span>
                    </button>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
        <div className="sidebarBottom">wallet connect</div>
      </div>
    </div>
  );
}

export default Sidebar;
