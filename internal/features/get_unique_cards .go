package features

import (
	"context"
	"database/sql"
	"errors"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
)

func GetUniqueCards(ctx context.Context, cardName string) ([]string, error) {

	db := GetCardsDB()

	query := mtgDB.New(db)

	cards, err := query.ListAllNames(ctx, sql.NullString{Valid: true, String: cardName})

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCardNotFound
	}

	return cards, nil

}
