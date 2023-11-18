const hre = require("hardhat");
const { parseUnits } = require("ethers");

async function main() {
  const abi = [
    "function mint(address to, uint256 amount) public returns (bool)",
  ];

  // const USDC = await hre.ethers.getContractAt(abi, "0x353bE137dD7cc23396E3D4Ef7BffF930D1E3e2af");
  // const mint1 = await USDC.mint("0x1FAf588150749cAa9E87e1Cd45d29d606c1B8596", parseUnits("1000000", 6));
  // mint1.wait();
  // console.log("minted 1000000 USDC to 0x1FAf588150749cAa9E87e1Cd45d29d606c1B8596");
  const DAI = await hre.ethers.getContractAt(abi, "0xDa4d5a48efABC18c8584E11D545F80D4f1F4478C");
  const mint2 = await DAI.mint("0x1FAf588150749cAa9E87e1Cd45d29d606c1B8596", parseUnits("1000000", 18));
  mint2.wait();
  console.log("minted 1000000 DAI to 0x1FAf588150749cAa9E87e1Cd45d29d606c1B8596");

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
