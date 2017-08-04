package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeMongo struct{}

func (f *fakeMongo) CountTweets() (int, error) {
	return 999, nil
}

// fake := new(fakeMongo)
// log.Println(fake)

func TestCountTweets(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	fakeDB := new(fakeMongo)
	rh := new(requestHandler)
	rh.dbConn = fakeDB

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rh.countHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"alive": true}`
	expected := "999"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
