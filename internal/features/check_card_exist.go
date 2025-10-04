package features

import (
	"context"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

type CheckcardExistParams struct {
	CardId   string
	Finish   string
	Language string
}

func CheckcardExist(ctx context.Context, params CheckcardExistParams) (bool, error) {
	db := core.DB()
	queryApp := appDB.New(db)

	exists, err := queryApp.CardExists(ctx, appDB.CardExistsParams{
		ID:       params.CardId,
		Finish:   params.Finish,
		Language: params.Language,
	})

	return exists == 1, err
}
