package srv_test

import (
	"net/http"
	"net/http/httptest"
	srv "nfon-crud/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockItemStore struct{}

func TestServer(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	mockStore := &mockItemStore{}
	server := srv.NewServer(mockStore)
	server.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "healthy", res.Body.String())
}
