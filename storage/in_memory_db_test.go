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

func TestUpdateItemById(t *testing.T) {
	t.Run("correctly update the item given a valid id", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		err := inMemDb.UpdateItemByID(1, "mini-fridge")
		assert.NoError(t, err)

		expectedUpdatedItem := &models.Item{
			ID:   1,
			Name: "mini-fridge",
		}

		actualUpdatedItem, err := inMemDb.GetItemByID(1)
		assert.NoError(t, err)
		assert.Equal(t, expectedUpdatedItem, actualUpdatedItem)
	})
	t.Run("return error when trying to update an id that is not in the store", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		err := inMemDb.UpdateItemByID(4, "chair")
		assert.Error(t, err)
		assert.ErrorContains(t, err, "item not found")
	})
}
