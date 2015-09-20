package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/brandur/spaced/errors"
	"github.com/brandur/spaced/model"
	"github.com/brandur/spaced/store"
	"github.com/gorilla/mux"
)

type CardHandler struct {
	Store store.Store
}

func (s *CardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var card *model.Card
	var err error

	id := mux.Vars(r)["id"]
	log.Infof("vars = %v", mux.Vars(r))
	log.WithFields(log.Fields{
		"handler": "card",
		"id": id,
	}).Infof("Looking up card %v", id)

	switch r.Method {
	case "GET":
		card, err = s.Store.GetCard(id)
		if err != nil {
			panic(err)
		}
		if card == nil {
			apiErr := errors.NotFoundRecord
			apiErr.WriteHTTP(w)
			return
		}

	case "PUT":
		rBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		err = json.Unmarshal(rBytes, &card)
		if err != nil {
			panic(err)
		}

		// make sure that out identifier field is consistent to what it's keyed
		// to
		card.ID = id

		err = s.Store.PutCard(card)
		if err != nil {
			panic(err)
		}

	default:
		apiErr := errors.NotFoundEndpoint
		apiErr.WriteHTTP(w)
		return
	}

	encoded, err := json.Marshal(card)
	if err != nil {
		panic(err)
	}
	w.Write(encoded)
}
