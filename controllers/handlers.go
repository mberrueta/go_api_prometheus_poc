package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mberrueta/go_api_prometheus_poc/models"
)

func BooksIndex(w http.ResponseWriter, r *http.Request) {
	list, err := models.GetBooks()

	render(w, list, err)
}

func BooksShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sId := params["id"]
	id, err := strconv.ParseUint(sId, 10, 16)

	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	book, err := models.GetBook(uint16(id))

	render(w, book, err)
}

func BooksCreate(w http.ResponseWriter, r *http.Request) {
	var book *models.Book = &models.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	models.CreateBook(*book)
}

func BooksDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sId := params["id"]
	id, err := strconv.ParseUint(sId, 10, 16)

	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	book, err := models.GetBook(uint16(id))

	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	book.Delete()
}

func BooksUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sId := params["id"]
	id, err := strconv.ParseUint(sId, 10, 16)

	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	book, err := models.GetBook(uint16(id))

	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	book.Update(params)
}

func render(w http.ResponseWriter, v interface{}, err error) {
	if err != nil {
		w.Write([]byte("422: " + err.Error() + http.StatusText(422)))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
