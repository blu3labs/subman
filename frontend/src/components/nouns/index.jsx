import React from "react";
import Logo1 from "@/assets/nouns/4.png";
import Logo2 from "@/assets/nouns/2.png";
import Logo3 from "@/assets/nouns/3.png";
import Logo from "@/assets/logo.png";
import "./index.css";

function Nouns() {
  let items = [Logo1, Logo2, Logo3, Logo1, Logo2, Logo3];

  return <div className="nouns">
    {items.map((item) => (
   
        <img src={item} alt="" draggable="false"/>
  
    ))}

    <img src={Logo} alt="" draggable="false" className="nounsLogo"/>
  </div>;
}

export default Nouns;
