package srv

import "net/http"

type ItemStore interface{}

type Server struct {
	store ItemStore
}

func NewServer(store ItemStore) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}
