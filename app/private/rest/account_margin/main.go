package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := rest.New()
	accountMarginRes, err := client.AccountMargin()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", accountMarginRes)
}
