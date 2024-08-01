package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	"log"
)

func main() {
	client := rest.NewWithKeys("YOUR_API_KEY", "YOUR_API_SECRET")
	ordersRes, err := client.Orders(12345676879)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("ordersRes:%+v", ordersRes)
}
