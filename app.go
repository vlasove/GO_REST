package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize ...
func (a *App) Initialize(user, password, dbname, sslmode string) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)
	var err error
	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run ...
func (a *App) Run(port string) {}
