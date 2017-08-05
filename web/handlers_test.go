package web

import (
	"encoding/json"
	// "log"
	"errors"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeMongo struct {
	Error bool
}

func (f *fakeMongo) CountTweets() (int, error) {
	if f.Error == false {
		return 999, nil
	}
	return -1, errors.New("db error")
}

func (f *fakeMongo) GetAuthors() (interface{}, error) {
	if f.Error == true {
		return "", errors.New("DB get authors error")
	}
	var res interface{}
	str := `[{"_id":"Victorine Lenormand","tweets":12},{"_id":"Piki","tweets":4}]`
	JSONString := []byte(str)
	json.Unmarshal(JSONString, &res)
	return res, nil
}

func getFakeInstance(er bool) *requestHandler {
	fakeDB := &fakeMongo{Error: er}
	rh := new(requestHandler)
	rh.dbConn = fakeDB
	return rh
}

func TestCountHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(false)
	handler := http.HandlerFunc(rh.countHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"count\":999}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthorsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/authors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(false)
	handler := http.HandlerFunc(rh.authorsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "[{\"_id\":\"Victorine Lenormand\",\"tweets\":12},{\"_id\":\"Piki\",\"tweets\":4}]"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCountHandlerDBError(t *testing.T) {
	req, err := http.NewRequest("GET", "/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(true)
	handler := http.HandlerFunc(rh.countHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := "{\"error\":\"Internal Server Error\"}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthorsHandlerDBError(t *testing.T) {
	req, err := http.NewRequest("GET", "/authors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(true)
	handler := http.HandlerFunc(rh.authorsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := "{\"error\":\"Internal Server Error\"}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
