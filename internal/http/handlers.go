package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/service"
	"github.com/madeinly/core/validation"
)

func GetCard(w http.ResponseWriter, r *http.Request) {

	bag := validation.New()

	cardID := r.URL.Query().Get("card_id")
	cardFinish := r.URL.Query().Get("card_finish")
	cardLanguage := r.URL.Query().Get("card_language")

	validation.Validate(bag, cardID, card.IdRules)
	validation.Validate(bag, cardLanguage, card.LanguageRules)
	validation.Validate(bag, cardFinish, card.FinishRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	fmt.Println("cardID value", cardID)

	card, err := service.GetCardFromID(r.Context(), cardID, cardFinish, cardLanguage)

	if card.ID == "" {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(card)

}

func PostCreateCard(w http.ResponseWriter, r *http.Request) {

	bag := validation.New()

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cardId := r.PostFormValue("card_id")
	language := r.PostFormValue("card_language")
	finish := r.PostFormValue("card_finish")
	vendor := r.PostFormValue("card_vendor")
	stockSTR := r.PostFormValue("card_stock")
	visibility := r.PostFormValue("card_visibility")

	validation.Validate(bag, cardId, card.IdRules)
	validation.Validate(bag, language, card.LanguageRules)
	validation.Validate(bag, finish, card.FinishRules)
	validation.Validate(bag, vendor, card.VendorRules)
	validation.Validate(bag, stockSTR, card.StockRules)
	validation.Validate(bag, visibility, card.VisibilityRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	stock, err := strconv.ParseInt(stockSTR, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = service.RegisterCard(ctx, service.RegisterCardParams{
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

func BulkCreate(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Println("too big of a file")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	file, header, err := r.FormFile("cards_import")

	if err != nil {
		fmt.Println("cant parse the file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = service.RegisterBulk(ctx, file, header)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
