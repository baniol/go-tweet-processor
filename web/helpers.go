package web

import (
	"encoding/json"
	"net/http"
)

func internalErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json") // @TODO addHeaders() instead ?
	w.WriteHeader(http.StatusInternalServerError)
	retError := struct {
		Error string `json:"error"`
	}{"Internal Server Error"}
	js, _ := json.Marshal(retError)
	w.Write(js)
}

func badRequestResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json") // @TODO addHeaders() instead ?
	w.WriteHeader(http.StatusBadRequest)
	retError := struct {
		Error string `json:"error"`
	}{"Bad Request"}
	js, _ := json.Marshal(retError)
	w.Write(js)
}

func addHeaders(w http.ResponseWriter) {
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	// w.Header().Set("X-XSS-Protection", "1; mode=block")
	// w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	// w.Header().Set("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
}
