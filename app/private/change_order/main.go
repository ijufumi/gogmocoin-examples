package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/api/private"
	_ "github.com/joho/godotenv/autoload"
	"github.com/shopspring/decimal"
)

func main() {
	client := private.New()
	response, err := client.ChangeOrder(134572625, decimal.New(1115001, 0))
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", response)
}
