package web

import (
	// "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go-tweet-processor/mongo"
	"os"
)

var mgs *mongo.MongoDataStore

func StartServer() {

	mgs = mongo.NewMongoStore()
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Info("starting server")
	e.GET("/count", count)
	e.GET("/authors", authors)
	e.GET("/tags", tags)
	e.GET("/author/:name", authorTweets)
	e.Logger.Fatal(e.Start(":" + os.Getenv("WEB_PORT")))
}
