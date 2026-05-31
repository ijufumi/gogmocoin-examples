package main

import (
	"log"

	"github.com/ijufumi/gogmocoin-examples/internal/wsrunner"
	"github.com/ijufumi/gogmocoin/v2/api/private/ws"
)

func main() {
	client := ws.NewPositionEvents(true)
	if err := wsrunner.Run(client.Subscribe, client.Receive, client.Unsubscribe); err != nil {
		log.Fatal(err)
	}
}
