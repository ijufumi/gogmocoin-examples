package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	client := private.New()
	response, err := client.PositionSummary(consts.SymbolBTCJPY)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", response)
}
