package http

import (
	"encoding/json"
	"errors"
	"fmt"
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
	cardMatchType := r.URL.Query().Get("card_colorMatchType")
	cardColors := r.URL.Query().Get("card_colors")
	langEn := r.URL.Query().Has("card_langEn")
	langES := r.URL.Query().Has("card_langEs")

	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")

	fmt.Println("langEn value:", langEn)

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
		MatchType:    cardMatchType,
		Colors:       cardColors,
		CardEn:       langEn,
		CardES:       langES,
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
