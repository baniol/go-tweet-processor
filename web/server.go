package web

import (
	"go-tweet-processor/db"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func StartServer(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	r := mux.NewRouter()

	r.HandleFunc("/count", h.countHandler)
	r.HandleFunc("/authors", h.authorsHandler)
	r.HandleFunc("/tags", h.tagsHandler)
	r.HandleFunc("/author/{name}", h.authorTweetsHandler)
	http.Handle("/", r)
	log.Info("Handlers initiated")
	// Apache access logging
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":1323", loggedRouter)
}
