package api

import (
	"net/http"

	"github.com/baniol/go-tweet-processor/db"
	"github.com/baniol/go-tweet-processor/middleware"
	"github.com/baniol/go-tweet-processor/web"
	log "github.com/sirupsen/logrus"
)

type requestHandler struct {
	dbConn db.DBLayer
}

// TODO: app is of type *App - but the function returns http.Handler
func InitHandlers(dblayer db.DBLayer) http.Handler {

	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	log.Infoln("Handlers initiated")

	app.Handle("GET", "/count", h.countHandler)
	app.Handle("GET", "/authors", h.authorsHandler)
	app.Handle("GET", "/tags", h.tagsHandler)
	app.Handle("GET", "/author/:name", h.authorTweetsHandler)

	return app
}
