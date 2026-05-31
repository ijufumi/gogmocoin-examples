package main

import (
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	_ "github.com/joho/godotenv/autoload"
	"github.com/shopspring/decimal"
)

// 注意: ChangeOrder は実際に注文内容を変更する破壊的な操作。
// 以下の値はサンプル用のダミーであり、必ず自分の有効な注文IDと変更後の価格に
// 置き換えてから実行すること。存在しない注文IDではエラーになる。
const (
	// orderID は変更対象の注文ID。自分の注文IDに置き換える。
	orderID = 134572625
)

func main() {
	client := rest.New()

	// newPrice は変更後の価格。decimal.New(1115001, 0) = 1115001。
	// 取引対象・呼値に応じて妥当な値へ置き換えること。
	newPrice := decimal.New(1115001, 0)

	response, err := client.ChangeOrder(orderID, newPrice)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", response)
}
