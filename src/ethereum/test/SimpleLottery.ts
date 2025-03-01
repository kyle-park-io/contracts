import {
  time,
  loadFixture,
} from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { anyValue } from '@nomicfoundation/hardhat-chai-matchers/withArgs';
import { expect } from 'chai';
import hre from 'hardhat';
import * as ethers from 'ethers';
import { setTimeout } from 'timers/promises';

describe('Lottery Contract Tests with hardhat', function () {
  async function deployLotteryContractFixture() {
    // Contracts are deployed using the first signer/account by default
    const [owner, account1, account2, account3, account4] =
      await hre.ethers.getSigners();

    const Lottery = await hre.ethers.getContractFactory('Lottery');
    const lottery = await Lottery.connect(owner).deploy();
    const lotteryAddress = await lottery.getAddress();

    return {
      lottery,
      lotteryAddress,
      owner,
      account1,
      account2,
      account3,
      account4,
    };
  }

  describe('Deployment', function () {
    it('정상적인 컨트랙트 배포', async function () {
      const { lotteryAddress } = await loadFixture(
        deployLotteryContractFixture,
      );
      console.log('lottery contract address: ', lotteryAddress);
    });
  });

  describe('Enter Lottery!', function () {
    it('플레이어 투입', async function () {
      const { lottery, account1, account2, account3, account4 } =
        await loadFixture(deployLotteryContractFixture);

      await lottery
        .connect(account1)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account2)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account3)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account4)
        .enter({ value: ethers.parseEther('0.001') });

      const randomValue = await lottery.randomPublic();
      console.log(`Value in Random Func: ${randomValue}`);
    });

    it('3회 초과 입장 플레이어', async function () {
      const { lottery, account1 } = await loadFixture(
        deployLotteryContractFixture,
      );

      await lottery
        .connect(account1)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account1)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account1)
        .enter({ value: ethers.parseEther('0.001') });

      // revert
      await expect(
        lottery.connect(account1).enter({ value: ethers.parseEther('0.001') }),
      ).to.be.revertedWith('최대 3회까지만 참여 가능');
    });

    it('당첨자 선정 및 관련 데이터 조회', async function () {
      const { lottery, lotteryAddress, owner, account1, account2, account3 } =
        await loadFixture(deployLotteryContractFixture);

      await lottery
        .connect(account1)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account2)
        .enter({ value: ethers.parseEther('0.001') });
      await lottery
        .connect(account3)
        .enter({ value: ethers.parseEther('0.001') });

      // 1라운드 당첨자 선정 전 데이터 조회
      let round = await lottery.rounds();
      console.log('🔹 라운드: ', round);
      let prizeAmounts = await lottery.prizeAmount();
      console.log('🔹 라운드 당첨금: ', prizeAmounts);
      let players = await lottery.getPlayers();
      console.log('🔹 전체 플레이어: ', players);

      // 1라운드 당첨자 선정
      // create winner tx
      await expect(lottery.connect(owner).pickWinner()).not.to.be.reverted;

      // 2라운드 데이터 조회
      round = await lottery.rounds();
      console.log('🔹 라운드: ', round);
      prizeAmounts = await lottery.prizeAmount();
      console.log('🔹 라운드 당첨금: ', prizeAmounts);
      const winners = await lottery.getWinners(round - BigInt(1));
      console.log('🔹 우승자: ', winners);
      players = await lottery.getPlayers();
      console.log('🔹 전체 플레이어: ', players);
    });
  });
});
