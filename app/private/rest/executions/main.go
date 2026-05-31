package main

import (
	"log"
	"time"

	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
)

// orderID は約定照会の対象となる注文ID。サンプル用のダミー値なので、
// 実行前に自分の注文IDへ置き換えること（存在しないIDではエラーになる）。
const orderID = 103804777

func main() {
	client := rest.New()
	executionsRes, err := client.ExecutionsByOrderID(orderID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", executionsRes)

	for _, execution := range executionsRes.Data.List {
		// レート制限に配慮して各リクエストの間隔を空ける。
		time.Sleep(time.Second)
		// 1件の失敗で全体を止めないよう、エラー時は次の約定IDへ進む。
		byID, err := client.ExecutionsByExecutionID(execution.ExecutionID)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("result:%+v", byID)
	}
}
