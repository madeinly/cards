package features

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/madeinly/core"
)

func HandleImportFile(file multipart.File) (string, error) {

	importFolderPath := core.FeaturePath("cards/imports")
	now := time.Now()
	fileName := now.Format("2006-01-02_15-04-05") + ".csv"
	importFilePath := filepath.Join(importFolderPath, fileName)

	dst, err := os.Create(importFilePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	if err := dst.Close(); err != nil {
		return "", err
	}

	return importFilePath, nil
}
