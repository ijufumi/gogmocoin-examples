package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/private"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"time"
)

func main() {
	client := private.New()
	executionsRes, err := client.ExecutionsByOrderID(103804777)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", executionsRes)

	for _, execution := range executionsRes.Data.List {
		time.Sleep(time.Second)
		executionsRes, err = client.ExecutionsByExecutionID(execution.ExecutionID)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("result:%+v", executionsRes)
	}
}
