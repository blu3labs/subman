const hre = require("hardhat");

async function main() {
  const Token = await hre.ethers.getContractFactory("Token");

  const USDC = await Token.deploy("USDC", "USDC", 6);
  await USDC.waitForDeployment();
  const USDCAddress = await USDC.getAddress();
  console.log("USDC deployed to:", USDCAddress);

  const DAI = await Token.deploy("DAI", "DAI", 18);
  await DAI.waitForDeployment();
  const DAIAddress = await DAI.getAddress();
  console.log("DAI deployed to:", DAIAddress);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
