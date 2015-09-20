package endpoint

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brandur/spaced/model"
	"github.com/brandur/spaced/store"
	"github.com/brandur/spaced/store/memstore"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CardSuite struct {
	suite.Suite

	router *mux.Router
	store store.Store
}

func (s *CardSuite) SetupTest() {
	st, err := memstore.NewMemstore()
	assert.Nil(s.T(), err)
	s.router = BuildRouter(st)
	s.store = st
}

func TestCardSuite(t *testing.T) {
	suite.Run(t, new(CardSuite))
}

func (s *CardSuite) TestGetCard() {
	card := s.buildCard()

	err := s.store.PutCard(card)
	assert.Nil(s.T(), err)

	r := s.buildRequest("GET", "/cards/" + card.ID, nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, r)

	assert.Equal(s.T(), 200, w.Code)
	s.assertResponseBody(w, card)
}

func (s *CardSuite) TestGetNotFound() {
	r := s.buildRequest("GET", "/cards/doesnt-exist", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, r)

	assert.Equal(s.T(), 404, w.Code)
}

func (s *CardSuite) TestPut() {
	card := s.buildCard()
	encoded, err := json.Marshal(card)
	assert.Nil(s.T(), err)
	reader := bytes.NewReader(encoded)

	r := s.buildRequest("PUT", "/cards/" + card.ID, reader)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, r)

	assert.Equal(s.T(), 200, w.Code)
	s.assertResponseBody(w, card)

	actual, err := s.store.GetCard(card.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), card, actual)
}

func (s *CardSuite) TestUnsupportedMethod() {
	r := s.buildRequest("POST", "/cards", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, r)

	assert.Equal(s.T(), 404, w.Code)
}

func (s *CardSuite) assertResponseBody(w *httptest.ResponseRecorder, card *model.Card) {
	var actual *model.Card
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), card, actual)
}

func (s *CardSuite) buildCard() *model.Card {
	return &model.Card{
		ID: "earth-shape",
		Question: "What shape is the Earth?",
		Answer: "Spherical",
	}
}

func (s *CardSuite) buildRequest(method, path string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, "http://example.com" + path, body)
	assert.Nil(s.T(), err)
	return r
}
