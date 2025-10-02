package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/madeinly/cards/internal/flows"
)

func ListCardsAvailable(w http.ResponseWriter, r *http.Request) {

	cardName := r.URL.Query().Get("card_name")

	list, err := flows.ListCardsAvailable(r.Context(), cardName)

	if err != nil && errors.Is(err, flows.ErrResourceNotFound) {
		w.WriteHeader(http.StatusNotFound)

	}

	if err != nil {
		fmt.Println("this is the error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
