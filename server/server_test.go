package srv_test

import (
	"net/http"
	"net/http/httptest"
	srv "nfon-crud/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	server := srv.NewServer()
	server.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "healthy", res.Body.String())
}
