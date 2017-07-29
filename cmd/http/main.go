package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
	"twitter/mongo"
)

var mgs *mongo.MongoDataStore

func main() {

	mgs = mongo.NewMongoStore()

	e := echo.New()
	e.GET("/count", count)
	e.GET("/authors", authors)
	e.GET("/tags", tags)
	e.GET("/author/:name", authorTweets)
	e.Logger.Fatal(e.Start(":" + os.Getenv("WEB_PORT")))
}

func authorTweets(c echo.Context) error {
	names := mgs.GetAuthorTweets(c.Param("name"))
	return c.JSON(http.StatusOK, names)
}

func count(c echo.Context) error {
	count, _ := mgs.CountTweets()
	return c.String(http.StatusOK, "count: "+strconv.Itoa(count))
}

func authors(c echo.Context) error {
	results := mgs.GetAuthors()
	return c.JSON(http.StatusOK, results)
}

func tags(c echo.Context) error {
	results := mgs.CountTags()
	return c.JSON(http.StatusOK, results)
}
