package spaced

import (
	"net/http"

	"github.com/brandur/spaced/store"
)

type Server struct {
	store store.Store
}

func NewServer(st store.Store) (*Server, error) {
	server := &Server{
		store: st,
	}
	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello."))
}
