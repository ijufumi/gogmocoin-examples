package main

import (
	"log"
	"time"

	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest"
)

// countPerPage は 1 ページあたりの取得件数。
const countPerPage = 100

func main() {
	client := rest.New()

	// page は 1 始まり。ページを 1, 2, 3... と進めながら取得するサンプル。
	for page := int64(1); page <= 5; page++ {
		tradesRes, err := client.Trades(consts.SymbolXRPJPY, page, countPerPage)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("[page:%v]result:%+v", page, tradesRes)
		// レート制限に配慮して各リクエストの間隔を空ける。
		time.Sleep(time.Second)
	}
}
