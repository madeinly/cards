package http

import (
	"errors"
	"net/http"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/flows"
	"github.com/madeinly/core"
)

/*
expects 4 values:

3 that identifies uniquely a card; cardID (mtgJson), language, finish
1 the numeric value that will be updated to the stock of the unique cards
*/

func UpdateCardStock(w http.ResponseWriter, r *http.Request) {

	bag := core.Validate()

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cardId := r.PostFormValue("card_id")
	language := r.PostFormValue("card_language")
	finish := r.PostFormValue("card_finish")
	stock := r.PostFormValue("card_stock")

	bag.Validate(cardId, card.IdRules)
	bag.Validate(language, card.LanguageRules)
	bag.Validate(finish, card.FinishRules)
	bag.Validate(stock, card.StockRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	err = flows.UpdateCardStock(r.Context(), flows.UpdateCardStockParams{
		Id:       cardId,
		Finish:   finish,
		Language: language,
		Stock:    stock,
	})

	if err != nil && errors.Is(err, flows.ErrResourceNotFound) {
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}
