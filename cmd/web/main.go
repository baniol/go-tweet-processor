package main

import (
	"github.com/baniol/go-tweet-processor/api"
	"github.com/baniol/go-tweet-processor/db"
	"github.com/baniol/go-tweet-processor/web"
	"log"
	"net/http"
)

func main() {

	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}

	// web.InitHandlers(dblayer)
	// TODO: Graceful shutdown + error handling
	http.ListenAndServe(":1323", api.API(dblayer))
}
