package card

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// returns the folder path of the bin (probably should be in core)
func MadeinlyPath() (string, error) {

	binPath, err := os.Executable()

	if err != nil {
		return "", err
	}

	return filepath.Dir(binPath), nil

}

// makes sure that the cards folder exist and if not created or return error
func CardsPath() (string, error) {

	madeinlyPath, err := MadeinlyPath()

	if err != nil {
		return "", err
	}

	cardsFolderPath := path.Join(madeinlyPath, "cards")

	_, err = os.Stat(cardsFolderPath)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(cardsFolderPath, 0755)
			if err != nil {
				return "", fmt.Errorf("failed to create directory: %w", err)
			}
		} else {
			return "", fmt.Errorf("failed to check directory: %w", err)
		}
	}

	return cardsFolderPath, nil
}

func ImportsPath() string {

	CardsPath, err := CardsPath()

	if err != nil {
		fmt.Println(err.Error())
		panic(0)
	}

	importsPath := path.Join(CardsPath, "imports")

	_, err = os.Stat(importsPath)

	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(importsPath, 0755)

		if err != nil {
			panic(0)
		}

		return importsPath
	}

	if err != nil {
		panic(0)
	}

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

	cardsPath, err := CardsPath()

	if err != nil {
		return err
	}

	gzFile := path.Join(cardsPath, "AllPrices.json.gz")

	err = DownLoadFile(url, gzFile)

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
