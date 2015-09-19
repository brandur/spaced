package store

import (
	"testing"

	"github.com/brandur/spaced/model"
	"github.com/stretchr/testify/assert"
)

var (
	// A set of general purposes tests that can be run against any Store
	// implementation to make sure that it behaves in the way that we would
	// expect (purely as a black box).
	//
	// Notably these functions are not prefixed with "Test" so that Go won't
	// try to pick them up.
	StoreTests = []func(t *testing.T, st Store){
		EmptyGet,
		PutAndGet,
	}
)

func EmptyGet(t *testing.T, st Store) {
	actual, err := st.GetCard("not-an-identifier")
	assert.Nil(t, actual)
	assert.Nil(t, err)
}

func PutAndGet(t *testing.T, st Store) {
	card := &model.Card{ID: "my-id"}

	err := st.PutCard(card)
	assert.Nil(t, err)

	actual, err := st.GetCard(card.ID)
	assert.Nil(t, err)

	assert.Equal(t, card, actual)
}
