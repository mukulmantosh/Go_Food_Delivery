package storage

import (
	"log"
	"os"
)

func createUploadDirectory(basePath string) {
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}

	err := os.Chmod(basePath, 0755)
	if err != nil {
		log.Fatalf("Permission Error:: %v", err)
		return
	}
}
