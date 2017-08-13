package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (rh *requestHandler) countHandler(w http.ResponseWriter, r *http.Request) {
	count, err := rh.dbConn.CountTweets()
	if err != nil {
		log.Errorf("CountTweets error: %s", err)
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
		log.Errorf("GetAuthors error: %s", err)
		internalErrorResponse(w)
		return
	}
	js, _ := json.Marshal(authors)
	addHeaders(w)
	w.Write(js)
}

func (rh *requestHandler) tagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := rh.dbConn.GetTags()
	if err != nil {
		log.Errorf("GetTags error: %s", err)
		internalErrorResponse(w)
		return
	}
	js, _ := json.Marshal(tags)
	addHeaders(w)
	w.Write(js)
}

func (rh *requestHandler) authorTweetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	tweets, err := rh.dbConn.GetAuthorTweets(name)
	if err != nil {
		log.Errorf("GetAuthorTweets error: %s", err)
		internalErrorResponse(w)
		return
	}
	js, _ := json.Marshal(tweets)
	if string(js) == "[]" {
		log.Warnf("Author not found: %s", name)
		badRequestResponse(w)
		return
	}
	addHeaders(w)
	w.Write(js)
}
