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

// func Middleware(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Info("middleware", r.URL)
// 		h.ServeHTTP(w, r)
// 	})
// }

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func StartServer(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	r := mux.NewRouter()

	r.HandleFunc("/count", h.countHandler)
	r.HandleFunc("/authors", h.authorsHandler)
	r.HandleFunc("/tags", h.tagsHandler)
	r.HandleFunc("/author/{name}", appHandler(h.authorTweetsHandler))
	http.Handle("/", r)
	log.Info("Handlers initiated")
	// Apache access logging
	// loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, Middleware(r))
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
	http.ListenAndServe(":1323", loggedRouter)
}
