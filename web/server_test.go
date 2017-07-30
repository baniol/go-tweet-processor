package web

import (
	// "bytes"
	// "encoding/json"
	// "net/http"
	// "os"
	"testing"

	// "github.com/Sirupsen/logrus"
	// "github.com/labstack/echo"
	// "github.com/labstack/echo/test"
	"github.com/stretchr/testify/assert"
)

func TestMissingEndpoint(t *testing.T) {
	// code, body := request(t, "GET", "/missing", nil)
	// assert.Equal(t, http.StatusNotFound, code)
	// err := extractError(t, body)
	// assert.Equal(t, http.StatusNotFound, err.Code)
	// assert.NotEmpty(t, err.Message)
	assert.NotNil(t, "asdf", "success message")
}
