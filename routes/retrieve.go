package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetBooks : Retrieving all the books in the database
func GetBooks(w http.ResponseWriter, r *http.Request) {

	//Send data
	w.Header().Set("Content-Type", "application/json") //Response content header
	w.WriteHeader(http.StatusOK)                       //Http status code
	json.NewEncoder(w).Encode(Books)
}

// GetBook : Getting a specific book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {

	//Get parameters from request
	params := mux.Vars(r)

	//Itterate & find
	for _, item := range Books {
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
