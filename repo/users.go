package repo

import (
	"change_money/admin"
	"change_money/dto"
	"change_money/errs"
	contract "change_money/gen"
	"change_money/models"
	"change_money/rate"
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type UserRepo interface {
	sendTokenToContract(dto.UserResquestExchange, dto.Contract) (dto.UserResponseExchange, *errs.AppError)
	//sendETHToContract(req dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError)
	receiveToken(dto.UserResquestReceive, dto.Contract) (dto.UserResponseReceive, *errs.AppError)
	// receiveETH(dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError)
	ExchangeController(dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError)
	ReceiveController(dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError)
	GetAllTokenFromAddress(PublicAddress string, network1 string, network2 string) (*big.Int, *errs.AppError)
}

type DefaultUserRepo struct {
	db    *gorm.DB
	admin *admin.Admin
}

func NewUserRepo(db *gorm.DB, admin *admin.Admin) UserRepo {
	return DefaultUserRepo{
		db:    db,
		admin: admin,
	}
}

func (r DefaultUserRepo) ExchangeController(req dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError) {
	c := dto.Contract{}
	err := godotenv.Load("././.env")
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorGetData()
	}
	if req.Network2 != "eth" && req.Network2 != "bsc" {
		return dto.UserResponseExchange{}, errs.BadRequestError("From Network or To Network does not exist")
	}
	switch req.Network1 {
	case "eth":
		c.Address_Contract = os.Getenv("ADDRESS_CONTRACT_ETH")
		c.Link_Contract = os.Getenv("LINK_ETH")
		return r.sendTokenToContract(req, c)
	case "bsc":
		c.Address_Contract = os.Getenv("ADDRESS_CONTRACT_BNB")
		c.Link_Contract = os.Getenv("LINK_BSC")
		return r.sendTokenToContract(req, c)
	default:
		return dto.UserResponseExchange{}, errs.BadRequestError("From Network or To Network does not exist")
	}
}

func (r DefaultUserRepo) ReceiveController(req dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError) {
	c := dto.Contract{}
	err := godotenv.Load("././.env")
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorGetData()
	}
	if req.Network1 != "eth" && req.Network1 != "bsc" {
		return dto.UserResponseReceive{}, errs.BadRequestError("From Network or To Network does not exist")
	}
	switch req.Network2 {
	case "eth":
		c.Address_Contract = os.Getenv("ADDRESS_CONTRACT_ETH")
		c.Link_Contract = os.Getenv("LINK_ETH")
		return r.receiveToken(req, c)
	case "bsc":
		c.Address_Contract = os.Getenv("ADDRESS_CONTRACT_BNB")
		c.Link_Contract = os.Getenv("LINK_BSC")
		return r.receiveToken(req, c)
	default:
		return dto.UserResponseReceive{}, errs.BadRequestError("From Network or To Network does not exist")
	}
}

func (r DefaultUserRepo) sendTokenToContract(req dto.UserResquestExchange, c dto.Contract) (dto.UserResponseExchange, *errs.AppError) {
	client, err := ethclient.Dial(c.Link_Contract)
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorGetData()
	}
	defer client.Close()

	cAdd := common.HexToAddress(c.Address_Contract)

	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorInsertData()
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorInsertData()
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To: &cAdd,
	})
	if err == nil {
		gasLimit += 20000 * params.Wei
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorInsertData()
	}

	t, err := contract.NewContract(cAdd, client)
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorInsertData()
	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorInsertData()
	}

	tx.GasLimit = uint64(float64(gasLimit))
	tx.GasPrice = gasPrice
	tx.Value = big.NewInt(req.Value)

	tra, err := t.SendTokenToContract(tx, common.HexToAddress(req.ToAddress))
	if err != nil {
		return dto.UserResponseExchange{}, errs.ErrorInsertData()
	}

	return dto.UserResponseExchange{TXHash: tra.Hash().Hex()}, nil
}

func (r DefaultUserRepo) receiveToken(req dto.UserResquestReceive, c dto.Contract) (dto.UserResponseReceive, *errs.AppError) {
	client, err := ethclient.Dial(c.Link_Contract)
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorGetData()
	}
	defer client.Close()

	addressPublic := common.HexToAddress(req.PublicKey)
	cAdd := common.HexToAddress(c.Address_Contract)

	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorInsertData()
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorInsertData()
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To: &cAdd,
	})
	if err == nil {
		gasLimit += 20000 * params.Wei
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorInsertData()
	}

	t, err := contract.NewContract(cAdd, client)
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorInsertData()
	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorInsertData()
	}

	tx.GasLimit = uint64(float64(gasLimit))
	tx.GasPrice = gasPrice

	valueOfAddress, err1 := r.GetAllTokenFromAddress(req.PublicKey, req.Network1, req.Network2)
	if err1 != nil {
		return dto.UserResponseReceive{}, errs.ErrorInsertData()
	}

	if valueOfAddress.IsInt64() == true {
		if valueOfAddress.Int64() == 0 {
			return dto.UserResponseReceive{}, errs.ErrorReadData()
		}
	}

	//fmt.Println(valueOfAddress)
	signature := r.admin.Signature(req.PublicKey, valueOfAddress)
	fmt.Println(hexutil.Encode(signature))
	tra, err := t.ReceiveToken(tx, signature, addressPublic, valueOfAddress)
	if err != nil {
		return dto.UserResponseReceive{}, errs.ErrorUpdateData()
	}

	return dto.UserResponseReceive{TXHash: tra.Hash().Hex()}, nil
}

