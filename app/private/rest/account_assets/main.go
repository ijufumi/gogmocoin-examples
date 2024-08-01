package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	"log"
)

func main() {
	client := rest.NewWithKeys("YOUR_API_KEY", "YOUR_API_SECRET")
	accountAssetsRes, err := client.AccountAssets()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", accountAssetsRes)
}
