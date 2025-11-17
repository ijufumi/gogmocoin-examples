package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := rest.New()
	accountAssetsRes, err := client.AccountAssets()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", accountAssetsRes)
}
