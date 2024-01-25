package inMemDB

import "nfon-crud/models"

type InMemDB struct{}

func NewInMemDB() *InMemDB {
	return &InMemDB{}
}

func (i *InMemDB) CreateItem(name string) {

}

func (i *InMemDB) GetItemByID(id int) (*models.Item, error) {
	return nil, nil
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
