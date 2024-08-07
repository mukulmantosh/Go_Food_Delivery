package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetUnSplashImageURL(menuItem string) string {

	url := "https://api.unsplash.com/search/photos/?page=1&query=" + menuItem + "&w=400&h=400"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Client-ID", os.Getenv("UNSPLASH_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Decode the JSON response into the struct
	var apiResponse UnSplash
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	return apiResponse.Results[0].Urls.Small
}

func DownloadImageToDisk(url string, filepath string) error {
	// Send a GET request to the image URL
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: status code %d", response.StatusCode)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}
