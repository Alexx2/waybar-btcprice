package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BinancePrice struct {
	Price string `json:"price"`
}

func main() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
	if err != nil {
		fmt.Printf("{\"text\":\"ERR\"}\n")
		return
	}
	defer resp.Body.Close()

	var price BinancePrice
	err = json.NewDecoder(resp.Body).Decode(&price)
	if err != nil {
		fmt.Printf("{\"text\":\"ERR\"}\n")
		return
	}

	priceFloat, err := strconv.ParseFloat(price.Price, 64)
	if err != nil {
		fmt.Printf("{\"text\":\"ERR\"}\n")
		return
	}

	fmt.Printf("· ₿:%.0f\n", priceFloat)
}
