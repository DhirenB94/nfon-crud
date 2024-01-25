package inMemDB_test

import (
	"nfon-crud/models"
	inMemDB "nfon-crud/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	store := inMemDB.NewInMemDB()

	store.CreateItem("fridge")
	store.CreateItem("freezer")

	assert.Len(t, store.ItemStore, 2)
}

func TestGetItemByID(t *testing.T) {
	t.Run("return correct item given valid id", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		item, err := inMemDb.GetItemByID(2)
		expectedItem := &models.Item{
			ID:   2,
			Name: "freezer",
		}
		assert.NoError(t, err)
		assert.Equal(t, expectedItem, item)
	})
	t.Run("return error when given id does not exist", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		item, err := inMemDb.GetItemByID(4)
		assert.Nil(t, item)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "item not found")
	})
}
