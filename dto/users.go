package dto

type UserResquestExchange struct {
	PrivateKey string `json:"privateKey"`
	ToAddress  string `json:"toAddress"`
	Value      int64  `json:"value"`
	Network1   string
	Network2   string
}

type UserResponseExchange struct {
	TXHash string `json:"TXHash"`
}

type UserResquestReceive struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Network1   string
	Network2   string
}

type UserResponseReceive struct {
	TXHash string `json:"TXHash"`
}
