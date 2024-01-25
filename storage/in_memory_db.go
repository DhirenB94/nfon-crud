package inMemDB

import (
	"errors"
	"nfon-crud/models"
)

type InMemDB struct {
	ItemStore map[int]models.Item
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		ItemStore: make(map[int]models.Item),
	}
}

func (i *InMemDB) CreateItem(name string) {
	id := len(i.ItemStore) + 1
	i.ItemStore[id] = models.Item{
		ID:   id,
		Name: name,
	}
}

func (i *InMemDB) GetItemByID(id int) (*models.Item, error) {
	for _, item := range i.ItemStore {
		if id == item.ID {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (i *InMemDB) UpdateItemByID(id int, name string) error {
	return nil
}

func (i *InMemDB) DeleteItem(id int) error {
	return nil
}

func (i *InMemDB) GetAllItems(name string) (*[]models.Item, error) {
	return nil, nil
}
