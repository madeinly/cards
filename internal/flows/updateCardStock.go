package flows

import (
	"context"
	"errors"
	"strconv"

	"github.com/madeinly/cards/internal/features/cards"
)

type UpdateCardStockParams struct {
	Id       string
	Finish   string
	Language string
	Stock    string
}

func UpdateCardStock(ctx context.Context, params UpdateCardStockParams) error {

	stock, _ := strconv.ParseInt(params.Stock, 10, 64)

	err := cards.UpdateCardStock(ctx, cards.UpdateCardStockParams{
		Id:       params.Id,
		Finish:   params.Finish,
		Language: params.Language,
		Stock:    stock,
	})

	if err != nil && errors.Is(err, cards.ErrCardNotFound) {
		return ErrResourceNotFound
	}

	return nil

}
