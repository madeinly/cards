package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"path"
	"sync"

	"githube.com/madeinly/cards/internal/card"
	_ "modernc.org/sqlite"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func InitCardsDB() (*sql.DB, error) {
	var initErr error

	dbOnce.Do(func() {
		cardsPath, err := card.CardsPath()
		if err != nil {
			initErr = err
			return
		}

		dbPath := path.Join(cardsPath, "mtgDB.sqlite")

		db, initErr = sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro&immutable=1", dbPath))
		if initErr != nil {
			return
		}

		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(1)
		db.SetConnMaxLifetime(0)

		if initErr = db.Ping(); initErr != nil {
			db.Close()
			return
		}
	})

	fmt.Println("database online")
	return db, initErr
}

func GetCardsDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("database not initialized - call InitDB first")
	}
	return db, nil
}
