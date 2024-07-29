package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest"
	"log"
	"time"
)

func main() {
	client := rest.New()

	for i := 0; i < 5; i++ {
		tradesRes, err := client.Trades(consts.SymbolXRPJPY, int64(i), int64(i))
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("[%v]result:%+v", i, tradesRes)
		time.Sleep(time.Second)
	}
}
