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

		cardExist, _ := features.CheckcardExist(ctx, features.CheckcardExistParams{
			CardId:   cardId,
			Finish:   item.Finish,
			Language: item.Language,
		})

		if !cardExist {

			err := features.CreateCard(ctx, tx, features.CreateCardParams{})

			if err != nil {
				errs = append(errs, fmt.Errorf("row %s: %w", item.ScryfallId, err))
			}
		}

		if cardExist {

			hasVendor := item.Vendor != ""

			stockQty := features.GetCardStock(ctx, cardId, item.Language, item.Finish)

			importQty, _ := strconv.ParseInt(item.Stock, 10, 64)

			err := features.UpdateCardStock(ctx, features.UpdateCardStockParams{
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
