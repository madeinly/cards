package features

import (
	"context"
	"database/sql"
	"errors"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

func ListCardsAvailable(ctx context.Context, cardName string) ([]appDB.ListAvailableCardsRow, error) {

	db := core.DB()

	queryApp := appDB.New(db)

	cards, err := queryApp.ListAvailableCards(ctx, sql.NullString{Valid: true, String: cardName})

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []appDB.ListAvailableCardsRow{}, ErrCardNotFound
	}

	return cards, nil

}
