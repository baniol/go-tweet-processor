package main

import (
	"go-tweet-processor/db"
	"go-tweet-processor/web"
	"log"
	"net/http"
)

func main() {

	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}

	web.StartServer(dblayer)
	http.ListenAndServe(":1323", nil)
}
