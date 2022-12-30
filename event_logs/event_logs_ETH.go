package event_logs

import (
	"fmt"
	"log"
	"os"

	contract "change_money/gen"
	"change_money/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type ETHSubscription interface {
	WatchSendETHToContract()
	WatchReceivedETHToWallet()
}

type DefaultETHSubscription struct {
	db *gorm.DB
}

func NewETHSubscription(db *gorm.DB) ETHSubscription {
	return &DefaultETHSubscription{
		db: db,
	}
}

func (a DefaultETHSubscription) WatchSendETHToContract() {
	err := godotenv.Load("././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	client, err := ethclient.Dial(os.Getenv("LINK_ETH"))
	if err != nil {
		log.Panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_ETH"))

	t, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Panic(err)
	}

	var opts *bind.WatchOpts
	logs := make(chan *contract.ContractNewTransactionExchange, 10000)
	sub, err := t.WatchNewTransactionExchange(opts, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Panic(err)
		case vLog := <-logs:
			eventLogsETH := models.EventLogs{
				Block_Number:     int64(vLog.Raw.BlockNumber),
				Transaction_Hash: vLog.Raw.TxHash.String(),
				Value:            vLog.Arg2.String(),
				FromAddress:      vLog.Arg0.String(),
				ToAddress:        vLog.Arg1.String(),
				IsReceived:       false,
				FromNetworkId:    1,
				ToNetworkID:      2,
			}
			a.db.Create(&eventLogsETH)
		}
	}
}

func (a DefaultETHSubscription) WatchReceivedETHToWallet() {
	err := godotenv.Load("././.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	client, err := ethclient.Dial(os.Getenv("LINK_ETH"))
	if err != nil {
		log.Panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_ETH"))

	t, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Panic(err)
	}

	var opts *bind.WatchOpts
	logs := make(chan *contract.ContractNewTransactionReceived, 10000)
	sub, err := t.WatchNewTransactionReceived(opts, logs)
	for {
		select {
		case err := <-sub.Err():
			log.Panic(err)
		case vLog := <-logs:
			event_log_BSC := []models.EventLogs{}
			a.db.Where("to_address = ? AND is_received = ? AND from_network_id = ? AND to_network_id = ?", vLog.Arg0.String(), "0", "2", "1").Find(&event_log_BSC)
			for _, event := range event_log_BSC {
				event.IsReceived = true
				a.db.Save(&event)
			}
		}
	}
}
