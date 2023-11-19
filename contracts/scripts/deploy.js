const hre = require("hardhat");

async function main() {
  const SubMan = await hre.ethers.getContractFactory("SubMan");
  const subMan = await SubMan.deploy("SubMan", "1");
  await subMan.waitForDeployment();
  const address = await subMan.getAddress();
  console.log("SubMan deployed to:", address);

  await hre.run("verify:verify", {
    address: address,
    constructorArguments: ["SubMan", "1"],
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
