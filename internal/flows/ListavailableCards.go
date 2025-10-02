package flows

import (
	"context"
	"errors"

	"github.com/madeinly/cards/internal/features/cards"
)

type MinCard struct {
	SetCode  string `json:"card_set"`
	Name     string `json:"card_name"`
	CardId   string `json:"card_id"`
	Language string `json:"card_language"`
	Finish   string `json:"card_finish"`
	Stock    int64  `json:"card_stock"`
}

func ListCardsAvailable(ctx context.Context, cardName string) ([]MinCard, error) {

	list, err := cards.ListCardsAvailable(ctx, cardName)

	if err != nil && errors.Is(err, cards.ErrCardNotFound) {
		return nil, ErrResourceNotFound
	}

	if err != nil {
		return nil, ErrServerFailure
	}

	var MinCardList []MinCard

	for _, item := range list {
		MinCardList = append(MinCardList, MinCard{
			SetCode:  item.SetCode,
			Name:     item.NameEn,
			CardId:   item.ID,
			Language: item.Language,
			Finish:   item.Finish,
			Stock:    item.Stock,
		})
	}

	return MinCardList, nil
}
