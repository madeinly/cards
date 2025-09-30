package http

import (
	"fmt"
	"net/http"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/flows"
	"github.com/madeinly/core"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {

	bag := core.Validate()

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cardId := r.PostFormValue("card_id")
	language := r.PostFormValue("card_language")
	finish := r.PostFormValue("card_finish")
	vendor := r.PostFormValue("card_vendor")
	stock := r.PostFormValue("card_stock")
	visibility := r.PostFormValue("card_visibility")

	bag.Validate(cardId, card.IdRules)
	bag.Validate(language, card.LanguageRules)
	bag.Validate(finish, card.FinishRules)
	bag.Validate(vendor, card.VendorRules)
	bag.Validate(stock, card.StockRules)
	bag.Validate(visibility, card.VisibilityRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = flows.RegisterCard(ctx, flows.RegisterCardParams{
		ID:         cardId,
		Vendor:     vendor,
		Language:   language,
		Finish:     finish,
		Stock:      stock,
		Visibility: visibility,
	})

	if err != nil && err.Error() == "the element already exist" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
