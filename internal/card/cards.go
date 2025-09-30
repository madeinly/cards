package card

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/madeinly/core"
)

func ImportsPath() string {

	CardsPath := core.FeaturePath("cards")

	importsPath := core.FeaturePath(path.Join(CardsPath, "imports"))

	return importsPath

}

func DownLoadFile(url string, path string) error {

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func Ungz(gzPath string, filePath string) error {

	gzFile, err := os.Open(gzPath)
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gzReader, err := gzip.NewReader(gzFile)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	cardsDB, err := os.Create(filePath)

	if err != nil {
		return err
	}
	defer cardsDB.Close()

	_, err = io.Copy(cardsDB, gzReader)

	if err != nil {
		return err
	}

	return nil

}

func FetchHistoricPrices() error {

	url := "https://mtgjson.com/api/v5/AllPrices.json.gz"

	cardsPath := core.FeaturePath("cards")

	gzFile := path.Join(cardsPath, "AllPrices.json.gz")

	err := DownLoadFile(url, gzFile)

	if err != nil {
		return err
	}

	pricesPath := path.Join(cardsPath, "mtgPrices.json")

	err = Ungz(gzFile, pricesPath)

	if err != nil {
		return err
	}

	defer os.Remove(gzFile)

	err = os.Chmod(pricesPath, 0600)

	if err != nil {
		return err
	}

	return nil

}
