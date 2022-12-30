package dto

type LogEventBSCExchange struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		BlockHash        string   `json:"blockHash"`
		TimeStamp        string   `json:"timeStamp"`
		GasPrice         string   `json:"gasPrice"`
		GasUsed          string   `json:"gasUsed"`
		LogIndex         string   `json:"logIndex"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
	} `json:"result"`
}

type LogEventBSCRecieved struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		BlockHash        string   `json:"blockHash"`
		TimeStamp        string   `json:"timeStamp"`
		GasPrice         string   `json:"gasPrice"`
		GasUsed          string   `json:"gasUsed"`
		LogIndex         string   `json:"logIndex"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
	} `json:"result"`
}
