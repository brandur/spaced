package memstore

import (
	"github.com/brandur/spaced/model"
)

// Memstore is a non-persistent Store implemetation that saves all information
// into ephemeral memory.
type Memstore struct {
	cards map[string]*model.Card
}

func NewMemstore() (*Memstore, error) {
	store := &Memstore{
		cards: make(map[string]*model.Card),
	}
	return store, nil
}

func (s *Memstore) AddLearningResult(card *model.Card, success bool) error {
	panic("not implemented")
}

func (s *Memstore) GetCard(id string) (*model.Card, error) {
	return s.cards[id], nil
}

func (s *Memstore) PutCard(card *model.Card) error {
	s.cards[card.ID] = card
	return nil
}
