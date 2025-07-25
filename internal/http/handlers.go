package http

import (
	"database/sql"
	"encoding/json"
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

	validation.Validate(bag, cardID, card.IdRules)
	validation.Validate(bag, cardFinish, card.FinishRules)

	if bag.HasErrors() {
		_ = bag.WriteHTTP(w)
		return
	}

	card, err := service.GetCardFromID(r.Context(), cardID, cardFinish)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil && err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(card)

}

func CreateCard(w http.ResponseWriter, r *http.Request) {

	bag := validation.New()

	cardId := r.URL.Query().Get("card_id")
	language := r.URL.Query().Get("card_language")
	finish := r.URL.Query().Get("card_finish")
	vendor := r.URL.Query().Get("card_vendor")
	stockSTR := r.URL.Query().Get("card_stock")
	condition := r.URL.Query().Get("card_condition")

	validation.Validate(bag, cardId, card.IdRules)
	validation.Validate(bag, language, card.LanguageRules)
	validation.Validate(bag, finish, card.FinishRules)
	validation.Validate(bag, vendor, card.VendorRules)
	validation.Validate(bag, stockSTR, card.StockRules)
	validation.Validate(bag, condition, card.ConditionRules)

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
		ID:        cardId,
		Vendor:    vendor,
		Language:  language,
		Finish:    finish,
		Stock:     stock,
		Condition: condition,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}
