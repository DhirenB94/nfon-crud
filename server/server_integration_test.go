package srv_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	models "nfon-crud/models"
	srv "nfon-crud/server"
	inMemDB "nfon-crud/storage"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerIntegration(t *testing.T) {
	store := inMemDB.NewInMemDB()
	server := srv.NewServer(store)

	itemOne := strings.NewReader(`{"name": "fridge"}`)
	itemTwo := strings.NewReader(`{"name": "freezer"}`)

	req1, _ := http.NewRequest(http.MethodPost, "/item/create", itemOne)
	req2, _ := http.NewRequest(http.MethodPost, "/item/create", itemTwo)
	server.Router.ServeHTTP(httptest.NewRecorder(), req1)
	server.Router.ServeHTTP(httptest.NewRecorder(), req2)

	reqAll, _ := http.NewRequest(http.MethodGet, "/items", nil)
	response := httptest.NewRecorder()
	server.Router.ServeHTTP(response, reqAll)
	assert.Equal(t, http.StatusOK, response.Code)

	var items []models.Item
	err := json.NewDecoder(response.Body).Decode(&items)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of items, '%v'", response.Body, err)
	}

	expectedResponse := []models.Item{
		{ID: 1, Name: "fridge"},
		{ID: 2, Name: "freezer"},
	}
	assert.ElementsMatch(t, expectedResponse, items)
	assert.Len(t, items, 2)
}
