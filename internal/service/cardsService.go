package service

import (
	"context"
	"fmt"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/repository"
	"github.com/madeinly/cards/internal/repository/queries/cardsQuery"
	"github.com/madeinly/core"
)

func GetCardFromID(ctx context.Context, cardID string) (card.Card, error) {

	cardsDB, err := repository.GetCardsDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	qCards := cardsQuery.New(cardsDB)

	repoCard, err := qCards.GetCard(ctx, cardID)

	if err != nil {
		return card.Card{}, err
	}

	db := core.DB()

	qCore := cardsQuery.New(db)

	price, err := qCore.GetPrice(ctx, repoCard.Uuid)

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
		Number:    repoCard.Number,
		Rarity:    repoCard.Rarity,
		Colors:    card.TransformColors(repoCard.Colors),
		Types:     card.TransformTypes(repoCard.Types),
		Price:     price,
	}, nil
}
