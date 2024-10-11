package unsplash

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ImageClient interface {
	Get(url string) (*http.Response, error)
}

type FileSystem interface {
	Create(name string) (*os.File, error)
}

type DefaultHTTPImageClient struct{}

func (c *DefaultHTTPImageClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

type DefaultFileSystem struct{}

func (fs *DefaultFileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func DownloadImageToDisk(client ImageClient, fs FileSystem, url string, filepath string) error {
	response, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: status code %d", response.StatusCode)
	}

	file, err := fs.Create(filepath)
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
