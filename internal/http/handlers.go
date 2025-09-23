package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/service"
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
	stockSTR := r.PostFormValue("card_stock")
	visibility := r.PostFormValue("card_visibility")

	bag.Validate(cardId, card.IdRules)
	bag.Validate(language, card.LanguageRules)
	bag.Validate(finish, card.FinishRules)
	bag.Validate(vendor, card.VendorRules)
	bag.Validate(stockSTR, card.StockRules)
	bag.Validate(visibility, card.VisibilityRules)

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

func GetDashboardCards(w http.ResponseWriter, r *http.Request) {

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

	cards, err := service.GetDashboardCards(ctx, service.GetDashboardCardsParams{
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

func GetSets(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	sets, err := service.GetSets(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)

}
