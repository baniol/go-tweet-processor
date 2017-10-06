package web

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data interface{}, code int) {
	// TODO: is interface type ok?; what about validation against data type?
	js, _ := json.Marshal(data)
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	// w.Header().Set("X-XSS-Protection", "1; mode=block")
	// w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	// w.Header().Set("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(js)
}
