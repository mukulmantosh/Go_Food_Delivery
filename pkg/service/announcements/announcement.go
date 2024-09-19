package announcements

import (
	"Go_Food_Delivery/pkg/abstract/announcements"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func (eventSrv *AnnouncementService) FlashEvents() (*[]announcements.FlashEvents, error) {
	currentDir, err := getCurrentDirectory()
	if err != nil {
		return nil, errors.New("unable to get current directory")
	}
	fileName := filepath.Join(currentDir, "events.json")

	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, errors.New("unable to open file")
	}
	defer file.Close()

	var newsEvents []announcements.FlashEvents
	err = json.NewDecoder(file).Decode(&newsEvents)
	if err != nil {
		return nil, errors.New("unable to decode JSON")
	}

	return &newsEvents, nil
}

func getCurrentDirectory() (string, error) {
	// Get the caller's information
	_, currentFile, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("unable to get current file path")
	}

	// Get the directory from the current file path
	currentDir := filepath.Dir(currentFile)

	return currentDir, nil
}
