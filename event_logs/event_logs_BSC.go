package event_logs

import (
	"change_money/models"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// // import (
// // 	"change_money/dto"
// // 	"change_money/models"
// // 	"encoding/json"
// // 	"fmt"
// // 	"io/ioutil"
// // 	"log"
// // 	"net/http"
// // 	"os"
// // 	"strconv"
// // 	"strings"
// // 	"time"

// // 	"github.com/joho/godotenv"
// // 	"gorm.io/gorm"
// // )

// // type BSCSubscription interface {
// // 	WatchSendBSCToContract()
// // 	WatchReceivedBSCToWallet()
// // }

// // type DefaultBSCSubscription struct {
// // 	db *gorm.DB
// // }

// // func NewBSCSubscription(db *gorm.DB) BSCSubscription {
// // 	return &DefaultBSCSubscription{
// // 		db: db,
// // 	}
// // }

// // func (a DefaultBSCSubscription) WatchSendBSCToContract() {
// // 	for {
// // 		err := godotenv.Load("./././.env")
// // 		if err != nil {
// // 			fmt.Printf("Error loading .env file")
// // 		}
// // 		eventLogsBSC := models.EventLogs{}
// // 		a.db.Last(&eventLogsBSC)
// // 		EventLogsBSCdto := &dto.LogEventBSCExchange{}
// // 		client := &http.Client{}
// // 		fromBlock := fmt.Sprint(eventLogsBSC.Block_Number + 1)
// // 		contractAddress := os.Getenv("ADDRESS_CONTRACT_BNB")
// // 		request := "https://api-testnet.bscscan.com/api?module=logs&action=getLogs&fromBlock=" + fromBlock + "&address=" + contractAddress + "&topic0=0xe303bfcae5504a87ac4aee4818422409a3da17fecbfc891de411a02ffbacd2a5&apikey=HK9XBMW21IR6VINPA4EWX2MXSEIAIB199J"
// // 		req, err := http.NewRequest("GET", request, nil)
// // 		if err != nil {
// // 			log.Print(err)
// // 			os.Exit(1)
// // 		}
// // 		resp, err := client.Do(req)
// // 		if err != nil {
// // 			fmt.Println("Error sending request to server")
// // 			os.Exit(1)
// // 		}
// // 		respBody, _ := ioutil.ReadAll(resp.Body)
// // 		json.Unmarshal(respBody, EventLogsBSCdto)
// // 		for _, event := range EventLogsBSCdto.Result {
// // 			eventLogsETH := models.EventLogs{
// // 				Block_Number:     HexToInt(event.BlockNumber),
// // 				Transaction_Hash: event.TransactionHash,
// // 				Value:            HexToInt(event.Data[131:]),
// // 				FromAddress:      "0x" + event.Data[26:66],
// // 				ToAddress:        "0x" + event.Data[90:130],
// // 				IsReceived:       false,
// // 				NetWorkId:        1,
// // 			}
// // 			a.db.Create(&eventLogsETH)
// // 		}
// // 		time.Sleep(5 * time.Second)
// // 	}
// // }

// // func (a DefaultBSCSubscription) WatchReceivedBSCToWallet() {
// // 	for {
// // 		err := godotenv.Load("./././.env")
// // 		if err != nil {
// // 			fmt.Printf("Error loading .env file")
// // 		}
// // 		eventLogsBSC := models.EventLogs{}
// // 		a.db.Last(&eventLogsBSC)
// // 		EventLogsBSCdto := &dto.LogEventBSCRecieved{}
// // 		client := &http.Client{}
// // 		fromBlock := fmt.Sprint(eventLogsBSC.Block_Number + 1)
// // 		contractAddress := os.Getenv("ADDRESS_CONTRACT_BNB")
// // 		request := "https://api-testnet.bscscan.com/api?module=logs&action=getLogs&fromBlock=" + fromBlock + "&address=" + contractAddress + "&topic0=0x4f0222c6a15da55907b401d3d25463ff0c3b38e50ca641d12c43aa4a85c9b9c5&apikey=HK9XBMW21IR6VINPA4EWX2MXSEIAIB199J"
// // 		req, err := http.NewRequest("GET", request, nil)
// // 		if err != nil {
// // 			log.Print(err)
// // 			os.Exit(1)
// // 		}
// // 		resp, err := client.Do(req)
// // 		if err != nil {
// // 			fmt.Println("Error sending request to server")
// // 			os.Exit(1)
// // 		}
// // 		respBody, _ := ioutil.ReadAll(resp.Body)
// // 		json.Unmarshal(respBody, EventLogsBSCdto)
// // 		event_log_ETh := []models.EventLogs{}
// // 		for _, event := range EventLogsBSCdto.Result {
// // 			a.db.Where("to_address = ? AND is_received = ?", "0x"+event.Data[26:66], "0").Find(&event_log_ETh)
// // 			for _, event := range event_log_ETh {
// // 				event.IsReceived = true
// // 				a.db.Save(&event)
// // 			}
// // 		}
// // 		time.Sleep(5 * time.Second)
// // 	}
// // }

// func HexToInt(hex string) int64 {
// 	numberStr := strings.Replace(hex, "0x", "", -1)
// 	numberValue, err := strconv.ParseInt(numberStr, 16, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0
// 	}
// 	return numberValue
// }

// package event_logs

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	contract "change_money/gen_copy"
// 	"change_money/models"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/joho/godotenv"
// 	"gorm.io/gorm"
// )

type BSCSubscription interface {
	WatchEventContract()
}

type DefaultBSCSubscription struct {
	db *gorm.DB
}

func NewBSCSubscription(db *gorm.DB) BSCSubscription {
	return &DefaultBSCSubscription{
		db: db,
	}
}

// func (a DefaultBSCSubscription) WatchSendBSCToContract() {
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
// 	logs := make(chan *contract.ContractNewTransactionExchange, 10000)
// 	sub, err := t.WatchNewTransactionExchange(opts, logs)
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Panic(err)
// 		case vLog := <-logs:
// 			eventLogsBSC := models.EventLogs{
// 				Block_Number:      int64(vLog.Raw.BlockNumber),
// 				Transaction_Hash:  vLog.Raw.TxHash.String(),
// 				Value:             vLog.Arg2.String(),
// 				FromAddress:       vLog.Arg0.String(),
// 				ToAddress:         vLog.Arg1.String(),
// 				IsReceived:        false,
// 				NetWorkId:         2,
// 				NetWorkRecievedId: 1,
// 			}
// 			a.db.Create(&eventLogsBSC)
// 		}
// 	}
// }

// func (a DefaultBSCSubscription) WatchReceivedBSCToWallet() {
// 	err := godotenv.Load("././.env")
// 	if err != nil {
// 		fmt.Printf("Error loading .env file")
// 	}

// 	client, err := ethclient.Dial(os.Getenv("LINK_BSC_WSS"))
// 	if err != nil {
// 		//fmt.Println(err)
// 	}

// 	contractAddress := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_BNB"))

// 	t, err := contract.NewContract(contractAddress, client)
// 	if err != nil {
// 		//fmt.Println(err)
// 	}

// 	var opts *bind.WatchOpts
// 	logs := make(chan *contract.ContractNewTransactionReceived, 10000)
// 	sub, err := t.WatchNewTransactionReceived(opts, logs)
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Panic(err)
// 		case vLog := <-logs:
// 			event_log_BSC := []models.EventLogs{}
// 			fmt.Println(vLog.Arg0)
// 			a.db.Where("to_address = ? AND is_received = ? AND net_work_id = ? AND net_work_recieved_id = ?", vLog.Arg0.Hex(), "0", 1, 2).Find(&event_log_BSC)
// 			for _, event := range event_log_BSC {
// 				event.IsReceived = true
// 				a.db.Save(&event)
// 			}
// 		}
// 	}
// }

func (a DefaultBSCSubscription) WatchEventContract() {
	for {
		err := godotenv.Load("././.env")
		if err != nil {
			fmt.Printf("Error loading .env file")
		}

		client, err := ethclient.Dial(os.Getenv("LINK_BSC_WSS"))
		if err != nil {
			fmt.Println(err)
			break
		}

		//topic := common.HexToHash("0x3434032ea91078551b2d5e559bd385096c32e1d79b46cb703ff913f5bf106574")
		contractAddress := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_BNB"))
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			fmt.Println(err)
			break
		}

		for {
			select {
			case <-sub.Err():
				break
			case vLog := <-logs:
				if vLog.Topics[0].String() == "0x3434032ea91078551b2d5e559bd385096c32e1d79b46cb703ff913f5bf106574" {
					value := new(big.Int)
					event := hex.EncodeToString(vLog.Data)
					fmt.Println(event)
					value.SetString(event[129:], 16)
					eventLogsETH := models.EventLogs{
						Block_Number:     int64(vLog.BlockNumber),
						Transaction_Hash: vLog.TxHash.Hex(),
						Value:            value.String(),
						FromAddress:      "0x" + event[24:64],
						ToAddress:        "0x" + event[88:128],
						IsReceived:       false,
						FromNetworkId:    2,
						ToNetworkID:      1,
					}
					a.db.Create(&eventLogsETH)
				} else if vLog.Topics[0].String() == "0x4f0222c6a15da55907b401d3d25463ff0c3b38e50ca641d12c43aa4a85c9b9c5" {
					event_log_BSC := []models.EventLogs{}
					event := hex.EncodeToString(vLog.Data)
					toAddress := "0x" + event[24:]
					a.db.Where("to_address = ? AND is_received = ? AND from_network_id = ? AND to_network_id = ?", toAddress, "0", 1, 2).Find(&event_log_BSC)
					for _, event := range event_log_BSC {
						event.IsReceived = true
						a.db.Save(&event)
					}
				}
			}
			break
		}
	}
}
