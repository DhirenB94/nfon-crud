package srv_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	models "nfon-crud"
	srv "nfon-crud/server"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockItemStore struct {
	items []models.Item
}

func (m *mockItemStore) CreateItem(name string) {
	m.items = append(m.items, models.Item{
		ID:   len(m.items) + 1,
		Name: name},
	)
}
func (m *mockItemStore) GetItemByID(id int) (*models.Item, error) {
	for _, item := range m.items {
		if id == item.ID {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}
func (m *mockItemStore) UpdateItemByID(id int, name string) error {
	return nil
}
func (m *mockItemStore) DeleteItem(id int) error {
	return nil
}
func (m *mockItemStore) GetAllItems() []models.Item {
	return nil
}

func TestServer(t *testing.T) {
	mockStore := &mockItemStore{}
	server := srv.NewServer(mockStore)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	server.Router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "healthy", res.Body.String())
}

func TestCreateItem(t *testing.T) {
	t.Run("create an item successfully", func(t *testing.T) {
		mockItems := []models.Item{
			{ID: 1, Name: "fridge"},
		}
		mockStore := &mockItemStore{
			items: mockItems,
		}
		server := srv.NewServer(mockStore)

		requestBody := strings.NewReader(`{"name": "freezer"}`)
		req, _ := http.NewRequest(http.MethodPost, "/item/create", requestBody)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "item created", res.Body.String())

		expectedOutput := models.Item{ID: 2, Name: "freezer"}
		assert.Equal(t, expectedOutput, mockStore.items[1])
		assert.Len(t, mockStore.items, 2)
	})
}

func TestGetItemByID(t *testing.T) {
	t.Run("get item by valid id", func(t *testing.T) {
		mockItems := []models.Item{
			{ID: 1, Name: "fridge"},
			{ID: 2, Name: "freezer"},
		}
		itemStore := &mockItemStore{
			items: mockItems,
		}
		server := srv.NewServer(itemStore)

		req, _ := http.NewRequest(http.MethodGet, "/item/2", nil)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		var item models.Item
		err := json.NewDecoder(res.Body).Decode(&item)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of items, '%v'", res.Body, err)
		}

		expectedItem := models.Item{
			ID:   2,
			Name: "freezer",
		}
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedItem, item)
	})
	t.Run("return 400 on invalid id format", func(t *testing.T) {
		itemStore := &mockItemStore{}
		server := srv.NewServer(itemStore)

		req, _ := http.NewRequest(http.MethodGet, "/item/abc", nil)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("return a 404 if item does not exist in the store", func(t *testing.T) {
		mockItems := []models.Item{
			{ID: 1, Name: "chair"},
		}
		itemStore := &mockItemStore{
			items: mockItems,
		}
		server := srv.NewServer(itemStore)

		req, _ := http.NewRequest(http.MethodGet, "/item/2", nil)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
