package api

import (
	"github.com/baniol/go-tweet-processor/db"
	"github.com/baniol/go-tweet-processor/middleware"
	"github.com/baniol/go-tweet-processor/web"
	"log"
	"net/http"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func InitHandlers(dblayer db.DBLayer) http.Handler {

	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	log.Println("Handlers initiated")

	app.Handle("GET", "/count", h.countHandler)
	app.Handle("GET", "/authors", h.authorsHandler)
	app.Handle("GET", "/tags", h.tagsHandler)
	app.Handle("GET", "/author/:name", h.authorTweetsHandler)

	return app
}
