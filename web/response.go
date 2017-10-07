package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Invalid describes a validation error belonging to a specific field.
type Invalid struct {
	Fld string `json:"field_name"`
	Err string `json:"error"`
}

// InvalidError is a custom error type for invalid fields.
type InvalidError []Invalid

// Error implements the error interface for InvalidError.
func (err InvalidError) Error() string {
	var str string
	for _, v := range err {
		str = fmt.Sprintf("%s,{%s:%s}", str, v.Fld, v.Err)
	}
	return str
}

// JSONError is the response for errors that occur within the API.
type JSONError struct {
	Error  string       `json:"error"`
	Fields InvalidError `json:"fields,omitempty"`
}

var (
	// ErrNotAuthorized occurs when the call is not authorized.
	ErrNotAuthorized = errors.New("Not authorized")

	// ErrDBNotConfigured occurs when the DB is not initialized.
	ErrDBNotConfigured = errors.New("DB not initialized")

	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = errors.New("Entity not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in it's proper form")

	// ErrValidation occurs when there are validation errors.
	ErrValidation = errors.New("Validation errors occurred")
)

// Error handles all error responses for the API.
func Error(cxt context.Context, w http.ResponseWriter, err error) {
	switch errors.Cause(err) {
	case ErrNotFound:
		RespondError(cxt, w, err, http.StatusNotFound)
		return

	case ErrInvalidID:
		RespondError(cxt, w, err, http.StatusBadRequest)
		return

	case ErrValidation:
		RespondError(cxt, w, err, http.StatusBadRequest)
		return

	case ErrNotAuthorized:
		RespondError(cxt, w, err, http.StatusUnauthorized)
		return
	}

	switch e := errors.Cause(err).(type) {
	case InvalidError:
		v := JSONError{
			Error:  "field validation failure",
			Fields: e,
		}

		Respond(w, v, http.StatusBadRequest)
		return
	}

	RespondError(cxt, w, err, http.StatusInternalServerError)
}

// RespondError sends JSON describing the error
func RespondError(ctx context.Context, w http.ResponseWriter, err error, code int) {
	Respond(w, JSONError{Error: err.Error()}, code)
}

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
