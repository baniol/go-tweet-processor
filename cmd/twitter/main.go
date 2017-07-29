package main

import (
	"go-tweet-processor/mongo"
	"go-tweet-processor/twitter"
)

var mgs *mongo.MongoDataStore

var filterString = "syria"

func main() {
	mgs = mongo.NewMongoStore()
	fn := mgs.InsertTweet
	twitter.Listen(filterString, fn)
}
