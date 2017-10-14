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

	// TODO: compare db mongo package with
	// github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/db/db.go
	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}

	// TODO: Graceful shutdown + error handling
	// github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/main.go

	m := api.InitHandlers(dblayer)

	server := http.Server{
		Addr:           ":1323",
		Handler:        m,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infoln("Starting twit processor server")

	server.ListenAndServe()
}
