package web

import (
	"go-tweet-processor/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type requestHandler struct {
	dbConn db.DBLayer
}

// @TODO separate file for server ?
func StartServer(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	log.Println("Handlers initiated")

	r := mux.NewRouter()

	r.HandleFunc("/count", h.countHandler)
	r.HandleFunc("/authors", h.authorsHandler)
	r.HandleFunc("/tags", h.tagsHandler)
	r.HandleFunc("/author/{name}", h.authorTweetsHandler)
	http.Handle("/", r)
}
