package main

import (
	"database/sql"
	"errors"
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
	var id int
	db.QueryRow("UPDATE books SET name = $1, price = $2  WHERE id = $3 returning id", b.Name, b.Price, b.ID).Scan(&id)
	if id != b.ID {
		b.createBook(db)
		return errors.New("Book with this id not found. Successfully created new!")
	}
	return nil
}

//createBook
func (b *Book) createBook(db *sql.DB) error {

	err := db.QueryRow(
		"INSERT INTO books(name, price) VALUES($1, $2) RETURNING id",
		b.Name, b.Price).Scan(&b.ID)
	if err != nil {
		return err
	}
	return nil

}

//deleteBook
func (b *Book) deleteBook(db *sql.DB) error {
	var id int
	db.QueryRow("DELETE FROM books WHERE id = $1 returning id", b.ID).Scan(&id)
	if id != b.ID {
		return errors.New("Object not found")
	}
	return nil
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
