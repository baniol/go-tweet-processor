package main

import (
	"github.com/baniol/go-tweet-processor/db"
	"github.com/baniol/go-tweet-processor/twitter"
	"log"
)

var filterString = "docker"

func main() {
	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}
	fn := dblayer.InsertTweet
	twitter.Listen(filterString, fn)
}
