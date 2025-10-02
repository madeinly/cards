package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/madeinly/cards/internal/flows"
)

func GetFilteredCards(w http.ResponseWriter, r *http.Request) {

	// bag := core.Validate()

	cardType := r.URL.Query().Get("card_type")
	cardName := r.URL.Query().Get("card_name")
	cardFinish := r.URL.Query().Get("card_finish")
	cardMv := r.URL.Query().Get("card_mv")
	cardPriceMin := r.URL.Query().Get("card_priceMin")
	cardPriceMax := r.URL.Query().Get("card_priceMax")
	cardColors := r.URL.Query().Get("card_colors")
	cardEn := r.URL.Query().Has("card_en")
	cardES := r.URL.Query().Has("card_es")

	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	// bag.Validate(cardID, card.IdRules)
	// bag.Validate(cardLanguage, card.LanguageRules)
	// bag.Validate(cardFinish, card.FinishRules)

	cardsList, err := flows.ListFilteredCards(r.Context(), flows.ListFilteredCardsParams{
		CardName:     cardName,
		CardType:     cardType,
		CardFinish:   cardFinish,
		CardMv:       cardMv,
		CardPriceMin: cardPriceMin,
		CardPriceMax: cardPriceMax,
		Colors:       cardColors,
		CardEn:       cardEn,
		CardES:       cardES,
		Limit:        limit,
		Page:         page,
	})

	if err != nil && errors.Is(err, flows.ErrResourceNotFound) {
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cardsList)

}
