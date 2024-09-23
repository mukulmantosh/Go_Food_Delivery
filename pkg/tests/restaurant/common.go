package restaurant

import (
	"bytes"
	"mime/multipart"
	"time"
)

type FakeRestaurant struct {
	Name        string
	File        []byte
	Description string
	Address     string
	City        string
	State       string
}

type MenuItem struct {
	MenuID       int       `json:"menu_id"`
	RestaurantID int       `json:"restaurant_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Photo        string    `json:"photo"`
	Price        float64   `json:"price"`
	Category     string    `json:"category"`
	Available    bool      `json:"available"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
}

func GenerateData(restaurant FakeRestaurant) (*bytes.Buffer, string, error) {
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
