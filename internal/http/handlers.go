package http

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/service"
	"github.com/madeinly/core/validation"
)

func GetCard(w http.ResponseWriter, r *http.Request) {

	bag := validation.New()

	cardID := r.URL.Query().Get("card_id")

	validation.Validate(bag, cardID, card.IdRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	card, err := service.GetCardFromID(r.Context(), cardID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(card)

}
