package main

import (
	"database/sql"
	"errors"
)

//Book ...
type Book struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//getBook ...
func (b *Book) getBook(db *sql.DB) error {
	return errors.New("Not yet")
}

//updateBook
func (b *Book) updateBook(db *sql.DB) error {
	return errors.New("Not yet")
}

//createBook
func (b *Book) createBook(db *sql.DB) error {
	return errors.New("Not yet")
}

//deleteBook
func (b *Book) deleteBook(db *sql.DB) error {
	return errors.New("Not yet")
}

//getBooks ...
func getBooks(db *sql.DB) ([]Book, error) {
	return nil, errors.New("Not yet")
}
