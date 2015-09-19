package model

// Card (as in "flash card") represents a combination of a question and answer
// which will be used for learning. It also includes a unique identifier so
// that it can be referenced by other entities.
type Card struct {
	// Answer to the card's question.
	Answer string

	// Unique identifier of the card.
	ID string

	// Challenge question posed on the card.
	Question string
}
