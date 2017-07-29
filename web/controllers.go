package web

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

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
