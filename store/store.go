package store

import (
	"github.com/brandur/spaced"
)

// Store is an abstract interface used to represent a storage mechanism for
// cards and learning information.
type Store interface {
	// Adds the result of a learning challenge for a card.
	AddLearningResult(card *spaced.Card, success bool) error

	// Gets a card by its unique identifier. Returns nil if a card for that
	// identifier was not found.
	GetCard(id string) (*spaced.Card, error)

	// Puts a card into the store, including a unique identifier for it, its
	// question, and that question's answer.
	PutCard(card *spaced.Card) error
}
