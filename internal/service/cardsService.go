package service

import (
	"context"
	"fmt"

	"githube.com/madeinly/cards/internal/card"
	"githube.com/madeinly/cards/internal/repository"
	"githube.com/madeinly/cards/internal/repository/queries/cardsQuery"
)

func GetCardFromID(cardID string) (card.Card, error) {

	ctx := context.Background()

	cardsDB, err := repository.GetCardsDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	q := cardsQuery.New(cardsDB)

	repoCard, err := q.GetCard(ctx, cardID)

	if err != nil {
		return card.Card{}, err
	}

	return card.Card{
		ID:        repoCard.Uuid,
		Name:      repoCard.Name,
		Sku:       "",
		ImageURL:  card.GetImageURL(cardID),
		SetCode:   repoCard.Setcode,
		SetName:   repoCard.Setname,
		ManaValue: int64(repoCard.Manavalue),
		Rarity:    repoCard.Rarity,
		Colors:    card.TransformColors(repoCard.Colors),
		Types:     card.TransformTypes(repoCard.Types),
	}, nil
}
