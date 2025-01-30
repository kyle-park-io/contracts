import fs from 'fs';
import path from 'path';
import { ethers } from 'hardhat';

async function main() {
  const Ie = await ethers.getContractFactory('Example');
  const ie = await Ie.deploy();
  await ie.deployed();

  console.log(`deployed to ${ie.address}`);
  const result = {
    deployer: ie.signer.address,
    contract: 'ie',
    contractAddress: ie.address,
  };
  const jsonString = JSON.stringify(result);

  try {
    await fs.promises.writeFile('config/Ie.contract.json', jsonString);
  } catch (err) {
    console.error('Error writing JSON file:', err);
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
