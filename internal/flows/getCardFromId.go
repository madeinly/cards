package flows

import (
	"context"

	"github.com/madeinly/cards/internal/features/cards"
)

func GetCardfromId(ctx context.Context, scryfallId string, finish string, language string) Card {

	rawCard := cards.GetRawCard(ctx, scryfallId)

	if rawCard.Uuid == "" {
		return Card{}
	}

	nameEs := cards.GetEsName(ctx, rawCard.Uuid)

	setName := cards.GetSetName(ctx, rawCard.Setcode)

	cardPrice := cards.GetCardPrice(ctx, rawCard.Uuid, finish)

	stock := cards.GetCardStock(ctx, rawCard.Uuid, language, finish)

	return Card{
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
	}

}
