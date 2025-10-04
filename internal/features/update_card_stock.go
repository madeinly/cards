package features

import (
	"context"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

type UpdateCardStockParams struct {
	Id       string
	Finish   string
	Language string
	Stock    int64
}

func UpdateCardStock(ctx context.Context, params UpdateCardStockParams) error {

	db := core.DB()

	queryApp := appDB.New(db)

	err := queryApp.UpdateCardStock(ctx, appDB.UpdateCardStockParams{
		ID:       params.Id,
		Language: params.Language,
		Finish:   params.Finish,
		Stock:    params.Stock,
	})

	if err != nil {
		return err
	}

	return nil

}
