package db

import (
	"database/sql"
	"log"

	"github.com/mberrueta/go_api_prometheus_poc/models"

	_ "github.com/lib/pq"
)

func Open() []models.Book {
	db, err := sql.Open("postgres", "postgres://postgres@localhost/go_api_prometheus_poc_dev?sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalf("Wrong connection string: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Db ping fail: %v", err)
	}

	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		log.Fatalf("query fail: %v", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		bk := models.Book{}
		err := rows.Scan(&bk.ID, &bk.Title, &bk.AuthorId)
		if err != nil {
			log.Fatalf("query fail: %v", err)
		}
		books = append(books, bk)
	}

	return books
}
