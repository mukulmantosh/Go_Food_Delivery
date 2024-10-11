package common

import (
	"Go_Food_Delivery/pkg/service/restaurant/unsplash"
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
)

type MockHTTPClient struct{}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("image data")),
	}, nil
}

type MockFileSystem struct {
	Files map[string]*bytes.Buffer
}

func (m *MockFileSystem) Create(name string) (*os.File, error) {
	m.Files[name] = &bytes.Buffer{}
	return os.Create(name)
}

func TestDownloadImageToDisk(t *testing.T) {
	mockClient := &MockHTTPClient{}
	mockFS := &MockFileSystem{Files: make(map[string]*bytes.Buffer)}

	err := unsplash.DownloadImageToDisk(mockClient, mockFS, "https://example.com/image.jpg", "/tmp/image.jpg")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validate that the file was created
	if _, exists := mockFS.Files["/tmp/image.jpg"]; !exists {
		t.Fatalf("expected file to be created, but it wasn't")
	}
}
