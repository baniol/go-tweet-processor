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
	Error      bool
	BadRequest bool
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

func (f *fakeMongo) GetTags() (interface{}, error) {
	if f.Error == true {
		return "", errors.New("DB get tags error")
	}
	var res interface{}
	str := `
	[
		{
		"_id": "syria",
		"count": 115
		},
		{
		"_id": "isis",
		"count": 13
		},
		{
		"_id": "raqqa",
		"count": 11
		}
	]`
	JSONString := []byte(str)
	json.Unmarshal(JSONString, &res)
	return res, nil
}

func (f *fakeMongo) GetAuthorTweets(string) (interface{}, error) {
	if f.Error == true {
		return "", errors.New("DB get tags error")
	}
	var res interface{}
	var str string
	if f.BadRequest == true {
		str = "[]"
	} else {
		str = `[{"Name": "some example text"}]`
	}
	JSONString := []byte(str)
	json.Unmarshal(JSONString, &res)
	return res, nil
}

func getFakeInstance(er bool, badReq bool) *requestHandler {
	fakeDB := &fakeMongo{Error: er, BadRequest: badReq}
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
	rh := getFakeInstance(false, false)
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
	rh := getFakeInstance(false, false)
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
	rh := getFakeInstance(true, false)
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
	rh := getFakeInstance(true, false)
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

func TestTagsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/tags", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(false, false)
	handler := http.HandlerFunc(rh.tagsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "[{\"_id\":\"syria\",\"count\":115},{\"_id\":\"isis\",\"count\":13},{\"_id\":\"raqqa\",\"count\":11}]"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTagsHandlerDBError(t *testing.T) {
	req, err := http.NewRequest("GET", "/tags", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(true, false)
	handler := http.HandlerFunc(rh.tagsHandler)
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

func TestAuthorTweetsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/author/some", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(false, false)
	handler := http.HandlerFunc(rh.authorTweetsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"Name":"some example text"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAuthorTweetsHandlerDBError(t *testing.T) {
	req, err := http.NewRequest("GET", "/author/example", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(true, false)
	handler := http.HandlerFunc(rh.authorTweetsHandler)
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

func TestAuthorTweetsHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/author/example", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rh := getFakeInstance(false, true)
	handler := http.HandlerFunc(rh.authorTweetsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"error":"Bad Request"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
