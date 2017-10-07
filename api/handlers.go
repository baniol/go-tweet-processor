package web

import (
	"github.com/baniol/go-tweet-processor/db"
	"log"
	"net/http"

	"github.com/baniol/go-tweet-processor/middleware"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func API(dblayer db.DBLayer) http.Handler {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	log.Println("Handlers initiated")

	app.Handle("/count", h.countHandler)
	app.Handle("/authors", h.authorsHandler)
	app.Handle("/tags", h.tagsHandler)
	app.Handle("/author/", h.authorTweetsHandler)

	return app
}
