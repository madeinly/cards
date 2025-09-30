package cards

import (
	"context"
	"database/sql"
	"errors"

	appDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/app"
	"github.com/madeinly/core"
)

// cardId is the uuid from mtgJson
func GetCardPrice(ctx context.Context, cardId string, finish string) float64 {

	db := core.DB()

	queryApp := appDB.New(db)

	cardPrice, err := queryApp.GetPrice(ctx, appDB.GetPriceParams{
		CardID: cardId,
		Finish: "normal",
	})

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return 0
	}

	if err != nil {
		core.Fatal(err, "could not retrieve the price of the card")
	}

	return cardPrice

}
