package web

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/baniol/go-tweet-processor/web"
)

func (rh *requestHandler) countHandler(w http.ResponseWriter, r *http.Request) {
	count, err := rh.dbConn.CountTweets()
	if err != nil {
		log.Println("CountTweets error: ", err)
		// internalErrorResponse(w)
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
		// TODO: to middleware
		// internalErrorResponse(w)
		return
	}

	web.Respond(w, authors, 200)
}

func (rh *requestHandler) tagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := rh.dbConn.GetTags()
	if err != nil {
		log.Println("GetAuthors error: ", err)
		// internalErrorResponse(w)
		return
	}
	web.Respond(w, tags, 200)
}

func (rh *requestHandler) authorTweetsHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("/author/(.*)")
	match := re.FindAllStringSubmatch(r.URL.String(), -1)
	name, err := url.QueryUnescape(match[0][1])
	if err != nil {
		log.Println("Error parsing query string param: ", err)
		// badRequestResponse(w)
		return
	}
	tweets, err := rh.dbConn.GetAuthorTweets(name)
	if err != nil {
		log.Println("GetAuthorTweets error: ", err)
		// internalErrorResponse(w)
		return
	}
	// js, _ := json.Marshal(tweets)
	// TODO: different error handling
	// if string(js) == "[]" {
	// 	log.Printf("Author not found: %s\n", name)
	// 	badRequestResponse(w)
	// 	return
	// }
	// addHeaders(w)
	// w.Write(js)
	web.Respond(w, tweets, 200)
}
