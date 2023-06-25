import fs from "fs";
import path from "path";
import { ethers } from "hardhat";

async function main() {
  const Ic = await ethers.getContractFactory("IC");
  const ic = await Ic.deploy();
  await ic.deployed();

  console.log(`deployed to ${ic.address}`);
  const result = {
    deployer: ic.signer.address,
    contract: "ic",
    contractAddress: ic.address,
  };
  const jsonString = JSON.stringify(result);

  try {
    await fs.promises.writeFile("config/Ic.contract.json", jsonString);
  } catch (err) {
    console.error("Error writing JSON file:", err);
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
