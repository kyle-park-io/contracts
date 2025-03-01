//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import 'hardhat/console.sol';

contract Lottery {
  address public owner;
  address[] public players;

  mapping(address => uint256) public entryCount; // 주소 당 enter 호출 횟수 기록

  struct WinnerInfo {
    address[3] winners; // 3명의 우승자
    uint256 prizeAmount; // 전체 당첨 금액
  }
  uint256 public rounds; // 현재 라운드
  uint256 public prizeAmount; // 라운드 당첨금
  mapping(uint256 => mapping(address => bool)) public hasEntered; // 라운드 별 주소 참여 확인 맵
  mapping(uint256 => WinnerInfo) public roundWinners; // 각 라운드의 3명의 우승자 저장
  mapping(address => uint256) public winnerCount; // 주소 당 우승 횟수

  constructor() {
    owner = msg.sender;
    // 1라운드 시작
    rounds = 1;
  }

  // 1. enter 함수 수정
  function enter() public payable {
    require(msg.value == .001 ether, unicode'정확한 0.001 ether value 전송 필');
    require(entryCount[msg.sender] < 3, unicode'최대 3회까지만 참여 가능');

    // 플레이어 배열 중복 값 배제
    if (!hasEntered[rounds][msg.sender]) {
      players.push(msg.sender);
    }
    prizeAmount += msg.value;
    entryCount[msg.sender]++; // 호출 횟수 증가
  }

  // 2. 리턴 값 테스트를 위해 public으로 변경 함수 추가
  function randomPublic() public view returns (uint, uint) {
    // players 배열 길이 0 modifier
    require(players.length > 0, 'No players in the lottery');

    console.log('In Contract');
    uint value = uint(keccak256(abi.encode(block.timestamp, players)));
    console.log(unicode'🔹 random value = ', value);
    uint index = value % players.length;
    console.log(unicode'🔹 random index = ', index);
    console.log('In Contract End');

    return (value, index);
  }

  function random() private view returns (uint) {
    return uint(keccak256(abi.encode(block.timestamp, players)));
  }

  // 3. 당첨 함수 수정
  function pickWinner() public onlyOwner {
    // 최소 3인의 플레이어 조건 추가
    require(players.length >= 3, 'At least 3 participants are required.');

    // 주어진 과제에서 index-1, index-2 로 추가 당첨자를 선정하나,
    // 0, 1번의 index가 당첨된 경우에 대한 언급이 따로 없어 0, 1번의 당첨의 경우 2번이 당첨된 것으로 판단.
    //    2 +
    // (uint(keccak256(abi.encode(block.timestamp, players))) %
    // (players.length - 2)); 이런 식으로 랜덤 함수의 로직 개선을 시도해볼 수 있으나 실용적인지에 대한 판단은 재고 필.
    uint index = random() % players.length;
    if (index < 2) {
      index = 2;
    }

    console.log('In Contract');
    console.log(unicode'🔹 당첨자 index = ', index);
    console.log(unicode'🔹 총 당첨금 = ', address(this).balance);
    console.log('In Contract End');

    // 블록 데이터 저장
    roundWinners[rounds] = WinnerInfo(
      [players[index], players[index - 1], players[index - 2]],
      prizeAmount
    );
    winnerCount[players[index]]++;
    winnerCount[players[index - 1]]++;
    winnerCount[players[index - 2]]++;
    rounds++;

    // 3명의 당첨자에게 이더 전송
    uint256 prizePerWinner = address(this).balance / 3;
    prizeAmount -= address(this).balance;
    payable(players[index]).transfer(prizePerWinner);
    payable(players[index - 1]).transfer(prizePerWinner);
    payable(players[index - 2]).transfer(prizePerWinner);

    players = new address[](0);
  }

  // 혹시 모르는 상황을 위해 구현했지만.. 필요한가?
  function withdraw() public payable onlyOwner {
    payable(msg.sender).transfer(address(this).balance);
  }

  function transferOwnership(address _newOwner) public onlyOwner {
    owner = _newOwner;
  }

  modifier onlyOwner() {
    _;
    require(msg.sender == owner, 'only Owner!');
  }

  /// Getter
  // get ETH
  function getETH() public view returns (uint256) {
    return address(this).balance;
  }

  // winners getter 함수 추가
  function getWinners(uint256 index) public view returns (WinnerInfo memory) {
    return roundWinners[index];
  }

  // players getter 함수 추가
  function getPlayers() public view returns (address[] memory) {
    return players;
  }
}
