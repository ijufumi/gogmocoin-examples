package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := rest.New()
	ordersRes, err := client.Orders(12345676879)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("ordersRes:%+v", ordersRes)
}
