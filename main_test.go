package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

var app App

func TestMain(m *testing.M) {
	app = App{}
	var conf Config
	_, err := toml.DecodeFile("config/app.toml", &conf)
	if err != nil {
		log.Fatal(err)
	}

	app.Initialize(conf.User, conf.Password, conf.DBname, conf.SSLmode)

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM books")
	app.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT books_pkey PRIMARY KEY (id)
)`

func ExecuteR(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func StatusCheker(t *testing.T, expect int, origin int) {
	if expect != origin {
		t.Error("Expected value of status code does not match origin.")
	}

}

func TestEmpty(t *testing.T) {
	clearTable()

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := ExecuteR(req)

	StatusCheker(t, 200, resp.Code)

	if resp.Body.String() != "[]" {
		t.Error("Body not empty", resp.Body.String())
	}

}

func TestNotExists(t *testing.T) {
	clearTable()
	req, err := http.NewRequest("GET", "/books/11", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := ExecuteR(req)
	StatusCheker(t, 404, resp.Code)

}
