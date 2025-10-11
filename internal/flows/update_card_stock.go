package flows

import (
	"context"
	"strconv"

	"github.com/madeinly/cards/internal/features"
)

type UpdateCardStockParams struct {
	Id        string
	Finish    string
	Language  string
	Stock     string
	HasVendor bool
}

func UpdateCardStock(ctx context.Context, params UpdateCardStockParams) error {

	cardExist, err := features.CheckcardExist(ctx, features.CheckcardExistParams{
		CardId:   params.Id,
		Finish:   params.Finish,
		Language: params.Language,
	})

	if err != nil {
		return err
	}

	if !cardExist {
		return ErrResourceNotFound
	}

	stock, _ := strconv.ParseInt(params.Stock, 10, 64)

	err = features.UpdateCardStock(ctx, features.UpdateCardStockParams{
		Id:       params.Id,
		Finish:   params.Finish,
		Language: params.Language,
		Stock:    stock,
	})

	if err != nil {
		return err
	}

	return nil

}
