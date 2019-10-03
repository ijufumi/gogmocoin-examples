package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/rest"
)

func main() {
	client := rest.New()

	orderbooks, err := client.OrderBooks(configuration.SymbolBCHJPY)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[result]%+v", orderbooks)
}
