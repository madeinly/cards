package flows

import (
	"context"

	mtgDB "github.com/madeinly/cards/internal/drivers/sqlite/sqlc/cards"
	"github.com/madeinly/cards/internal/features"
)

type Set struct {
	SetCode string `json:"card_setCode"`
	SetName string `json:"card_setName"`
}

func GetSets(ctx context.Context) ([]Set, error) {

	cardsDB := features.GetCardsDB()

	qCards := mtgDB.New(cardsDB)

	repoSets, err := qCards.GetSets(ctx)

	if err != nil {
		return nil, err
	}

	var sets []Set

	for _, repoSet := range repoSets {
		sets = append(sets, Set{
			SetCode: repoSet.Code,
			SetName: repoSet.Name,
		})
	}

	return sets, nil

}
