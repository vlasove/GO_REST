package main

import (
	"database/sql"
	"fmt"
)

//Book ...
type Book struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//getBook ...
func (b *Book) getBook(db *sql.DB) error {
	err := db.QueryRow("SELECT * FROM books WHERE id=$1", b.ID).Scan(&b.ID, &b.Name, &b.Price)
	return err
}

//updateBook
func (b *Book) updateBook(db *sql.DB) error {
	_, err := db.Exec("UPDATE books SET name = $1, price = $2  WHERE id = $3", b.Name, b.Price, b.ID)
	return err
}

//createBook
func (b *Book) createBook(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO books (name, price) VALUES ($1, $2)", b.Name, b.Price)
	return err
}

//deleteBook
func (b *Book) deleteBook(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", b.ID)
	return err
}

//getBooks ...
func getBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	books := make([]Book, 0)
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Name, &b.Price); err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}

	return books, nil
}
