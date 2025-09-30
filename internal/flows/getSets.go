package flows

import (
	"context"

	"github.com/madeinly/cards/internal/card"
	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
	"github.com/madeinly/cards/internal/features"
)

func GetSets(ctx context.Context) ([]card.Set, error) {

	cardsDB := features.GetCardsDB()

	qCards := mtgDB.New(cardsDB)

	repoSets, err := qCards.GetSets(ctx)

	if err != nil {
		return nil, err
	}

	var sets []card.Set

	for _, repoSet := range repoSets {
		sets = append(sets, card.Set{
			SetCode: repoSet.Code,
			SetName: repoSet.Name,
		})
	}

	return sets, nil

}
