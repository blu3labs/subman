import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.jsx";
import "./index.css";
import "./reset.css";

import { ThemeProvider } from "styled-components";
import { ThorinGlobalStyles, lightTheme } from "@ensdomains/thorin";

import { createWeb3Modal, defaultWagmiConfig } from "@web3modal/wagmi/react";
import { WagmiConfig } from "wagmi";
import {  baseGoerli, scrollSepolia } from "wagmi/chains";


const projectId = "cfe6a84206725199dbc2e6b7b8ffbccc"; // blu3 - hackathon

const metadata = {
  name: "Web3Modal",
  description: "Web3Modal Example",
  url: "https://web3modal.com",
  icons: ["https://avatars.githubusercontent.com/u/37784886"],
};

const chains = [ baseGoerli,scrollSepolia];
const wagmiConfig = defaultWagmiConfig({ chains, projectId, metadata });

createWeb3Modal({ wagmiConfig, projectId, chains, themeMode: "light" });

ReactDOM.createRoot(document.getElementById("root")).render(
  <ThemeProvider theme={lightTheme}>
    <ThorinGlobalStyles />
    <WagmiConfig config={wagmiConfig}>
      <App />
    </WagmiConfig>
  </ThemeProvider>
);
