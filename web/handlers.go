package web

import (
	"fmt"
	// "gopkg.in/mgo.v2/dbtest"
	// "encoding/json"
	"go-tweet-processor/db"
	"log"
	"net/http"
	"strconv"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func InitHandlers(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	log.Println("Handlers initiated")

	http.HandleFunc("/count", h.countHandler)
	// http.HandleFunc("/authors", h.authorsHandler)
	// http.HandleFunc("/tags", h.tagsHandler)
}

func (rh *requestHandler) countHandler(w http.ResponseWriter, r *http.Request) {
	count, _ := rh.dbConn.CountTweets()
	fmt.Fprint(w, strconv.Itoa(count))
}

/*
func (rh *requestHandler) authorsHandler(w http.ResponseWriter, r *http.Request) {
	authors := rh.dbConn.GetAuthors()
	js, err := json.Marshal(authors) // encoder ?
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(js) // io.WriteString(w, `{"alive": true}`) ?
}

func (rh *requestHandler) tagsHandler(w http.ResponseWriter, r *http.Request) {
	tags := rh.dbConn.GetTags()
	js, err := json.Marshal(tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(js)
}
*/
// func authorTweets(c echo.Context) error {
// 	names := mgs.GetAuthorTweets(c.Param("name"))
// 	return c.JSON(http.StatusOK, names)
// }
