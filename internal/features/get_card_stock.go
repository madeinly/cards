package features

import (
	"context"
	"database/sql"
	"errors"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

// cardId is the uuid from mtgJson
func GetCardStock(ctx context.Context, tx *sql.Tx, cardId string, language string, finish string) int64 {
	var conn appDB.DBTX = core.DB()

	if tx != nil {
		conn = tx
	}

	queryApp := appDB.New(conn)

	stock, err := queryApp.GetCardStockById(ctx, appDB.GetCardStockByIdParams{ID: cardId, Language: language, Finish: finish})

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return 0
	}

	if err != nil {
		core.Fatal(err, err.Error())
	}

	return stock

}
