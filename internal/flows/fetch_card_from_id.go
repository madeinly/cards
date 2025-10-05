package flows

import (
	"context"
	"errors"

	"github.com/madeinly/cards/internal/features"
)

func GetCardfromId(ctx context.Context, scryfallId string, finish string, language string) (CardBase, error) {

	rawCard, err := features.GetRawCard(ctx, scryfallId)

	if err != nil && errors.Is(err, features.ErrCardNotFound) {
		return CardBase{}, ErrResourceNotFound
	}

	nameEs := features.GetEsName(ctx, rawCard.Uuid)

	setName := features.GetSetName(ctx, rawCard.Setcode)

	cardPrice, err := features.GetCardPrice(ctx, rawCard.Uuid, finish)

	if err != nil && errors.Is(err, features.ErrCardPriceNotFound) {
		return CardBase{}, ErrResourceNotFound
	}

	stock := features.GetCardStock(ctx, rawCard.Uuid, language, finish)

	return CardBase{
		ID:        rawCard.Uuid,
		NameEN:    rawCard.Name,
		NameES:    nameEs,
		ImageURL:  features.BuildImageURL(scryfallId),
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
