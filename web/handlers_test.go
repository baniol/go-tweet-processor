package web

import (
	"encoding/json"
	// "log"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeMongo struct{}

func (f *fakeMongo) CountTweets() (int, error) {
	return 999, nil
}

func (f *fakeMongo) GetAuthors() (interface{}, error) {

	var res interface{}
	str := `[{"_id":"Victorine Lenormand","tweets":12},{"_id":"Piki","tweets":4}]`
	JSONString := []byte(str)
	// json.NewDecoder(JSONString).Decode(&res)
	json.Unmarshal(JSONString, &res)
	return res, nil
}

func getFakeInstance() *requestHandler {
	fakeDB := new(fakeMongo)
	rh := new(requestHandler)
	rh.dbConn = fakeDB
	return rh
}

func TestCountHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	rh := getFakeInstance()
	handler := http.HandlerFunc(rh.countHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "999"
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

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	rh := getFakeInstance()
	handler := http.HandlerFunc(rh.authorsHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "[{\"_id\":\"Victorine Lenormand\",\"tweets\":12},{\"_id\":\"Piki\",\"tweets\":4}]"
	fmt.Printf("%T\n", rr.Body)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
