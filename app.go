package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"net/http"

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
	a.findRoutes()
}

// Run ...
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

//getBook
func (a *App) getBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondError(w, 400, "Invalid ID")
		return
	}

	b.ID = id
	if err := b.getBook(a.DB); err != nil {
		RespondError(w, 404, "Can not be found in DB")
		return
	}

	RespondJSON(w, 200, b)
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {

	booksSlice, err := getBooks(a.DB)
	if err != nil {
		RespondError(w, 500, "BAD REQUEST")
		return
	}
	RespondJSON(w, 200, booksSlice)

}

func (a *App) createBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	fmt.Println("POST REQUEST STARTED")
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&b); err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()

	if err := b.createBook(a.DB); err != nil {
		RespondError(w, 500, err.Error())
		return
	}
}

func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondError(w, 400, "Invalid ID")
		return
	}

	b.ID = id
	if err := b.deleteBook(a.DB); err != nil {
		RespondError(w, 404, "Can not be found in DB")
		return
	}

	RespondJSON(w, 200, map[string]string{"message": "successfull deleted"})
}

func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondError(w, 400, "Invalid ID")
		return
	}

	dec := json.NewDecoder(r.Body)
	_ = dec.Decode(&b)

	b.ID = id
	if err := b.updateBook(a.DB); err != nil {
		RespondError(w, 201, err.Error())
		return
	}

	RespondJSON(w, 200, map[string]string{"message": "successfull "})

}

func (a *App) findRoutes() {
	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
	a.Router.HandleFunc("/books/{id}", a.getBook).Methods("GET")
	a.Router.HandleFunc("/books/{id}", a.updateBook).Methods("PUT")
	a.Router.HandleFunc("/books/{id}", a.deleteBook).Methods("DELETE")
	a.Router.HandleFunc("/books", a.createBook).Methods("POST")
}

func RespondError(w http.ResponseWriter, statusCode int, message string) {
	RespondJSON(w, statusCode, map[string]string{"error": message})
}

func RespondJSON(w http.ResponseWriter, statusCode int, p interface{}) {
	resp, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
