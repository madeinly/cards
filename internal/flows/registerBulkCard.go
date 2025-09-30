package flows

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/core"
)

func RegisterBulk(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	// 1. save the file (your existing code)

	fmt.Println("started bulk")
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
	fmt.Println("started reading and parsing")

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
		params = append(params, RegisterCardParams{
			ID:         row[0],
			Language:   row[1],
			Stock:      row[2],
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
