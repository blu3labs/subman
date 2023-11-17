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
    bscTestnet: {
      url: "https://bsc-testnet.public.blastapi.io",
      accounts: [
        process.env.PV_KEY
      ],
    },
  },
};
