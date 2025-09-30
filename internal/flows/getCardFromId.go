package flows

import (
	"context"
	"errors"

	"github.com/madeinly/cards/internal/features/cards"
)

func GetCardfromId(ctx context.Context, scryfallId string, finish string, language string) (CardBase, error) {

	rawCard, err := cards.GetRawCard(ctx, scryfallId)

	if err != nil && errors.Is(err, cards.ErrCardNotFound) {
		return CardBase{}, ErrResourceNotFound
	}

	nameEs := cards.GetEsName(ctx, rawCard.Uuid)

	setName := cards.GetSetName(ctx, rawCard.Setcode)

	cardPrice, err := cards.GetCardPrice(ctx, rawCard.Uuid, finish)

	if err != nil && errors.Is(err, cards.ErrCardPriceNotFound) {
		return CardBase{}, ErrResourceNotFound
	}

	stock := cards.GetCardStock(ctx, rawCard.Uuid, language, finish)

	return CardBase{
		ID:        rawCard.Uuid,
		NameEN:    rawCard.Name,
		NameES:    nameEs,
		ImageURL:  cards.GetImageURL(scryfallId),
		SetCode:   rawCard.Setcode,
		SetName:   setName,
		ManaValue: int64(rawCard.Manavalue),
		Number:    rawCard.Number,
		Rarity:    rawCard.Rarity,
		Colors:    rawCard.Colors,
		Types:     rawCard.Types,
		Price:     cardPrice,
		Stock:     stock,
	}, nil

}
