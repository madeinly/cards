package flows

import (
	"database/sql"
	"time"

	"fmt"
	"os"
	"path"

	_ "modernc.org/sqlite"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/cards/internal/features"
)

func UpdateCardsDB() error {
	cardsPath, err := card.CardsPath()
	if err != nil {
		return fmt.Errorf("failed to get cards path: %w", err)
	}

	mtgjsonURL := "https://mtgjson.com/api/v5/AllPrintings.sqlite.gz"
	gzPath := path.Join(cardsPath, "AllPrintings.sqlite.gz.tmp")
	defer os.Remove(gzPath)

	if err := card.DownLoadFile(mtgjsonURL, gzPath); err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	tempDBPath := path.Join(cardsPath, "mtgDB.sqlite.tmp")
	if err := card.Ungz(gzPath, tempDBPath); err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	if err := verifySQLiteDB(tempDBPath); err != nil {
		os.Remove(tempDBPath)
		return fmt.Errorf("database verification failed: %w", err)
	}

	if err := os.Chmod(tempDBPath, 0444); err != nil {
		os.Remove(tempDBPath)
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	db, _ := features.GetCardsDB()

	var dbWasInUse bool

	if db != nil {
		dbWasInUse = true
	}

	if dbWasInUse {
		db.Close()
		time.Sleep(time.Millisecond * 100)
	}

	finalPath := path.Join(cardsPath, "mtgDB.sqlite")
	if err := os.Rename(tempDBPath, finalPath); err != nil {
		return fmt.Errorf("failed to replace database: %w", err)
	}

	if dbWasInUse {
		features.InitCardsDB()
	}

	return nil
}

func verifySQLiteDB(path string) error {
	db, err := sql.Open("sqlite", fmt.Sprintf(
		"file:%s?mode=ro&immutable=1",
		path,
	))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec("PRAGMA quick_check"); err != nil {
		return err
	}

	return nil
}

func InitCardPrices() error {

	err := card.FetchHistoricPrices()

	if err != nil {
		return err
	}

	err = card.SetupPriceTable()

	if err != nil {
		return err
	}

	return nil
}
