package spaced

import "net/http"

type Server struct {
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello."))
}
