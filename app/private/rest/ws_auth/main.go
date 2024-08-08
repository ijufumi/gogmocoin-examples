package main

import (
	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	"log"
)

func main() {
	client := rest.New()
	token, err := client.CreateWSAuthToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("token is %v\n", token)
	err = client.ExtendWSAuthToken(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("updated")
	err = client.RevokeWSAuthToken(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finished")
}
