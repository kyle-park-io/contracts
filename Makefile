sol:
	solc --optimize --abi ./contracts/Lock.sol -o build
	solc --optimize --bin ./contracts/Lock.sol -o build
	./abigen --abi=./build/Lock.abi --bin=./build/Lock.bin --pkg=api --type Lock --out=./api/Lock.go

ganache:
	yarn
	npx hardhat run scripts/deploy.other.ts --network ganache
	npx hardhat run scripts/deploy.a.ts --network ganache

sepolia:
	yarn
	npx hardhat run scripts/deploy.mint.ts --network sepolia
	npx hardhat run scripts/deploy.validatorInfo.ts --network sepolia
	npx hardhat run scripts/deploy.validator.ts --network sepolia
	npx hardhat run scripts/deploy.vault.ts --network sepolia