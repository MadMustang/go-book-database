package routes

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Database object
var Db *sql.DB

//Init the data
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

//NewRouter initiates a new router
func NewRouter() *mux.Router {

	//Initialize the data
	Db = initDat()

	//Start new router
	router := mux.NewRouter()

	//Define handlers and their functions
	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/book", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")

	//Return the router
	return router
}
