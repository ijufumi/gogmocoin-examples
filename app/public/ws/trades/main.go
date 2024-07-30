package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws"
	"log"
	"time"
)

func main() {
	client := ws.NewTrades(consts.SymbolBTCJPY, nil)
	timeoutCnt := 0
	err := client.Subscribe()
	if err != nil {
		log.Fatalln(err)
	}
	for {
		select {
		case v := <-client.Receive():
			log.Printf("msg:%+v\n", v)
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
