package features

import (
	"context"
	"database/sql"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

type UpdateCardStockParams struct {
	Id        string
	Finish    string
	Language  string
	Stock     int64
	HasVendor bool
}

func UpdateCardStock(ctx context.Context, tx *sql.Tx, params UpdateCardStockParams) error {

	var dbConn appDB.DBTX = core.DB()

	if tx != nil {
		dbConn = tx
	}

	queryApp := appDB.New(dbConn)

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
