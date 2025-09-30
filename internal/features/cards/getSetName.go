package cards

import (
	"context"
	"database/sql"
	"errors"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
	"github.com/madeinly/cards/internal/features"
	"github.com/madeinly/core"
)

func GetSetName(ctx context.Context, setCode string) string {

	cardsDB := features.GetCardsDB()

	queryCards := mtgDB.New(cardsDB)

	setName, err := queryCards.GetSetName(ctx, setCode)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return ""
	}

	if err != nil {
		core.Fatal(err, "could not get the set name ")
	}

	return setName

}
