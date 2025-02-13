import fs from 'fs';
import path from 'path';
import { ethers } from 'hardhat';

async function main() {
  const other = fs.readFileSync(`./config/Other.contract.json`);
  const parsedOther = JSON.parse(other);

  const A = await ethers.getContractFactory('A');
  const a = await A.deploy(parsedOther.contractAddress);
  await a.deployed();

  console.log(`deployed to ${a.address}`);
  const result = {
    deployer: a.signer.address,
    contract: 'a',
    contractAddress: a.address,
  };
  const jsonString = JSON.stringify(result);

  try {
    await fs.promises.writeFile('config/A.contract.json', jsonString);
  } catch (err) {
    console.error('Error writing JSON file:', err);
  }
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