func (r DefaultUserRepo) GetAllTokenFromAddress(PublicAddress string, network1 string, network2 string) (*big.Int, *errs.AppError) {
	value := big.NewInt(0)
	event_log := []models.EventLogs{}
	network_id_1 := models.NetWork{}
	network_id_2 := models.NetWork{}
	r.db.Where("name_of_network = ?", network1).First(&network_id_1)
	r.db.Where("name_of_network = ?", network2).First(&network_id_2)
	r.db.Where("to_address = ? AND is_received = ? AND from_network_id = ? AND to_network_id = ?", PublicAddress, "0", network_id_1.Id, network_id_2.Id).Find(&event_log)
	for _, event := range event_log {
		n := new(big.Int)
		n, ok := n.SetString(event.Value, 10)
		if !ok {
			fmt.Println("SetString: error")
			return nil, errs.ErrorGetData()
		}
		fmt.Println(n.String())
		value = value.Add(value, n)
	}
	rate := rate.ExchangeRate()
	if network2 == "eth" {
		RateInt := int64(rate * 90000000000)
		rateBigInt := big.NewInt(RateInt)
		value = value.Mul(value, rateBigInt)
		value = value.Div(value, big.NewInt(100000000000))
	} else if network2 == "bsc" {
		RateInt := int64(rate * 100000000000)
		rateBigInt := big.NewInt(RateInt)
		value = value.Div(value, rateBigInt)
		value = value.Mul(value, big.NewInt(90000000000))
	}
	return value, nil
}

// func GetAllBSCFromAddress(PublicAddress string) (uint64, *errs.AppError) {
// 	var value uint64
// 	event_log_BSC := []models.EventLogs{}
// 	config.DB.Where("to_address = ? AND is_received = ? AND net_work_id = ? AND net_work_recieved_id = ?", PublicAddress, "0").Find(&event_log_BSC)
// 	for _, event := range event_log_BSC {
// 		value += uint64(event.Value)
// 	}
// 	if value == 0 {
// 		return 0, errs.ErrorDataNotSurvive()
// 	}
// 	return value, nil
// }

// func (r DefaultUserRepo) sendETHToContract(req dto.UserResquestExchange) (dto.UserResponseExchange, *errs.AppError) {
// 	err := godotenv.Load("././.env")
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorGetData()
// 	}

// 	client, err := ethclient.Dial(os.Getenv("LINK_ETH"))
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorGetData()
// 	}
// 	defer client.Close()

// 	cAdd := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_ETH"))

// 	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorInsertData()
// 	}

// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorInsertData()
// 	}

// 	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 		To: &cAdd,
// 	})
// 	if err == nil {
// 		gasLimit += 20000 * params.Wei
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorInsertData()
// 	}

// 	t, err := contract.NewContract(cAdd, client)
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorInsertData()
// 	}

// 	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorInsertData()
// 	}

// 	tx.GasLimit = uint64(float64(gasLimit))
// 	tx.GasPrice = gasPrice
// 	tx.Value = big.NewInt(req.Value)

// 	tra, err := t.SendTokenToContract(tx, common.HexToAddress(req.ToAddress))
// 	if err != nil {
// 		return dto.UserResponseExchange{}, errs.ErrorInsertData()
// 	}

// 	return dto.UserResponseExchange{TXHash: tra.Hash().Hex()}, nil
// }

// func (r DefaultUserRepo) receiveETH(req dto.UserResquestReceive) (dto.UserResponseReceive, *errs.AppError) {
// 	signature := r.admin.Signature()
// 	message := r.admin.GetHashMessage().Bytes()
// 	var messageHash [32]byte

// 	copy(messageHash[:], message)

// 	err := godotenv.Load("././.env")
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorGetData()
// 	}

// 	client, err := ethclient.Dial(os.Getenv("LINK_ETH"))
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorGetData()
// 	}
// 	defer client.Close()
// 	addressPublic := common.HexToAddress(req.PublicKey)
// 	cAdd := common.HexToAddress(os.Getenv("ADDRESS_CONTRACT_ETH"))

// 	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorInsertData()
// 	}

// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorInsertData()
// 	}

// 	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 		To: &cAdd,
// 	})
// 	if err == nil {
// 		gasLimit += 20000 * params.Wei
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorInsertData()
// 	}

// 	t, err := contract.NewContract(cAdd, client)
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorInsertData()
// 	}

// 	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorInsertData()
// 	}

// 	tx.GasLimit = uint64(float64(gasLimit))
// 	tx.GasPrice = gasPrice
// 	fmt.Println(GetAllBSCFromAddress(req.PublicKey))

// 	valueOfAddress, err1 := GetAllBSCFromAddress(req.PublicKey)
// 	if err1 != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorInsertData()
// 	}

// 	value := uint64(float64(valueOfAddress) * rate.ExchangeRate())
// 	fmt.Println(value)
// 	fmt.Println(req.PublicKey)

// 	tra, err := t.ReceiveToken(tx, messageHash, signature, addressPublic, value)
// 	if err != nil {
// 		return dto.UserResponseReceive{}, errs.ErrorUpdateData()
// 	}

// 	return dto.UserResponseReceive{TXHash: tra.Hash().Hex()}, nil
// }
