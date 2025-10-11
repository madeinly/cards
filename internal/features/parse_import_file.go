package features

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type newCardParams struct {
	ScryfallId string
	Language   string
	Stock      string
	Vendor     string
	Finish     string
	Visibility string
}

func ParseCardsImportFile(ctx context.Context, file *os.File, additive bool) ([]newCardParams, error) {

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv read: %w", err)
	}
	if len(records) == 0 {
		return nil, errors.New("empty csv")
	}

	var newCards []newCardParams
	for i, row := range records {
		if i == 0 { // header
			continue
		}
		if len(row) < 6 {
			return nil, fmt.Errorf("row %d: not enough columns", i+1)
		}
		newCards = append(newCards, newCardParams{
			ScryfallId: row[0],
			Language:   row[1],
			Stock:      row[2],
			Vendor:     row[3],
			Finish:     row[4],
			Visibility: row[5],
		})
	}

	return newCards, nil

}
