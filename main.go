package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var symbol string
var amount string
var offset float64
var side string
var env string
var reoccur int

func init() {
	flag.StringVar(&symbol, "s", "", "SYMBOL: The symbol for the new order")
	flag.StringVar(&amount, "a", "", "AMOUNT: Quoted decimal amount to purchase")
	flag.Float64Var(&offset, "o", 0, "PRICE OFFSET: Quoted decimal amount to ADD TO PRICE")
	flag.StringVar(&side, "t", "buy", "TYPE: buy or sell")
	flag.StringVar(&env, "e", "sand", "ENVIRONMENT: prod or sand")
	flag.IntVar(&reoccur, "r", 0, `REOCCUR: frequency in hours of reocurrence (default "0")`)
}

func main() {
	flag.Parse()

	var baseurl string

	switch env {
	case "prod":
		baseurl = "https://api.gemini.com"
	case "sand":
		baseurl = "https://api.sandbox.gemini.com"
	default:
		fmt.Println(`enter a value of either "prod" or "sand".`)
	}

	for {
		price := PriceFeed(symbol, baseurl, offset)

		payload, err := PayloadBuilder(symbol, amount, price, side)
		if err != nil {
			fmt.Print(err)
			return
		}

		signature := SigBuilder(payload)

		response, err := NewOrder(baseurl, payload, signature)
		if err != nil {
			fmt.Print(err)
			return
		}

		log.Printf("%+v\n", response)

		if reoccur <= 0 {
			return
		} else {
			time.Sleep(time.Hour * time.Duration(reoccur))
		}
	}
}
