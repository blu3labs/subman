const hre = require("hardhat");
const { parseUnits } = require("ethers");

async function main() {
  const abi = [
    "function mint(address to, uint256 amount) public returns (bool)",
  ];

  const USDC = await hre.ethers.getContractAt(abi, "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d");
  const mint1 = await USDC.mint("0x1FAf588150749cAa9E87e1Cd45d29d606c1B8596", parseUnits("1000000", 6));
  mint1.wait();

  const DAI = await hre.ethers.getContractAt(abi, "0x353bE137dD7cc23396E3D4Ef7BffF930D1E3e2af");
  const mint2 = await DAI.mint("0x1FAf588150749cAa9E87e1Cd45d29d606c1B8596", parseUnits("1000000", 18));
  mint2.wait();

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
