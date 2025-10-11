package features

import (
	"context"
	"database/sql"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
)

func UpdateCardStockWithTx(ctx context.Context, tx *sql.Tx, params UpdateCardStockParams) error {

	queryApp := appDB.New(tx)

	err := queryApp.UpdateCardStock(ctx, appDB.UpdateCardStockParams{
		ID:        params.Id,
		Language:  params.Language,
		Finish:    params.Finish,
		Stock:     params.Stock,
		HasVendor: params.HasVendor,
	})

	if err != nil {
		return err
	}

	return nil

}
