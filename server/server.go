package srv

import (
	"net/http"

	"github.com/bmizerany/pat"
)

type ItemStore interface{}

type Server struct {
	store  ItemStore
	router http.Handler
}

func NewServer(store ItemStore) *Server {
	server := new(Server)
	server.store = store

	router := pat.New()
	server.router = router

	return server

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}
