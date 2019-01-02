package services

import (
	// "errors"
	// "strconv"

	"github.com/mberrueta/go_api_prometheus_poc/db"
	"github.com/mberrueta/go_api_prometheus_poc/models"
)

var books []models.Book

func GetBooks() ([]models.Book, error) {
	if len(books) == 0 {
		// mockData()
		books = *db.GetAll()
	}
	return books, nil
}

func GetBook(id uint16) (models.Book, error) {
	// for _, book := range books {
	// 	if book.ID == id {
	// 		return book, nil
	// 	}
	// }

	result, err := db.GetById(id)
	return *result, err
}

func CreateBook(book models.Book) uint16 {
	book.ID = uint16(len(books) + 1)
	books = append(books, book)
	return book.ID
}

func Delete(id uint16) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	return nil
}

func Update(id uint16, params map[string]string) error {
	for _, book := range books {
		if book.ID == id {
			book.Title = params["title"]
		}
	}

	return nil
}

func mockData() {
	books = append(books, models.Book{
		ID:    1,
		Title: "Saving dogs",
		Author: &models.Author{
			FirstName: "John",
			LastName:  "Wick",
		},
	})

	books = append(books, models.Book{
		ID:    2,
		Title: "Hiding",
		Author: &models.Author{
			FirstName: "Carmen",
			LastName:  "Sandiego",
		},
	})
}
