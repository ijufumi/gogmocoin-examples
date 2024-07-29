package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest"
	"log"
)

func main() {
	client := rest.New()

	orderbooks, err := client.OrderBooks(consts.SymbolBCHJPY)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[result]%+v", orderbooks)
}
