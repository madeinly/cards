package flows

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/madeinly/cards/internal/card"
	"github.com/madeinly/core"
)

func SetupPriceTable() error {
	start := time.Now()

	// Memory measurement before
	var mBefore, mAfter runtime.MemStats
	runtime.ReadMemStats(&mBefore)

	cardsPath := core.FeaturePath("cards")

	pricesPath := path.Join(cardsPath, "mtgPrices.json")
	pricesFile, err := os.Open(pricesPath)
	if err != nil {
		return err
	}
	defer pricesFile.Close()

	var jsonModel card.MtgJsonModel
	decoder := json.NewDecoder(pricesFile)
	if err := decoder.Decode(&jsonModel); err != nil {
		return fmt.Errorf("JSON decode failed: %w", err)
	}

	// Process records in batches
	records := card.BuildPriceRecords(jsonModel)
	const batchSize = 500
	totalRecords := len(records)

	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		batch := records[i:end]
		if err := card.BatchInsertPrices(batch); err != nil {
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
