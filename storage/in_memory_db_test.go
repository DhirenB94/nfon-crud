package inMemDB_test

import (
	"nfon-crud/models"
	inMemDB "nfon-crud/storage"
	"nfon-crud/utils"
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
		assert.ErrorIs(t, err, utils.ErrNoItemByID)
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
		assert.ErrorIs(t, err, utils.ErrNoItemByID)
	})
}

func TestDeleteItem(t *testing.T) {
	t.Run("delete item given valid id", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		err := inMemDb.DeleteItem(2)
		assert.NoError(t, err)

		item, err := inMemDb.GetItemByID(2)
		assert.Nil(t, item)
		assert.ErrorIs(t, err, utils.ErrNoItemByID)
	})
	t.Run("return error when trying to delete an id that is not in the store", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		err := inMemDb.DeleteItem(4)
		assert.ErrorIs(t, err, utils.ErrNoItemByID)
	})
}

func TestGetAllItems(t *testing.T) {
	t.Run("return all items if there are stored items", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")

		expectedItems := &[]models.Item{
			{ID: 1, Name: "fridge"},
			{ID: 2, Name: "freezer"},
		}
		items, err := inMemDb.GetAllItems("")
		assert.NoError(t, err)
		assert.Equal(t, expectedItems, items)
	})
	t.Run("return empty array if no items stored", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		items, err := inMemDb.GetAllItems("")
		assert.NoError(t, err)
		assert.Empty(t, items)
	})
	t.Run("return all items with a valid name", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		inMemDb.CreateItem("fridge")
		inMemDb.CreateItem("freezer")
		inMemDb.CreateItem("freezer")

		expectedItems := &[]models.Item{
			{ID: 2, Name: "freezer"},
			{ID: 3, Name: "freezer"},
		}
		items, err := inMemDb.GetAllItems("freezer")
		assert.NoError(t, err)
		assert.Equal(t, expectedItems, items)
	})
	t.Run("return an error with an invalid name", func(t *testing.T) {
		inMemDb := inMemDB.NewInMemDB()

		items, err := inMemDb.GetAllItems("fridge")
		assert.Nil(t, items)
		assert.ErrorIs(t, err, utils.ErrNoItemByName)
	})
}
