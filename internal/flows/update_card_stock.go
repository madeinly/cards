package flows

import (
	"context"
	"fmt"
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

	cardExist, err := features.CheckCardExist(ctx, nil, features.CheckcardExistParams{
		CardId:   params.Id,
		Finish:   params.Finish,
		Language: params.Language,
	})

	fmt.Println()

	if err != nil {
		return err
	}

	if !cardExist {
		return ErrResourceNotFound
	}

	stock, _ := strconv.ParseInt(params.Stock, 10, 64)

	err = features.UpdateCardStock(ctx, nil, features.UpdateCardStockParams{
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
