package main

import (
	"go-tweet-processor/db"
	"go-tweet-processor/twitter"
	"log"
)

var filterString = "syria"

func main() {
	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}
	fn := dblayer.InsertTweet
	twitter.Listen(filterString, fn)
}
