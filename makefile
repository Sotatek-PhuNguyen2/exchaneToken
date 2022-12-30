build: 
	solcjs --bin --abi contract/change_money.sol -o build

gen:
	abigen --bin=build/contract_change_money_sol_change_money.bin --abi=build/contract_change_money_sol_change_money.abi --pkg=contract --out=gen/contract_change.go

deployBNB:
	go run deploy/deployBNBcontract/deployBNB.go

deployETh:
	go run deploy/deployETHcontract/deployETH.go