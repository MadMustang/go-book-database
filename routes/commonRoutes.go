package routes

import (
	"fmt"
	"go-book-database/models"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Books : Declare list of books
var Books []models.Book

// Database object
var Db *sql.DB

//Init the data
func initDat() *sql.DB {
	Books = append(Books, models.Book{
		ID:     "1",
		Title:  "IT",
		Author: "Stephen King",
	})

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "rionaldy:woozie1210@tcp(127.0.0.1:3306)/acme")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened!!")
	}

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()

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
