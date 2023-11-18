require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    compilers: [
      {
        version: "0.8.23",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  networks: {
    baseGoerli: {
      url: "https://base-goerli.public.blastapi.io",
      accounts: [
        process.env.PV_KEY
      ],
    },
    scrollSepolia: {
      url: "https://rpc.ankr.com/scroll_sepolia_testnet",
      accounts: [
        process.env.PV_KEY
      ],
    },
  },
  etherscan: {
    apiKey: {
     "base-goerli": "PLACEHOLDER_STRING",
     "scroll-sepolia": process.env.SCROLL_API_KEY
    },
    customChains: [
      {
        network: "base-goerli",
        chainId: 84531,
        urls: {
         apiURL: "https://api-goerli.basescan.org/api",
         browserURL: "https://goerli.basescan.org"
        }
      },
      {
        network: "scroll-sepolia",
        chainId: 534351,
        urls: {
         apiURL: "https://api-sepolia.scrollscan.com/api",
         browserURL: "https://sepolia.scrollscan.com/"
        }
      }
    ]
  },
};
