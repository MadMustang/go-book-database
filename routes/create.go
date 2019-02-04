package routes

import (
	"encoding/json"
	"go-book-database/models"
	"net/http"
)

// CreateBook : inserting a new book into the library
func CreateBook(w http.ResponseWriter, r *http.Request) {

	//Making the insertion from the request body
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	Books = append(Books, book) //Argument: the list, the new item

	//Returns the created index
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
