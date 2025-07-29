package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/database"
	appDB "github.com/madeinly/cards/internal/sqlc/app"
	mtgDB "github.com/madeinly/cards/internal/sqlc/cards"
	"github.com/madeinly/core"
)

func GetCardFromID(ctx context.Context, cardScryFallID string, finish string, language string) (card.Card, error) {

	cardsDB, err := database.GetCardsDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	qCards := mtgDB.New(cardsDB)

	repoCard, err := qCards.GetCard(ctx, cardScryFallID)

	if err != nil {
		return card.Card{}, fmt.Errorf("could not find the card: %w", err)
	}

	nameES, err := qCards.GetCardNameES(ctx, sql.NullString{Valid: true, String: repoCard.Uuid})

	if err != nil && err != sql.ErrNoRows {
		nameES = sql.NullString{String: ""}
	}

	db := core.DB()

	qApp := appDB.New(db)

	stock, err := qApp.GetCardStockById(ctx, appDB.GetCardStockByIdParams{
		ID:       repoCard.Uuid,
		Language: language,
		Finish:   finish,
	})

	if err != nil && err != sql.ErrNoRows {
		return card.Card{}, fmt.Errorf("there was an error getting the stock")
	}

	if err == sql.ErrNoRows {
		stock = 0
	}

	price, err := qApp.GetPrice(ctx, appDB.GetPriceParams{
		CardID: repoCard.Uuid,
		Finish: finish,
	})

	if err != nil {
		return card.Card{}, fmt.Errorf("there was an error getting the price")
	}

	return card.Card{
		ID:        repoCard.Uuid,
		NameEN:    repoCard.Name,
		NameES:    nameES.String,
		ImageURL:  card.GetImageURL(cardScryFallID),
		SetCode:   repoCard.Setcode,
		SetName:   repoCard.Setname,
		ManaValue: int64(repoCard.Manavalue),
		Number:    repoCard.Number,
		Rarity:    repoCard.Rarity,
		Colors:    repoCard.Colors,
		Types:     repoCard.Types,
		Price:     price,
		Stock:     stock,
	}, nil
}

func RegisterCard(ctx context.Context, params RegisterCardParams) error {

	card, err := GetCardFromID(ctx, params.ID, params.Finish, params.Language)

	if err != nil && err != sql.ErrNoRows {

		return err
	}

	db := core.DB()

	qApp := appDB.New(db)

	exists, err := qApp.CardExists(ctx, appDB.CardExistsParams{
		ID:       params.ID,
		Finish:   params.Finish,
		Language: params.Language,
	})

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("the element already exist")
	}

	hasVendor, err := qApp.GetCardHasVendorById(ctx, params.ID)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("there was a problem getting the hasVendor")
	}

	if err == sql.ErrNoRows {
		hasVendor = false
	}

	hasVendor = hasVendor || params.Vendor != ""

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("there was a problem getting the stock")
	}

	sku := strings.ToLower(params.Language) + "-" + params.Finish + "-" + card.SetCode + "-" + card.Number

	err = qApp.CreateCard(ctx, appDB.CreateCardParams{
		ID:         params.ID,
		NameEn:     card.NameEN,
		NameEs:     card.NameES,
		Sku:        sku,
		UrlImage:   card.ImageURL,
		SetName:    card.SetName,
		SetCode:    card.SetCode,
		ManaValue:  card.ManaValue,
		Colors:     card.Colors,
		Types:      card.Types,
		Finish:     params.Finish,
		HasVendor:  hasVendor,
		Language:   params.Language,
		Visibility: params.Visibility,
		ImagePath:  sql.NullString{Valid: false},
		ImageUrl:   sql.NullString{Valid: false},
		Stock:      params.Stock,
	})

	if err != nil {
		return err
	}

	return nil

}

