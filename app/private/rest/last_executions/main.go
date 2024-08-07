package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := rest.New()
	executionsRes, err := client.LastExecutions(consts.SymbolBTCJPY, 0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", executionsRes)
}
