package rate

import (
	"change_money/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func rate() float64 {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.binance.us/exchange-api/v1/public/asset-service/product/get-products", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	ExchangeMoney := &dto.ExchangeMoney{}
	json.Unmarshal(respBody, ExchangeMoney)
	var ETH, BNB float64
	for _, coin := range ExchangeMoney.Data {
		if coin.S == "ETHUSD" {
			ETH = coin.C
		} else if coin.S == "BNBUSD" {
			BNB = coin.C
		}
	}
	return BNB / ETH
}

func ExchangeRate() float64 {
	Rate := rate()
	return Rate
}
