import fs from "fs";
import path from "path";
import { ethers } from "hardhat";

async function main() {
  const It = await ethers.getContractFactory("InterfaceTest");
  const it = await It.deploy();
  await it.deployed();

  console.log(`deployed to ${it.address}`);
  const result = {
    deployer: it.signer.address,
    contract: "it",
    contractAddress: it.address,
  };
  const jsonString = JSON.stringify(result);

  try {
    await fs.promises.writeFile("config/It.contract.json", jsonString);
  } catch (err) {
    console.error("Error writing JSON file:", err);
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
