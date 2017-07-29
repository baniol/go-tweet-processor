package main

import (
	"twitter/mongo"
	"twitter/twitter"
)

var mgs *mongo.MongoDataStore

var filterString = "syria"

func main() {
	mgs = mongo.NewMongoStore()
	fn := mgs.InsertTweet
	twitter.Listen(filterString, fn)
}
