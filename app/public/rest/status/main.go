package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/public/rest"
)

func main() {
	client := rest.New()
	status, err := client.Status()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", status)
}
