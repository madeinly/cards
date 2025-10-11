package features

import (
	"context"
	"database/sql"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

type CheckcardExistParams struct {
	CardId   string
	Finish   string
	Language string
}

func CheckCardExist(ctx context.Context, tx *sql.Tx, params CheckcardExistParams) (bool, error) {
	var conn appDB.DBTX = core.DB()
	if tx != nil {
		conn = tx
	}

	q := appDB.New(conn)

	exists, err := q.CardExists(ctx, appDB.CardExistsParams{
		ID:       params.CardId,
		Finish:   params.Finish,
		Language: params.Language,
	})
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
