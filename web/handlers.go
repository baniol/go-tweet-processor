package web

import (
	"encoding/json"
	"go-tweet-processor/db"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

type requestHandler struct {
	dbConn db.DBLayer
}

func InitHandlers(dblayer db.DBLayer) {
	h := new(requestHandler)
	h.dbConn = dblayer // @TODO change name of dblayer

	log.Println("Handlers initiated")

	// http.HandleFunc("/", home)
	http.HandleFunc("/count", h.countHandler)
	http.HandleFunc("/authors", h.authorsHandler)
	http.HandleFunc("/tags", h.tagsHandler)
	http.HandleFunc("/author/", h.authorTweetsHandler)
}

// @TODO only for testing
// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Server", "A Go Web Server")
// 	w.WriteHeader(200)
// }

func (rh *requestHandler) countHandler(w http.ResponseWriter, r *http.Request) {
	count, err := rh.dbConn.CountTweets()
	if err != nil {
		log.Println("CountTweets error: ", err)
		internalErrorResponse(w)
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
		internalErrorResponse(w)
		return
	}
	js, _ := json.Marshal(authors) // encoder ? move marshaling to db llayer ?
	addHeaders(w)
	// json.NewEncoder(w).Encode(authors)
	w.Write(js) // io.WriteString(w, `{"alive": true}`) ?
}

func (rh *requestHandler) tagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := rh.dbConn.GetTags()
	if err != nil {
		log.Println("GetAuthors error: ", err)
		internalErrorResponse(w)
		return
	}
	js, _ := json.Marshal(tags)
	addHeaders(w)
	w.Write(js)
}

func (rh *requestHandler) authorTweetsHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("/author/(.*)")
	match := re.FindAllStringSubmatch(r.URL.String(), -1)
	name, err := url.QueryUnescape(match[0][1])
	if err != nil {
		log.Println("Error parsing query string param: ", err)
		badRequestResponse(w)
		return
	}
	tweets, err := rh.dbConn.GetAuthorTweets(name)
	if err != nil {
		log.Println("GetAuthorTweets error: ", err)
		internalErrorResponse(w)
		return
	}
	js, _ := json.Marshal(tweets)
	if string(js) == "[]" {
		log.Printf("Author not found: %s\n", name)
		badRequestResponse(w)
		return
	}
	addHeaders(w)
	w.Write(js)
}
