package cards

import (
	"context"
	"database/sql"
	"errors"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
	"github.com/madeinly/cards/internal/features"
	"github.com/madeinly/core"
)

// CardId is the uuid from mtgJson
func GetEsName(ctx context.Context, cardId string) string {

	cardsDB := features.GetCardsDB()

	queryCards := mtgDB.New(cardsDB)

	esName, err := queryCards.GetCardNameES(ctx, sql.NullString{String: cardId, Valid: true})

	if err != nil && errors.Is(err, sql.ErrNoRows) {

		return ""

	}

	if err != nil {
		core.Fatal(err, "could not get the ES name of a card")
	}

	return esName.String

}
