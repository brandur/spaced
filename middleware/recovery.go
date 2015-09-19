package middleware

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/brandur/spaced/errors"
)

// Recovery provides a middleware that recovers from panics. If one does occur,
// a 500 is written as an HTTP response and logrus' Errorf is invoked, which
// will send the error to Rollbar if a token is configured.
type Recovery struct {
}

// Creates a new Recovery middleware from the specified options.
func NewRecovery() *Recovery {
	return &Recovery{}
}

// Serves an HTTP request, recovering from a panic if necessary.
func (rec *Recovery) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			apiErr := errors.InternalServer
			apiErr.WriteHTTP(w)

			log.WithFields(log.Fields{
				"err": err,
			}).Errorf("PANIC: %v", err)
		}
	}()

	next(w, r)
}
