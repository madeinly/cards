package features

import (
	"database/sql"
	"fmt"
	"path"
	"sync"

	"github.com/madeinly/core"
	_ "modernc.org/sqlite"
)

var (
	cardsDB *sql.DB
	dbOnce  sync.Once
)

func InitCardsDB() (*sql.DB, error) {
	var initErr error

	dbOnce.Do(func() {
		cardsPath := core.FeaturePath("cards")

		dbPath := path.Join(cardsPath, "mtgDB.sqlite")

		cardsDB, initErr = sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro&immutable=1", dbPath))

		if initErr != nil {
			return
		}

		cardsDB.SetMaxOpenConns(1)
		cardsDB.SetMaxIdleConns(1)
		cardsDB.SetConnMaxLifetime(0)

		if initErr = cardsDB.Ping(); initErr != nil {
			cardsDB.Close()
			return
		}
	})

	return cardsDB, initErr
}

func GetCardsDB() *sql.DB {
	if cardsDB != nil {
		return cardsDB
	}

	cardsDB, err := InitCardsDB()

	if err != nil {
		core.Fatal(err, "could not connect to cards database")
	}

	return cardsDB
}
