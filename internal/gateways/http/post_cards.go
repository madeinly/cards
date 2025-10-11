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

	ctx := r.Context()

	err = flows.RegisterCard(ctx, flows.RegisterCardParams{
		ScryfallId: cardId,
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

func BulkCreate(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Println("too big of a file")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registerAdition := r.URL.Query().Has("card_additive")

	ctx := r.Context()

	file, header, err := r.FormFile("card_import")

	if err != nil {
		fmt.Println("cant parse the file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = flows.RegisterBulk(ctx, file, header, registerAdition)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
