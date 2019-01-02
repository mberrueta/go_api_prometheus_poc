package db

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/mberrueta/go_api_prometheus_poc/models"

	_ "github.com/lib/pq"
)

var pool *sql.DB

func RowToBook(row *sql.Row) *models.Book {
	bk := models.Book{}
	row.Scan(&bk.ID, &bk.Title, &bk.AuthorId)

	return &bk
}

func RowsToBooks(rows *sql.Rows) *[]models.Book {
	var books []models.Book

	for rows.Next() {
		bk := models.Book{}
		err := rows.Scan(&bk.ID, &bk.Title, &bk.AuthorId)
		if err != nil {
			log.Fatalf("query fail: %v", err)
		}
		books = append(books, bk)
	}

	return &books
}

func GetAll() *[]models.Book {
	rows, err := pool.Query("SELECT * FROM books;")
	if err != nil {
		log.Fatalf("Issue with select query: %v", err)
	}

	return RowsToBooks(rows)
}

func GetById(id uint16) (*models.Book, error) {
	rows, err := pool.Query("SELECT * FROM books WHERE id = $1;", id)
	if err != nil {
		log.Fatalf("Issue with select query: %v", err)
	}

	result := RowsToBooks(rows)

	if len(*result) == 0 {
		return &models.Book{}, errors.New("Not found")
	}

	if len(*result) > 1 {
		return &models.Book{}, errors.New("More than 1 result")
	}

	return &(*result)[0], nil
}

func Create(title string, authorId uint16) (uint16, error) {
	result, err := pool.Exec("INSERT INTO books VALUES($1, $2);", title, authorId)
	if err != nil {
		log.Fatalf("Fail to insert: %v", err)
	}

	id, err := result.LastInsertId()

	return uint16(id), err
}

func init() {
	var err error
	pool, err = sql.Open("postgres", "postgres://postgres@localhost/go_api_prometheus_poc_dev?sslmode=disable")
	// defer pool.Close()
	if err != nil {
		log.Fatalf("Wrong connection string: %v", err)
	}

	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(time.Hour)
	pool.SetMaxIdleConns(5)

	err = pool.Ping()
	if err != nil {
		log.Fatalf("Db ping fail: %v", err)
	}
}

// func Open() []models.Book {
// 	db, err := sql.Open("postgres", "postgres://postgres@localhost/go_api_prometheus_poc_dev?sslmode=disable")
// 	defer db.Close()
// 	if err != nil {
// 		log.Fatalf("Wrong connection string: %v", err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Db ping fail: %v", err)
// 	}

// 	rows, err := db.Query("SELECT * FROM books;")

// 	err := db.QueryRow("SELECT * FROM users WHERE name = $1", name).Scan(&u.Id, &u.Name, &u.Score)

// 	if err != nil {
// 		log.Fatalf("query fail: %v", err)
// 	}
// 	defer rows.Close()

// 	var books []models.Book
// 	for rows.Next() {
// 		bk := models.Book{}
// 		err := rows.Scan(&bk.ID, &bk.Title, &bk.AuthorId)
// 		if err != nil {
// 			log.Fatalf("query fail: %v", err)
// 		}
// 		books = append(books, bk)
// 	}

// 	return books
// }
