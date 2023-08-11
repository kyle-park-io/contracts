import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("Create2", function () {
  async function init() {
    const [admin] = await ethers.getSigners();
    const Create2Contract = await ethers.getContractFactory("Create2");
    const create2 = await Create2Contract.deploy();
    await create2.deployed();

    return { create2, admin };
  }

  describe("Calculate contract address", function () {
    it("Test", async function () {
      const { create2, admin } = await loadFixture(init);
      await create2.connect(admin).Calc();
    });
  });
});
