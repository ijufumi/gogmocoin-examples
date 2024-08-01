package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	client := rest.New()
	response, err := client.PositionSummary(consts.SymbolBTCJPY)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", response)
}
