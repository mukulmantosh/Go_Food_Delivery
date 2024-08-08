package restaurant

import (
	"bytes"
	"mime/multipart"
)

type FakeRestaurant struct {
	Name        string
	File        []byte
	Description string
	Address     string
	City        string
	State       string
}

func generateData(restaurant FakeRestaurant) (*bytes.Buffer, string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	_ = writer.WriteField("name", restaurant.Name)

	fileWriter, _ := writer.CreateFormFile("file", "restaurant.jpg")
	_, _ = fileWriter.Write(restaurant.File)

	_ = writer.WriteField("description", restaurant.Description)
	_ = writer.WriteField("address", restaurant.Address)
	_ = writer.WriteField("city", restaurant.City)
	_ = writer.WriteField("state", restaurant.State)

	_ = writer.Close()

	return &buffer, writer.FormDataContentType(), nil

}
