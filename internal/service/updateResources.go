package service

import (
	"os"
	"path"

	"githube.com/madeinly/cards/internal/card"
)

func FetchCardsDB() error {

	mtgjsonURL := "https://mtgjson.com/api/v5/AllPrintings.sqlite.gz"

	cardsPath, err := card.CardsPath()

	if err != nil {
		return err
	}

	gzPath := path.Join(cardsPath, "AllPrintings.sqlite.gz")
	err = card.DownLoadFile(mtgjsonURL, gzPath)

	defer os.Remove(gzPath)

	if err != nil {
		return err
	}

	pathMtgDB := path.Join(cardsPath, "mtgDB.sqlite")

	err = card.Ungz(gzPath, pathMtgDB)

	if err != nil {
		return err
	}

	err = os.Chmod(pathMtgDB, 0600)

	if err != nil {
		return err
	}

	return nil

}
