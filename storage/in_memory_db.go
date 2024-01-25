package inMemDB

import (
	"nfon-crud/models"
	"nfon-crud/utils"
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
	return nil, utils.ErrNoItemByID
}

func (i *InMemDB) UpdateItemByID(id int, name string) error {
	for index, item := range i.ItemStore {
		if id == item.ID {
			item.Name = name
			i.ItemStore[index] = item
			return nil
		}
	}
	return utils.ErrNoItemByID
}

func (i *InMemDB) DeleteItem(id int) error {
	for index, item := range i.ItemStore {
		if id == item.ID {
			delete(i.ItemStore, index)
			return nil
		}
	}
	return utils.ErrNoItemByID
}

func (i *InMemDB) GetAllItems(name string) (*[]models.Item, error) {
	var items []models.Item
	if name == "" {
		for _, item := range i.ItemStore {
			items = append(items, item)
		}
		return &items, nil
	}

	for _, item := range i.ItemStore {
		if name == item.Name {
			items = append(items, item)
		}
	}
	if len(items) == 0 {
		return nil, utils.ErrNoItemByName
	}
	return &items, nil
}
