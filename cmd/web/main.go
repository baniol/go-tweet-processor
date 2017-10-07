package main

import (
	"github.com/baniol/go-tweet-processor/api"
	"github.com/baniol/go-tweet-processor/db"
	"log"
	"net/http"
	"time"
)

func main() {

	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}

	// TODO: Graceful shutdown + error handling

	m := api.InitHandlers(dblayer)

	server := http.Server{
		Addr:           ":1323",
		Handler:        m,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
