package memstore

import (
	"testing"

	"github.com/brandur/spaced"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MemstoreSuite struct {
	suite.Suite

	store *Memstore
}

func (s *MemstoreSuite) SetupTest() {
	store, err := NewMemstore()
	assert.Nil(s.T(), err)
	s.store = store
}

func TestMemstoreSuite(t *testing.T) {
	suite.Run(t, new(MemstoreSuite))
}

func (s *MemstoreSuite) TestEmptyGet() {
	actual, err := s.store.GetCard("not-an-identifier")
	assert.Nil(s.T(), actual)
	assert.Nil(s.T(), err)
}

func (s *MemstoreSuite) TestPutAndGet() {
	card := &spaced.Card{ID: "my-id"}

	err := s.store.PutCard(card)
	assert.Nil(s.T(), err)

	actual, err := s.store.GetCard(card.ID)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), card, actual)
}
