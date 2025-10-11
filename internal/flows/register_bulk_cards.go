package flows

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/madeinly/cards/internal/features"
	"github.com/madeinly/core"
)

func RegisterBulkCards(ctx context.Context, file multipart.File, addititive bool) error {

	/// ================= Handle the file - Fetch data

	importFilePath, err := features.HandleImportFile(file)
	if err != nil {
		return err
	}

	importFile, err := os.Open(importFilePath)
	if err != nil {
		return err
	}
	defer importFile.Close()

	cardsToImport, err := features.ParseCardsImportFile(ctx, importFile, true)
	if err != nil {
		return err
	}

	/// =================== Process the data - DB Conn

	db := core.DB()

	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var errs []error

	for _, item := range cardsToImport {

		cardId := features.GetCardId(ctx, item.ScryfallId)

		cardExist, err := features.CheckCardExist(ctx, tx, features.CheckcardExistParams{
			CardId:   cardId,
			Finish:   item.Finish,
			Language: item.Language,
		})

		if err != nil {
			return err
		}

		if !cardExist {

			rawCard, err := features.GetRawCard(ctx, item.ScryfallId)

			if err != nil {
				return err

			}

			nameES := features.GetEsName(ctx, tx, rawCard.Uuid)

			visibility, _ := strconv.ParseInt(item.Visibility, 10, 64)

			hasVendor := item.Vendor == ""

			skuParams := features.SkuParams{
				Language: item.Language,
				Finish:   item.Finish,
				SetCode:  rawCard.Setcode,
				Number:   rawCard.Number,
			}

			err = features.CreateCard(ctx, tx, features.CreateCardParams{
				ID:         rawCard.Uuid,
				NameEn:     rawCard.Name,
				NameEs:     nameES,
				Sku:        features.BuildCardSku(skuParams),
				SetName:    rawCard.Setname,
				SetCode:    rawCard.Setcode,
				ManaValue:  int64(rawCard.Manavalue),
				Colors:     rawCard.Colors,
				Types:      rawCard.Types,
				Finish:     item.Finish,
				Rarity:     rawCard.Rarity,
				Number:     rawCard.Number,
				HasVendor:  hasVendor,
				Language:   item.Language,
				Visibility: visibility,
				ImagePath:  "",
				ImageUrl:   features.BuildImageURL(item.ScryfallId),
			})

			if err != nil {
				errs = append(errs, fmt.Errorf("row %s: %w", item.ScryfallId, err))
			}
		}

		if cardExist {

			hasVendor := item.Vendor != ""

			stockQty := features.GetCardStock(ctx, tx, cardId, item.Language, item.Finish)

			importQty, _ := strconv.ParseInt(item.Stock, 10, 64)

			err := features.UpdateCardStockWithTx(ctx, tx, features.UpdateCardStockParams{
				Id:        cardId,
				Finish:    item.Finish,
				Language:  item.Language,
				Stock:     stockQty + importQty,
				HasVendor: hasVendor,
			})

			if err != nil {
				errs = append(errs, fmt.Errorf("row %s: %w", item.ScryfallId, err))
			}

		}

	}

	if len(errs) > 0 {
		return fmt.Errorf("bulk insert finished with %d error(s): %v", len(errs), errs)
	}

	return tx.Commit()
}
