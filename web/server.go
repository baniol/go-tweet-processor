package web

import (
	"go-tweet-processor/db"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// @TODO separate file for server ?
func StartServer(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	log.Info("Handlers initiated - info")

	r := mux.NewRouter()

	r.HandleFunc("/count", h.countHandler)
	r.HandleFunc("/authors", h.authorsHandler)
	r.HandleFunc("/tags", h.tagsHandler)
	r.HandleFunc("/author/{name}", h.authorTweetsHandler)
	http.Handle("/", r)
}
