package http

import (
	"encoding/json"
	"net/http"

	"github.com/madeinly/cards/internal/flows"
)

func GetSets(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	sets, err := flows.GetSets(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)

}
