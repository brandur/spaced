package spaced

import (
	"net/http"

	"github.com/brandur/spaced/errors"
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
	switch r.Method {
	case "PUT":
		
	default:
		goto missing
	}

missing:
	apiErr := errors.NotFoundEndpoint
	apiErr.WriteHTTP(w)
}
