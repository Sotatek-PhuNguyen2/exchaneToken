package dto

type ExchangeMoney struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          []struct {
		S          string        `json:"s"`
		St         string        `json:"st"`
		B          string        `json:"b"`
		Q          string        `json:"q"`
		Ba         string        `json:"ba"`
		Qa         string        `json:"qa"`
		I          float64       `json:"i"`
		Ts         float64       `json:"ts"`
		An         string        `json:"an"`
		Qn         string        `json:"qn"`
		O          float64       `json:"o"`
		H          float64       `json:"h"`
		L          float64       `json:"l"`
		C          float64       `json:"c"`
		V          float64       `json:"v"`
		Qv         float64       `json:"qv"`
		Y          int           `json:"y"`
		As         float64       `json:"as"`
		Pm         string        `json:"pm"`
		Pn         string        `json:"pn"`
		Cs         int64         `json:"cs"`
		Tags       []interface{} `json:"tags"`
		Pom        bool          `json:"pom"`
		Pomt       interface{}   `json:"pomt"`
		EnableOcbs bool          `json:"enableOcbs"`
		Kline      string        `json:"kline"`
		MarketCap  float64       `json:"marketCap"`
		Blogo      string        `json:"blogo"`
		Qlogo      string        `json:"qlogo"`
		Etf        bool          `json:"etf"`
	} `json:"data"`
	Success bool `json:"success"`
}
