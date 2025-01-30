import fs from 'fs';
import path from 'path';
import { ethers } from 'hardhat';

async function main() {
  const Create2 = await ethers.getContractFactory('Create2');
  const create2 = await Create2.deploy();
  await create2.deployed();

  console.log(`deployed to ${create2.address}`);
  const result = {
    deployer: create2.signer.address,
    contract: 'create2',
    contractAddress: create2.address,
  };
  const jsonString = JSON.stringify(result);

  try {
    await fs.promises.writeFile('config/Create2.contract.json', jsonString);
  } catch (err) {
    console.error('Error writing JSON file:', err);
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
