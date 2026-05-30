package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
)

// orderID は照会対象の注文ID。サンプル用のダミー値なので、
// 実行前に自分の注文IDへ置き換えること（存在しないIDではエラーになる）。
const orderID = 12345676879

func main() {
	client := rest.New()
	ordersRes, err := client.Orders(orderID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", ordersRes)
}
