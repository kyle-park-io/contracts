import { HardhatUserConfig } from 'hardhat/config';
import '@nomicfoundation/hardhat-toolbox';
import * as dotenv from 'dotenv';
dotenv.config();

const INFURA_API_KEY = process.env.INFURA_API_KEY;
const GANACHE_DEPLOYER_PRIVATE_KEY = process.env.GANACHE_DEPLOYER_PRIVATE_KEY;
const SEPOLIA_PRIVATE_KEY = process.env.SIGNER_PRIVATE_KEY;

const config: HardhatUserConfig = {
  defaultNetwork: 'hardhat',
  networks: {
    hardhat: {
      chainId: 1337,
    },
    ganache: {
      url: 'http://127.0.0.1:7545',
      accounts: [GANACHE_DEPLOYER_PRIVATE_KEY],
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_API_KEY}`,
      accounts: [SEPOLIA_PRIVATE_KEY],
    },
  },
  solidity: {
    compilers: [
      { version: '0.8.20' },
      { version: '0.4.24' },
      { version: '0.6.12' },
    ],
    settings: {
      optimizer: {
        enabled: true,
        runs: 2000,
      },
      viaIR: true,
    },
  },
  paths: {
    sources: './src/ethereum/contracts',
    tests: './src/ethereum/test',
    cache: './src/ethereum/cache',
    artifacts: './src/ethereum/artifacts',
  },
  mocha: {
    timeout: 40000,
  },
};

export default config;
