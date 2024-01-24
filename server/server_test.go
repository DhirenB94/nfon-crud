package srv_test

import (
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
	return nil, nil
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
