package memstore

import (
	"testing"

	"github.com/brandur/spaced/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MemstoreSuite struct {
	suite.Suite

	store *Memstore
}

func (s *MemstoreSuite) SetupTest() {
	st, err := NewMemstore()
	assert.Nil(s.T(), err)
	s.store = st
}

func TestMemstoreSuite(t *testing.T) {
	suite.Run(t, new(MemstoreSuite))
}

func (s *MemstoreSuite) TestAgainstStoreSuite() {
	for _, test := range store.StoreTests {
		test(s.T(), s.store)
	}
}
