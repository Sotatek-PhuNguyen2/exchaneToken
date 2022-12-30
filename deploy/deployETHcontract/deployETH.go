package main

import (
	contract "change_money/gen"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	client, err := ethclient.Dial(os.Getenv("LINK_ETH"))
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_PRIVATE_KEY"))
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := contract.DeployContract(auth, client)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}
