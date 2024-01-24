package srv_test

import (
	"net/http"
	"net/http/httptest"
	models "nfon-crud"
	srv "nfon-crud/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockItemStore struct{}

func (m *mockItemStore) CreateItem(name string) {

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
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	mockStore := &mockItemStore{}
	server := srv.NewServer(mockStore)
	server.Router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "healthy", res.Body.String())
}
