package web

import (
	"encoding/json"
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

	http.HandleFunc("/", home)
	http.HandleFunc("/count", h.countHandler)
	http.HandleFunc("/authors", h.authorsHandler)
	// http.HandleFunc("/tags", h.tagsHandler)
}

// @TODO only for testing
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func (rh *requestHandler) countHandler(w http.ResponseWriter, r *http.Request) {
	count, err := rh.dbConn.CountTweets()
	if err != nil {
		log.Println("CountTweets error: ", err)
		errorResponse(w)
		return
	}
	c := struct {
		Count int `json:"count"`
	}{count}
	js, _ := json.Marshal(c)
	w.Write(js)
}

func (rh *requestHandler) authorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := rh.dbConn.GetAuthors()
	if err != nil {
		log.Println("GetAuthors error: ", err)
		errorResponse(w)
		return
	}
	js, _ := json.Marshal(authors) // encoder ? move marshaling to db llayer ?
	addHeaders(w)
	// json.NewEncoder(w).Encode(authors)
	w.Write(js) // io.WriteString(w, `{"alive": true}`) ?
}

/*
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
