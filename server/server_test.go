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
	for index, item := range m.items {
		if id == item.ID {
			m.items[index].Name = name
			return nil
		}
	}
	return errors.New("item not found")
}

func (m *mockItemStore) DeleteItem(id int) error {
	for index, item := range m.items {
		if id == item.ID {
			m.items[index] = models.Item{}
			return nil
		}
	}
	return errors.New("item not found")
}

func (m *mockItemStore) GetAllItems(name string) (*[]models.Item, error) {
	return &m.items, nil
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

func TestUpdate(t *testing.T) {
	t.Run("update item with a valid id", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "mini fridge"}`)

		items := []models.Item{
			{ID: 5, Name: "fridge"},
			{ID: 6, Name: "freezer"},
		}
		mockStore := &mockItemStore{
			items: items,
		}
		server := srv.NewServer(mockStore)

		req, _ := http.NewRequest(http.MethodPatch, "/item/5", requestBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "item updated", res.Body.String())

		expectedUpdatedItem := models.Item{
			ID:   5,
			Name: "mini fridge",
		}
		assert.Equal(t, expectedUpdatedItem, mockStore.items[0])
		assert.Len(t, mockStore.items, 2)
	})
	t.Run("return a 404 if item does not exist", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "mini fridge"}`)
		mockItems := []models.Item{
			{ID: 1, Name: "fridge"},
		}
		mockStore := &mockItemStore{
			items: mockItems,
		}
		server := srv.NewServer(mockStore)

		req, _ := http.NewRequest(http.MethodPatch, "/item/3", requestBody)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete item with valid id", func(t *testing.T) {
		items := []models.Item{
			{ID: 5, Name: "fridge"},
			{ID: 7, Name: "freezer"},
			{ID: 10, Name: "mini fridge"},
		}
		mockStore := &mockItemStore{
			items: items,
		}
		server := srv.NewServer(mockStore)

		req, _ := http.NewRequest(http.MethodDelete, "/item/5", nil)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "item deleted", res.Body.String())

		assert.Equal(t, models.Item{}, mockStore.items[0])

	})
	t.Run("return 404 if item does not exist in the store", func(t *testing.T) {
		mockItems := []models.Item{
			{ID: 1, Name: "fridge"},
		}
		mockStore := &mockItemStore{
			items: mockItems,
		}
		server := srv.NewServer(mockStore)

		req, _ := http.NewRequest(http.MethodDelete, "/item/2", nil)
		res := httptest.NewRecorder()
		server.Router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}

func TestGetAllItems(t *testing.T) {
	t.Run("when no name is provided in the query", func(t *testing.T) {
		t.Run("returns all items as JSON", func(t *testing.T) {
			allItems := []models.Item{
				{ID: 5, Name: "fridge"},
				{ID: 7, Name: "freezer"},
			}
			mockStore := &mockItemStore{
				items: allItems,
			}
			server := srv.NewServer(mockStore)

			req, _ := http.NewRequest(http.MethodGet, "/items", nil)
			res := httptest.NewRecorder()
			server.Router.ServeHTTP(res, req)

			var items []models.Item
			err := json.NewDecoder(res.Body).Decode(&items)
			if err != nil {
				t.Fatalf("Unable to parse response from server %q into slice of items, '%v'", res.Body, err)
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, allItems, items)
			assert.Equal(t, srv.JsonContentType, res.Header().Get("content-type"))
		})
		t.Run("returns message to the user if no items are in the store yet", func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/items", nil)
			res := httptest.NewRecorder()

			itemStore := &mockItemStore{}
			server := srv.NewServer(itemStore)

			server.Router.ServeHTTP(res, req)

			assert.Equal(t, "no items to display yet", res.Body.String())
			assert.Equal(t, http.StatusOK, res.Code)
		})
	})

	t.Run("when a name parameter is provided", func(t *testing.T) {
		t.Run("given an item exists in the store for the name", func(t *testing.T) {

		})
		t.Run("return 404 when no item exists in the store for that name", func(t *testing.T) {

		})
	})
}
