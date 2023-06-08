import fs from "fs";
import path from "path";
import { ethers } from "hardhat";

async function main() {
  const Other = await ethers.getContractFactory("Other");
  const other = await Other.deploy();
  await other.deployed();

  console.log(`deployed to ${other.address}`);
  const result = {
    deployer: other.signer.address,
    contract: "other",
    contractAddress: other.address,
  };
  const jsonString = JSON.stringify(result);

  try {
    await fs.promises.writeFile("config/Other.contract.json", jsonString);
  } catch (err) {
    console.error("Error writing JSON file:", err);
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
