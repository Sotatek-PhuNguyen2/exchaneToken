// package main

// import "change_money/app"

// func main() {
// 	app.SetUpApp()
// }

package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.HexToECDSA("342c8bcb3a1f8bc48aea7b82a2cbcb9753d8c10357a5163a957033bc9fa3ba1a")
	if err != nil {
		log.Fatal(err)
	}
	value := new(big.Int)
	value, ok := value.SetString("10000000000000000000", 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	address1 := common.HexToAddress("0xC4E12b4DB90D2042D70E42B6ef6D025cbCe5692b")
	value1 := common.BigToHash(value)
	network := []byte("ethereum")
	nonce := big.NewInt(3)
	nonce1 := common.BigToHash(nonce)
	messageHash := crypto.Keccak256Hash(address1.Bytes(), value1.Bytes(), network, nonce1.Bytes())
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

	// data := []byte("longphu")
	// data1 := []byte("\x19Ethereum Signed Message:\n32")
	// hash := crypto.Keccak256Hash(data)
	// hash1 := crypto.Keccak256Hash(data1, hash.Bytes())

	// signature := []byte("0x7fd504b69dddd77bef24d3bfc426c42a5890273af54547397deec9bf0dd0fcc534849845d8b9657cc19acf43055ae7aaedb8de6c1cedf22dcc4077ea5a42c7461c")
	// signature[64] -= 27

}

// // publicKeyBytes := []byte("0x6ee1725C1BcbAE17fBcE9FFC5fE4A23d627C6e35")

// // 	data := []byte("longphu")
// // 	data1 := []byte("\x19Ethereum Signed Message:\n32")
// // 	hash := crypto.Keccak256Hash(data)
// // 	hash1 := crypto.Keccak256Hash(data1, hash.Bytes())

// // 	signature := []byte("0x7fd504b69dddd77bef24d3bfc426c42a5890273af54547397deec9bf0dd0fcc534849845d8b9657cc19acf43055ae7aaedb8de6c1cedf22dcc4077ea5a42c7461c")
// // 	signature[64] -= 27

// // 	sigPublicKey, err := crypto.Ecrecover(hash1.Bytes(), signature)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
// // 	fmt.Println(matches) // true

// // 	sigPublicKeyECDSA, err := crypto.SigToPub(hash1.Bytes(), signature)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
// // 	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
// // 	fmt.Println(matches) // true

// // 	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
// // 	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
// // 	fmt.Println(verified) // true

// // package main

// // import (
// // 	"encoding/hex"
// // 	"fmt"
// // 	"log"
// // )

// // func main() {
// // 	const s = "48656c6c6f20476f7068657221"
// // 	decoded, err := hex.DecodeString(s)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	fmt.Printf("%x\n", decoded)

// //}
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"change_money/gen"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load("././.env")
// 	if err != nil {
// 		fmt.Printf("Error loading .env file")
// 	}

// 	client, err := ethclient.Dial(os.Getenv("LINK_BSC_WSS"))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	contractAddress := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_BNB"))

// 	t, err := contract.NewContract(contractAddress, client)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var opts *bind.WatchOpts
// 	logs := make(chan *contract.ContractNewTransaction, 10000)
// 	sub, err := t.WatchNewTransaction(opts, logs)
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Panic(err)
// 		case vLog := <-logs:
// 			fmt.Println(vLog)
// 		}
// 	}
// }
