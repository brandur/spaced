package errors

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//
// Errors
//

type ErrorsSuite struct {
	suite.Suite
}

func TestErrorsSuite(t *testing.T) {
	suite.Run(t, new(ErrorsSuite))
}

func (s *ErrorsSuite) TestError() {
	e := InternalServer
	str := e.Error()

	assert.NotEqual(s.T(), strings.Index(str, e.ID), -1)
	assert.NotEqual(s.T(), strings.Index(str, e.Message), -1)
	codeStr := fmt.Sprintf("%v", e.StatusCode)
	assert.NotEqual(s.T(), strings.Index(str, codeStr), -1)
}

func (s *ErrorsSuite) TestWriteHTTP() {
	e := InternalServer
	w := httptest.NewRecorder()

	e.WriteHTTP(w)

	// should call down to the handler that we passed above
	assert.Equal(s.T(), e.StatusCode, w.Code)
	assert.Equal(s.T(), "application/json", w.Header().Get("Content-Type"))

	var decoded map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &decoded)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), e.ID, decoded["id"])
	assert.Equal(s.T(), e.Message, decoded["message"])
}
