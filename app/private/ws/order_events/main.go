package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/private/ws"
	"log"
	"time"
)

func main() {
	client := ws.NewOrderEvents(true)
	if err := client.Subscribe(); err != nil {
		log.Fatal(err)
	}
	timeoutCnt := 0
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
	if err := client.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}
