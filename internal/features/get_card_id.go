package features

import (
	"context"
	"fmt"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
)

func GetCardId(ctx context.Context, scryfallId string) string {

	cardsDB := GetCardsDB()

	queryCards := mtgDB.New(cardsDB)

	cardId, err := queryCards.GetCardId(ctx, scryfallId)

	if err != nil {
		fmt.Println(err.Error())
	}

	return cardId.String
}
