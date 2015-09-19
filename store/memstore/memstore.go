package memstore

import (
	"github.com/brandur/spaced"
)

// Memstore is a non-persistent Store implemetation that saves all information
// into ephemeral memory.
type Memstore struct {
	cards map[string]*spaced.Card
}

func NewMemstore() (*Memstore, error) {
	store := &Memstore{
		cards: make(map[string]*spaced.Card),
	}
	return store, nil
}

func (s *Memstore) AddLearningResult(card *spaced.Card, success bool) error {
	panic("not implemented")
}

func (s *Memstore) GetCard(id string) (*spaced.Card, error) {
	return s.cards[id], nil
}

func (s *Memstore) PutCard(card *spaced.Card) error {
	s.cards[card.ID] = card
	return nil
}
