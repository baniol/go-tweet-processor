package main

import (
	"go-tweet-processor/db"
	"go-tweet-processor/twitter"
	"log"
)

// var mgs *mongo.MongoDataStore

var filterString = "syria"

func main() {
	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}

	// web.InitHandlers(dblayer)

	// mgs = mongo.NewMongoStore()
	fn := dblayer.InsertTweet
	twitter.Listen(filterString, fn)
}
