package routes

import (
	"go-book-database/models"

	"github.com/gorilla/mux"
)

//Books : Declare list of books
var Books []models.Book

//Init the data
func initDat() {
	Books = append(Books, models.Book{
		ID:     "1",
		Title:  "IT",
		Author: "Stephen King",
	})
}

//NewRouter initiates a new router
func NewRouter() *mux.Router {

	//Initialize the data
	initDat()

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
