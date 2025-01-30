sol:
  mkdir -p api
	npx solc --optimize --abi ./src/contracts/string/String.sol -o build
  npx solc --optimize --bin ./src/contracts/string/String.sol -o build
	./bin/abigen_arm --abi=./build/src_contracts_string_String_sol_String.abi --bin=./build/src_contracts_string_String_sol_String.bin --pkg=api --type String --out=./api/String.go
  ./bin/abigen_amd --abi=./build/src_contracts_string_String_sol_String.abi --bin=./build/src_contracts_string_String_sol_String.bin --pkg=api --type String --out=./api/String.go

ganache:
	yarn
	npx hardhat run scripts/call/deploy.other.ts --network ganache
	npx hardhat run scripts/call/deploy.a.ts --network ganache
	npx hardhat run scripts/interface/deploy.ie.ts --network ganache
	npx hardhat run scripts/interface/deploy.ic.ts --network ganache
	npx hardhat run scripts/interface/deploy.it.ts --network ganache

sepolia:
	yarn
	npx hardhat run scripts/deploy.mint.ts --network sepolia
	npx hardhat run scripts/deploy.validatorInfo.ts --network sepolia
	npx hardhat run scripts/deploy.validator.ts --network sepolia
	npx hardhat run scripts/deploy.vault.ts --network sepolia
