package routes

import (
	"encoding/json"
	"go-book-database/models"
	"net/http"

	"github.com/gorilla/mux"
)

// UpdateBook : Update the book's information
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	//Get Parametters from request
	params := mux.Vars(r)

	//Making the insertion from the request body
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	//Itterate & find
	for index, item := range Books {
		if item.ID == params["id"] {

			//Delete the book
			Books = append(Books[:index], Books[index+1:]...)

			//Update the book
			book.ID = item.ID
			Books = append(Books, book)

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
