package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalFileStorage struct {
	BasePath string
}

func (lf *LocalFileStorage) Upload(fileName string, file io.Reader) (string, error) {
	fullPath := filepath.Join(lf.BasePath, fileName)
	outFile, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}
