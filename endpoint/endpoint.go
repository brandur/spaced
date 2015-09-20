package endpoint

import (
	"net/http"

	"github.com/brandur/spaced/errors"
	"github.com/brandur/spaced/store"
	"github.com/gorilla/mux"
)

func BuildRouter(st store.Store) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = &NotFoundHandler{}

	r.Handle("/cards/{id}", &CardHandler{Store: st})

	return r
}

type NotFoundHandler struct {
}

// Writes a 404 to the response. Used as a default route so that 404 messages
// match everywhere.
func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiErr := errors.NotFoundEndpoint
	apiErr.WriteHTTP(w)
}
