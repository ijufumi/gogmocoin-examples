package main

import (
	"log"
	"time"

	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/public/ws/channel/orderbooks"
)

func main() {
	client := orderbooks.New(configuration.SymbolBTCJPY)
	timeoutCnt := 0
	for {
		select {
		case v := <-client.Receive():
			log.Printf("msg:%+v", v)
		case <-time.After(180 * time.Second):
			log.Println("timeout...")
			timeoutCnt++
		}
		if timeoutCnt > 20 {
			break
		}
	}
	e := client.Unsubscribe()
	if e != nil {
		log.Println(e)
		return
	}
}
