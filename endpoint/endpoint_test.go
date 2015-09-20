package endpoint

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brandur/spaced/errors"
	"github.com/brandur/spaced/store"
	"github.com/brandur/spaced/store/memstore"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EndpointSuite struct {
	suite.Suite

	router *mux.Router
	store  store.Store
}

func TestEndpointSuite(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}

// Doesn't really touch much of anything useful.
func (s *EndpointSuite) TestBuildRouter() {
	st, err := memstore.NewMemstore()
	assert.Nil(s.T(), err)
	_ = BuildRouter(st)
}

func (s *EndpointSuite) TestNotFoundHandlerServeHTTP() {
	handler := &NotFoundHandler{}

	r, err := http.NewRequest("GET", "http://example.com", nil)
	assert.Nil(s.T(), err)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	var actual *errors.APIError
	err = json.Unmarshal(w.Body.Bytes(), &actual)
	assert.Nil(s.T(), err)
	actual.StatusCode = w.Code
	assert.Equal(s.T(), errors.NotFoundEndpoint, *actual)
}
