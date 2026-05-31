package main

import (
	"log"

	"github.com/ijufumi/gogmocoin-examples/internal/wsrunner"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/ws"
)

func main() {
	client := ws.NewTrades(consts.SymbolBTCJPY, nil)
	if err := wsrunner.Run(client.Subscribe, client.Receive, client.Unsubscribe); err != nil {
		log.Fatal(err)
	}
}
