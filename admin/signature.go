package admin

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

func (a Admin) Signature(address string, value *big.Int) []byte {
	err := godotenv.Load("./././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("ADMIN_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	address1 := common.HexToAddress(address)
	value1 := common.BigToHash(value)
	messageHash := crypto.Keccak256Hash(address1.Bytes(), value1.Bytes())
	data1 := []byte("\x19Ethereum Signed Message:\n32")
	fmt.Println(messageHash.Hex())
	hash1 := crypto.Keccak256Hash(data1, messageHash.Bytes())
	fmt.Println(hash1.Hex())
	signature, err := crypto.Sign(hash1.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	signature[64] += 27
	fmt.Println(hexutil.Encode(signature))
	return signature
}
