package srv

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "nfon-crud"
	"strconv"

	"github.com/bmizerany/pat"
)

type ItemStore interface {
	CreateItem(name string)
	GetItemByID(id int) (*models.Item, error)
	UpdateItemByID(id int, name string) error
	DeleteItem(id int) error
	GetAllItems() []models.Item
}

type Server struct {
	Store  ItemStore
	Router http.Handler
}

func NewServer(store ItemStore) *Server {
	server := new(Server)
	server.Store = store

	router := pat.New()
	server.Router = router
	//Routes
	//Homepage acts as healthcheck
	router.Get("/", http.HandlerFunc(server.healthCheck))

	//Create an item
	router.Post("/item/create", http.HandlerFunc(server.createItemHandler))

	//Individual item operations
	router.Get("/item/:id", http.HandlerFunc(server.individualItemHandler))
	router.Patch("/item/:id", http.HandlerFunc(server.individualItemHandler))
	router.Del("/item/:id", http.HandlerFunc(server.individualItemHandler))

	//Get all items
	router.Get("/items", http.HandlerFunc(server.showAllItemsHandler))
	return server

}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}

func (s *Server) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItemName struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&newItemName)
	if err != nil {
		fmt.Println("here")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	s.Store.CreateItem(newItemName.Name)
	w.Write([]byte("item created"))
}

func (s *Server) individualItemHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		item, err := s.Store.GetItemByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(item)

	case http.MethodPatch:
		var updatedItemName struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(r.Body).Decode(&updatedItemName)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		err = s.Store.UpdateItemByID(id, updatedItemName.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Write([]byte("item updated"))
	}

}

func (s *Server) showAllItemsHandler(w http.ResponseWriter, r *http.Request) {

}
