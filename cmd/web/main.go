package main

import (
	"net/http"
	"os"
	"time"

	"github.com/baniol/go-tweet-processor/api"
	"github.com/baniol/go-tweet-processor/db"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {

	log.Infoln("Starting tweet processor server")

	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("DB connection failed")
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = ":1323"
	}

	m := api.InitHandlers(dblayer)

	server := http.Server{
		Addr:           host,
		Handler:        m,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

}
