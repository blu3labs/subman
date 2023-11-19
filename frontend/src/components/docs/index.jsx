import React from "react";
import Base from "@/assets/partner/base.jpeg";
import Certora from "@/assets/partner/certora.jpeg";
import Ens from "@/assets/partner/ens.png";
import Nouns from "@/assets/partner/nouns.png";
import Scroll from "@/assets/partner/scroll.png";
import Wc from "@/assets/partner/wc.png";
import Npm from "@/assets/npm.png";
import "./index.css";

function Docs() {
  let partners = [
    {
      link: "https://base.org/",
      logo: Base,
    },
    {
      link: "https://scroll.io/",
      logo: Scroll,
    },
    {
      link: "https://ens.domains/",
      logo: Ens,
    },
    {
      link: "https://walletconnect.com/",
      logo: Wc,
    },
    {
      link: "https://nouns.wtf/",
      logo: Nouns,
    },

    {
      link: "https://www.certora.com/",
      logo: Certora,
    },
  ];

  return (
    <div className="docsArea">
      <div className="docsAreaTop">
        <a
          href="https://www.npmjs.com/package/subman-plugin"
          target="_blank"
          className="docsBanner"
        >
          <span>Install SubMan Plugin</span>
          <img src={Npm} alt="" draggable="false" />
        </a>
        <div className="docsInfoArea">
          <div className="docsInfoAreaTitle">Installation</div>
          <div className="docsInfoBox">
            <span>npm install subman-plugin wagmi ethers</span>
          </div>
        </div>

      

        <div className="docsInfoArea">
          <div className="docsInfoAreaTitle">Usage</div>
          <div className="docsInfoBox">
            <span>{`import { Subman } from "subman-plugin"`}</span>
            <span>
              {`<Subman planId={1}  chainId={84531} showPrize={true} backgroundColor="#00a9ff" textColor="#fff" />`}
            </span>
          </div>
        </div>

        <div className="docsDesc">
           Lorem ipsum dolor sit, amet consectetur adipisicing elit. Doloribus sed, laudantium repudiandae deserunt exercitationem nesciunt possimus quidem, iusto cupiditate, obcaecati quos cumque nobis repellat consequuntur tempore recusandae aspernatur placeat dolores quasi animi porro eos consequatur natus consectetur. Voluptatem odio, repudiandae amet similique sapiente aspernatur exercitationem?
        </div>

        <div className="docsInfoArea">
          <div className="docsInfoAreaTitle">View</div>
          <div className="docsInfoBoxButton">
            <a
              href="https://www.npmjs.com/package/subman-plugin"
              target="_blank"
            >
              View on Demo App
            </a>
            <a href="https://github.com/blu3labs/subman" target="_blank">
              View on Github
            </a>
          </div>
        </div>
      </div>
      <div className="docsAreaBottom">
        <div className="docsAreaTitle">Build with</div>
        <div className="docsAreaItems">
          {partners.map((item, index) => (
            <a href={item.link} target="_blank" key={index}>
              <img src={item.logo} alt="" draggable="false" />
            </a>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Docs;
