package models

import (
	"errors"
	"strconv"
)

var books []Book

type Book struct {
	ID     uint16  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func GetBooks() ([]Book, error) {
	if len(books) == 0 {
		mockData()
	}
	return books, nil
}

func GetBook(id uint16) (Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, errors.New("Book with id " + strconv.Itoa(int(id)) + " not found")
}

func CreateBook(book Book) uint16 {
	book.ID = uint16(len(books) + 1)
	books = append(books, book)
	return book.ID
}

func (b Book) Delete() {
	for i, book := range books {
		if book.ID == b.ID {
			books = append(books[:i], books[i+1:]...)
		}
	}
}

func (b Book) Update(params map[string]string) {
	for _, book := range books {
		if book.ID == b.ID {
			book.Title = params["title"]
		}
	}
}

func mockData() {
	books = append(books, Book{
		ID:    1,
		Title: "Saving dogs",
		Author: &Author{
			FirstName: "John",
			LastName:  "Wick",
		},
	})

	books = append(books, Book{
		ID:    2,
		Title: "Hiding",
		Author: &Author{
			FirstName: "Carmen",
			LastName:  "Sandiego",
		},
	})
}
