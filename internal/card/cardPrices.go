package card

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/madeinly/core"
)

type MtgJsonModel struct {
	Meta MtgJsonMeta         `json:"meta"`
	Data map[string]CardType `json:"data"`
}

type MtgJsonMeta struct {
	Date    string `json:"date"`
	Version string `json:"version"`
}

type CardType struct {
	Paper Vendor `json:"paper"`
}

type Vendor struct {
	Cardkingdom vendorList `json:"cardkingdom"`
}

type vendorList struct {
	Retail finish `json:"retail,omitempty"`
	// Buylist finish `json:"buylist,omitempty"`
}

type finish struct {
	Foil   price `json:"foil,omitempty"`
	Normal price `json:"normal,omitempty"`
	Etched price `json:"etched,omitempty"`
}

type price map[string]float64

type PriceRecord struct {
	CardID string
	Finish string // 'normal', 'foil', or 'etched'
	Type   string // always 'retail' in your case
	Price  float64
	Date   string
}

const OptimalBatchSize = 500 // Good balance between speed and memory

func SetupPriceTable() error {
	start := time.Now()

	// Memory measurement before
	var mBefore, mAfter runtime.MemStats
	runtime.ReadMemStats(&mBefore)

	cardsPath, err := CardsPath()
	if err != nil {
		return err
	}

	pricesPath := path.Join(cardsPath, "mtgPrices.json")
	pricesFile, err := os.Open(pricesPath)
	if err != nil {
		return err
	}
	defer pricesFile.Close()

	var jsonModel MtgJsonModel
	decoder := json.NewDecoder(pricesFile)
	if err := decoder.Decode(&jsonModel); err != nil {
		return fmt.Errorf("JSON decode failed: %w", err)
	}

	// Process records in batches
	records := BuildPriceRecords(jsonModel)
	const batchSize = 500
	totalRecords := len(records)

	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		batch := records[i:end]
		if err := BatchInsertPrices(batch); err != nil {
			return fmt.Errorf("batch insert failed at batch %d-%d: %w", i, end, err)
		}

		// Print progress every 10 batches
		if (i/batchSize)%10 == 0 {
			fmt.Printf("Processed %d/%d records (%.1f%%)\n",
				end, totalRecords, float64(end)/float64(totalRecords)*100)
		}
	}

	// Memory measurement after processing
	runtime.GC()
	runtime.ReadMemStats(&mAfter)

	// Print stats
	fmt.Println("\nMeta:", jsonModel.Meta)
	fmt.Printf("Memory used: %d bytes (%.2f MB)\n",
		mAfter.HeapAlloc-mBefore.HeapAlloc,
		float64(mAfter.HeapAlloc-mBefore.HeapAlloc)/1024/1024)
	fmt.Printf("Processed %d records in %v\n", totalRecords, time.Since(start))

	return nil
}

func BuildPriceRecords(mtgJson MtgJsonModel) []PriceRecord {

	data := mtgJson.Data

	var Record []PriceRecord

	for cardID, cardData := range data {

		normalSet := cardData.Paper.Cardkingdom.Retail.Normal
		foilSet := cardData.Paper.Cardkingdom.Retail.Foil
		etchedSet := cardData.Paper.Cardkingdom.Retail.Etched

		if normalSet != nil {

			price, date := getMostRecentPrice(normalSet)

			Record = append(Record, PriceRecord{
				CardID: cardID,
				Finish: "normal",
				Type:   "retail",
				Price:  price,
				Date:   date,
			})
		}

		if foilSet != nil {

			price, date := getMostRecentPrice(foilSet)

			Record = append(Record, PriceRecord{
				CardID: cardID,
				Finish: "foil",
				Type:   "retail",
				Price:  price,
				Date:   date,
			})
		}

		if etchedSet != nil {

			price, date := getMostRecentPrice(foilSet)

			Record = append(Record, PriceRecord{
				CardID: cardID,
				Finish: "etched",
				Type:   "retail",
				Price:  price,
				Date:   date,
			})
		}

	}

	return Record
}

func getMostRecentPrice(prices map[string]float64) (float64, string) {
	var latestDate string
	var latestPrice float64
	first := true

	for date, price := range prices {
		if first || date > latestDate { // String comparison works for YYYY-MM-DD
			latestDate = date
			latestPrice = price
			first = false
		}
	}
	return latestPrice, latestDate
}

// BatchInsertPrices inserts a batch of price records efficiently
func BatchInsertPrices(batch []PriceRecord) error {
	if len(batch) == 0 {
		return nil
	}

	// Begin transaction
	tx, err := core.DB().Begin()
	if err != nil {
		return fmt.Errorf("transaction begin failed: %w", err)
	}
	defer tx.Rollback() // Will be ignored if tx.Commit() succeeds

	// Build the bulk insert query
	query, args := buildBulkInsertQuery(batch)

	// Execute the query
	_, err = tx.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("bulk insert failed: %w", err)
	}

	// Commit transaction
	return tx.Commit()
}

// buildBulkInsertQuery constructs the SQL and arguments for bulk insert
func buildBulkInsertQuery(batch []PriceRecord) (string, []interface{}) {
	var valueStrings []string
	var valueArgs []interface{}

	const valuesPerRow = 4 // card_id, finish, type, price
	valueArgs = make([]interface{}, 0, len(batch)*valuesPerRow)

	for i, record := range batch {
		// Add placeholders for this row
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)",
			i*valuesPerRow+1, i*valuesPerRow+2, i*valuesPerRow+3, i*valuesPerRow+4))

		// Add the actual values
		valueArgs = append(valueArgs,
			record.CardID,
			record.Finish,
			record.Type,
			record.Price)
	}

	query := fmt.Sprintf(`
		INSERT INTO cards_price 
		(card_id, finish, type, price)
		VALUES %s
		ON CONFLICT(card_id, finish, type) 
		DO UPDATE SET 
			price = excluded.price,
			updated_at = CURRENT_TIMESTAMP`,
		strings.Join(valueStrings, ","))

	return query, valueArgs
}
