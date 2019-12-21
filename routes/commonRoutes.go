package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Database object
var Db *sql.DB

// Init the data
func initDat() *sql.DB {

	// Open up our database connection.
	// The database is called acme
	db, err := sql.Open("mysql", "rionaldy:password@tcp(127.0.0.1:3306)/acme")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened!!")
	}

	// return db object
	return db
}

// NewRouter initiates a new router
func NewRouter() *mux.Router {

	// Initialize the data
	Db = initDat()

	// Start new router
	router := mux.NewRouter()

	// Define sub-routers base on path
	baseRoute := router.PathPrefix(baseEndpoint).Subrouter()
	booksRoute := router.PathPrefix(bookEndpoint).Subrouter()

	// Define handlers for each endpoint path
	baseRoute.HandleFunc("", hello).Methods(http.MethodGet)
	booksRoute.HandleFunc("", GetBooks).Methods(http.MethodGet)
	booksRoute.HandleFunc("/{id}", GetBook).Methods(http.MethodGet)
	booksRoute.HandleFunc("/{id}", CreateBook).Methods(http.MethodPost)
	booksRoute.HandleFunc("/{id}", UpdateBook).Methods(http.MethodPut)
	booksRoute.HandleFunc("/{id}", DeleteBook).Methods(http.MethodDelete)

	//Return the router
	return router
}
