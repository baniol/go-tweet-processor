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

// InitHandlers registers routes and middleware.
// Returns http.Handler interface
// Explanation: app is of type *App, but the function returns http.Handler
// This is possible because App embeds *httptreemux.TreeMux which implements ServeHTTP function:
// github.com/dimfeld/httptreemux/router.go
func InitHandlers(dblayer db.DBLayer) http.Handler {

	// TODO: to factory function ?
	h := new(requestHandler)
	h.dbConn = dblayer

	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	log.Infoln("Handlers initiated")

	app.Handle("GET", "/count", h.countHandler)
	app.Handle("GET", "/authors", h.authorsHandler)
	app.Handle("GET", "/tags", h.tagsHandler)
	app.Handle("GET", "/author/:name", h.authorTweetsHandler)

	return app
}
