package main

import (
	"go-tweet-processor/db"
	"go-tweet-processor/web"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {

	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatalf("Error connecting mongo: %s", err)
	}

	web.StartServer(dblayer)
}