func RegisterBulk(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	// 1. save the file (your existing code)
	importFolderPath := card.ImportsPath()
	now := time.Now()
	fileName := now.Format("2006-01-02_15-04-05") + ".csv"
	importFilePath := filepath.Join(importFolderPath, fileName)

	dst, err := os.Create(importFilePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := dst.Close(); err != nil {
		return err
	}

	// 2. re-open for reading & parse CSV
	f, err := os.Open(importFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("csv read: %w", err)
	}
	if len(records) == 0 {
		return errors.New("empty csv")
	}

	// header row: skip index 0
	var params []RegisterCardParams
	for i, row := range records {
		if i == 0 { // header
			continue
		}
		if len(row) < 6 {
			return fmt.Errorf("row %d: not enough columns", i+1)
		}
		stock, _ := strconv.ParseInt(row[2], 10, 64)
		params = append(params, RegisterCardParams{
			ID:         row[0],
			Language:   row[1],
			Stock:      stock,
			Vendor:     row[3],
			Finish:     row[4],
			Visibility: row[5],
		})
	}

	db := core.DB()

	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var errs []error
	for _, p := range params {
		if err := RegisterCardTx(ctx, tx, p); err != nil {
			errs = append(errs, fmt.Errorf("row %s: %w", p.ID, err))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("bulk insert finished with %d error(s): %v", len(errs), errs)
	}
	return tx.Commit()
}

func RegisterCardTx(ctx context.Context, tx *sql.Tx, params RegisterCardParams) error {

	card, err := GetCardFromIDTx(ctx, tx, params.ID, params.Finish, params.Language)

	if err != nil && err != sql.ErrNoRows {

		return err
	}

	qApp := appDB.New(tx)

	exists, err := qApp.CardExists(ctx, appDB.CardExistsParams{
		ID:       params.ID,
		Finish:   params.Finish,
		Language: params.Language,
	})

	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("the element already exist")
	}

	hasVendor, err := qApp.GetCardHasVendorById(ctx, params.ID)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("there was a problem getting the hasVendor")
	}

	if err == sql.ErrNoRows {
		hasVendor = false
	}

	hasVendor = hasVendor || params.Vendor != ""

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("there was a problem getting the stock")
	}

	sku := strings.ToLower(params.Language) + "-" + params.Finish + "-" + card.SetCode + "-" + card.Number

	err = qApp.CreateCard(ctx, appDB.CreateCardParams{
		ID:        card.ID,
		NameEn:    card.NameEN,
		NameEs:    card.NameES,
		Sku:       sku,
		UrlImage:  card.ImageURL,
		SetName:   card.SetName,
		SetCode:   card.SetCode,
		ManaValue: card.ManaValue,
		// falta rarity y number
		Colors:     card.Colors,
		Types:      card.Types,
		Finish:     params.Finish,
		HasVendor:  hasVendor,
		Language:   params.Language,
		Visibility: params.Visibility,
		ImagePath:  sql.NullString{Valid: false},
		ImageUrl:   sql.NullString{Valid: false},
		Stock:      params.Stock,
	})

	if err != nil {
		return err
	}

	return nil

}

func GetCardFromIDTx(ctx context.Context, tx *sql.Tx, cardScryFallID string, finish string, language string) (card.Card, error) {

	cardsDB, err := database.GetCardsDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	qCards := mtgDB.New(cardsDB)

	repoCard, err := qCards.GetCard(ctx, cardScryFallID)

	if err != nil {
		return card.Card{}, fmt.Errorf("could not find the card: %w", err)
	}

	nameES, err := qCards.GetCardNameES(ctx, sql.NullString{Valid: true, String: repoCard.Uuid})

	if err != nil && err != sql.ErrNoRows {
		nameES = sql.NullString{String: ""}
	}

	qApp := appDB.New(tx)

	stock, err := qApp.GetCardStockById(ctx, appDB.GetCardStockByIdParams{
		ID:       repoCard.Uuid,
		Language: language,
		Finish:   finish,
	})

	if err != nil && err != sql.ErrNoRows {
		return card.Card{}, fmt.Errorf("there was an error getting the stock")
	}

	if err == sql.ErrNoRows {
		stock = 0
	}

	price, err := qApp.GetPrice(ctx, appDB.GetPriceParams{
		CardID: repoCard.Uuid,
		Finish: finish,
	})

	if err != nil {
		return card.Card{}, nil
	}

	return card.Card{
		ID:        repoCard.Uuid,
		NameEN:    repoCard.Name,
		NameES:    nameES.String,
		ImageURL:  card.GetImageURL(cardScryFallID),
		SetCode:   repoCard.Setcode,
		SetName:   repoCard.Setname,
		ManaValue: int64(repoCard.Manavalue),
		Number:    repoCard.Number,
		Rarity:    repoCard.Rarity,
		Colors:    repoCard.Colors,
		Types:     repoCard.Types,
		Price:     price,
		Stock:     stock,
	}, nil
}

func GetDashboardCards(ctx context.Context, params GetDashboardCardsParams) ([]card.Card, error) {

	db := core.DB()

	qApp := appDB.New(db)

	page, err := strconv.ParseInt(params.Page, 10, 64)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.ParseInt(params.Limit, 10, 64)
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	repoCards, err := qApp.GetCardsWithPrice(ctx, appDB.GetCardsWithPriceParams{
		SetCode: params.SetCode,
		Name:    params.CardName,
		Offset:  offset,
		Limit:   limit,
	})

	fmt.Println("repo cards", repoCards)

	if err != nil && err == sql.ErrNoRows {

		return []card.Card{}, nil
	}

	if err != nil {
		return []card.Card{}, err
	}

	var cards []card.Card

	for _, repoCard := range repoCards {

		cards = append(cards, card.Card{
			ID:        repoCard.ID,
			NameEN:    repoCard.NameEn,
			NameES:    repoCard.NameEs,
			ImageURL:  repoCard.UrlImage,
			SetCode:   repoCard.SetCode,
			SetName:   repoCard.SetName,
			ManaValue: repoCard.ManaValue,
			Colors:    repoCard.Colors,
			Types:     repoCard.Types,
			Price:     repoCard.Price,
			Stock:     repoCard.Stock,
		})

	}

	return cards, nil
}

func GetSets(ctx context.Context) ([]card.Set, error) {

	cardsDB, err := database.GetCardsDB()

	if err != nil {
		fmt.Println(err.Error())
	}

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
