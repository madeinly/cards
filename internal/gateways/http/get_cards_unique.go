package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/flows"
	"github.com/madeinly/core"
)

func GetCard(w http.ResponseWriter, r *http.Request) {

	bag := core.Validate()

	cardID := r.URL.Query().Get("card_id")
	cardFinish := r.URL.Query().Get("card_finish")
	cardLanguage := r.URL.Query().Get("card_language")

	bag.Validate(cardID, card.IdRules)
	bag.Validate(cardLanguage, card.LanguageRules)
	bag.Validate(cardFinish, card.FinishRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	card, err := flows.GetCardfromId(r.Context(), cardID, cardFinish, cardLanguage)

	if err != nil && errors.Is(err, flows.ErrResourceNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(card)

}
