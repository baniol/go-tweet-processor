package web

import (
	"github.com/labstack/echo"
	"go-tweet-processor/mongo"
	"os"
)

var mgs *mongo.MongoDataStore

func StartServer() {

	mgs = mongo.NewMongoStore()

	e := echo.New()
	e.GET("/count", count)
	e.GET("/authors", authors)
	e.GET("/tags", tags)
	e.GET("/author/:name", authorTweets)
	e.Logger.Fatal(e.Start(":" + os.Getenv("WEB_PORT")))
}
