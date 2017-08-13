package main

import (
	"go-tweet-processor/db"
	"go-tweet-processor/twitter"
	"os"

	log "github.com/sirupsen/logrus"
)

var filterString = "syria"

func init() {
	// @TODO - separate module wrapper ? atm - code duplication in cmd/web
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatalf("Error connecting mongo: %s", err)
	}
	fn := dblayer.InsertTweet
	twitter.Listen(filterString, fn)
}
