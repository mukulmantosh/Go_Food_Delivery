package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func GetUnSplashImageURL(client HttpClient, menuItem string) string {

	imageUrl := "https://api.unsplash.com/search/photos/?page=1&query=" + url.QueryEscape(menuItem) + "&w=400&h=400"
	req, err := http.NewRequest("GET", imageUrl, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Client-ID", os.Getenv("UNSPLASH_API_KEY")))

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
		log.Fatalf("UnSplash::Failed to decode JSON response: %v", err)
	}

	return apiResponse.Results[0].Urls.Small
}
