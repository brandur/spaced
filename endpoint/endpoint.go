package endpoint

import (
	"net/http"

	"github.com/brandur/spaced/errors"
	"github.com/brandur/spaced/store"
	"github.com/gorilla/mux"
)

// Builds a master router that handles all application handlers that make up
// the API.
func BuildRouter(st store.Store) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = &NotFoundHandler{}

	r.Handle("/cards/{id}", &CardHandler{Store: st})

	return r
}

// NotFoundHandler writes a 404 to a response using the standard error
// convention.
type NotFoundHandler struct {
}

// Writes a 404 to a response.
func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiErr := errors.NotFoundEndpoint
	apiErr.WriteHTTP(w)
}
