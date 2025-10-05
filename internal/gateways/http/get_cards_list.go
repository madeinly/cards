package http

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/flows"
	"github.com/madeinly/core"
)

func ListCards(w http.ResponseWriter, r *http.Request) {

	setCode := r.URL.Query().Get("card_setCode")
	cardName := r.URL.Query().Get("card_name")
	cardPage := r.URL.Query().Get("card_page")
	cardLimit := r.URL.Query().Get("card_limit")

	bag := core.Validate()

	bag.Validate(setCode, card.SetCodeRules)
	bag.Validate(cardName, card.CardNameRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	ctx := r.Context()

	cards, err := flows.GetDashboardCards(ctx, flows.GetDashboardCardsParams{
		SetCode:  setCode,
		CardName: cardName,
		Page:     cardPage,
		Limit:    cardLimit,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cards)

}
