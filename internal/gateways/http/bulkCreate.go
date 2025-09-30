package http

import (
	"fmt"
	"net/http"

	"github.com/madeinly/cards/internal/flows"
)

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

	err = flows.RegisterBulk(ctx, file, header)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
