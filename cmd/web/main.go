package main

import (
	"go-tweet-processor/db"
	"go-tweet-processor/web"
	"log"
	"net/http"
)

// type fakeMongo struct{}

// func (f *fakeMongo) CountTweets() (int, error) {
// 	return 999, nil
// }

func main() {
	dblayer, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("db error")
	}

	// fake := new(fakeMongo)
	web.InitHandlers(dblayer)
	http.ListenAndServe(":1323", nil)
}
