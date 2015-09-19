package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RecoverySuite struct {
	suite.Suite

	requestLog *Recovery
}

func TestRecoverySuite(t *testing.T) {
	suite.Run(t, new(RecoverySuite))
}

func (s *RecoverySuite) SetupTest() {
	s.requestLog = NewRecovery()
}

func (s *RecoverySuite) TestServeHTTP() {
	f := func(w http.ResponseWriter, r *http.Request) {
		panic("major error")
	}
	w := httptest.NewRecorder()
	s.requestLog.ServeHTTP(w, nil, f)

	// should convert the panic to a 500
	assert.Equal(s.T(), 500, w.Code)
}
