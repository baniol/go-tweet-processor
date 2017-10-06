package web

import (
	"go-tweet-processor/db"
	"log"
	"net/http"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func InitHandlers(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	log.Println("Handlers initiated")

	http.HandleFunc("/count", h.countHandler)
	http.HandleFunc("/authors", h.authorsHandler)
	http.HandleFunc("/tags", h.tagsHandler)
	http.HandleFunc("/author/", h.authorTweetsHandler)
}
