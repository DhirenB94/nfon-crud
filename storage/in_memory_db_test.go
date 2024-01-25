package inMemDB_test

import (
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
