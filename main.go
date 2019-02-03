package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Main Server Function
func main() {

	//Initialize router
	r := mux.NewRouter()

	//Specify server port
	port := os.Getenv("Server_Port") //Grabs the port number from the dotenv

	//Validate that the port exists in the config file
	if port == "" {
		port = "3000"
	}

	//Test mock data
	books = append(books, Book{
		ID:     "1",
		Title:  "IT",
		Author: "Stephen King",
	})

	//Define handlers and their functions
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", GetBook).Methods("GET")
	r.HandleFunc("/book", CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")

	//Running the server
	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

//Test Handler
func hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//I used a map
	greet := map[string]string{
		"salutations": "Hello there",
		"statement":   "You seemed to be lost",
		"instruction": "Read the docs please",
	}

	json.NewEncoder(w).Encode(greet)
}

// Book object
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// GetBooks : Retrieving all the books in the database
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Response content header
	w.WriteHeader(http.StatusOK)                       //Http status code
	json.NewEncoder(w).Encode(books)
}

// GetBook : Getting a specific book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {

	//Get parameters from request
	params := mux.Vars(r)

	//Itterate & find
	for _, item := range books {
		if item.ID == params["id"] {

			//Return the found book
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	//If the book is not found, send error response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"Response": "No such book in this library",
	})
}

// CreateBook : inserting a new book into the library
func CreateBook(w http.ResponseWriter, r *http.Request) {

	//Making the insertion from the request body
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book) //Argument: the list, the new item

	//Returns the created index
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// DeleteBook : Deleting an existing book from the database/library
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	//Get Parametters from request
	params := mux.Vars(r)

	//Itterate & find
	for index, item := range books {
		if item.ID == params["id"] {

			//Getting the data of the book
			b := item

			//Delete book from databse
			books = append(books[:index], books[index+1:]...)

			//Return the data of the deleted book
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	//If the book is not found, send error response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"Response": "No such book in this library",
	})
}

// UpdateBook : Update the book's information
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	//Get Parametters from request
	params := mux.Vars(r)

	//Making the insertion from the request body
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	//Itterate & find
	for index, item := range books {
		if item.ID == params["id"] {

			//Delete the book
			books = append(books[:index], books[index+1:]...)

			//Update the book
			book.ID = item.ID
			books = append(books, book)

			//Return the data of the updated book
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	//If the book is not found, send error response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"Response": "No such book in this library",
	})
}
