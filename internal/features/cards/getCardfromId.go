package cards

import (
	"context"
	"database/sql"
	"errors"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
	"github.com/madeinly/cards/internal/features"
)

/*
uses scryfallID to retrieve mtgDB.GetCardRow:

	type GetCardRow struct {
		Uuid      string  `json:"uuid"`
		Name      string  `json:"name"`
		Setcode   string  `json:"setcode"`
		Manavalue float64 `json:"manavalue"`
		Rarity    string  `json:"rarity"`
		Colors    string  `json:"colors"`
		Types     string  `json:"types"`
		Number    string  `json:"number"`
		Setname   string  `json:"setname"`
	}
*/
func GetRawCard(ctx context.Context, scryfallId string) mtgDB.GetCardRow {

	cardsDB := features.GetCardsDB()

	queryCards := mtgDB.New(cardsDB)

	repoCard, err := queryCards.GetCard(ctx, scryfallId)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return mtgDB.GetCardRow{}
	}

	return repoCard

}
