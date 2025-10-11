package features

import (
	"context"
	"database/sql"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
)

func GetCardId(ctx context.Context, scryfallId string) string {

	cardsDB := GetCardsDB()

	queryCards := mtgDB.New(cardsDB)

	cardId, _ := queryCards.GetCardId(ctx, sql.NullString{Valid: true, String: scryfallId})

	return cardId
}
