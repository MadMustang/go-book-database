package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteBook : Deleting an existing book from the database/library
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	//Get Parametters from request
	params := mux.Vars(r)

	//Itterate & find
	for index, item := range Books {
		if item.ID == params["id"] {

			//Getting the data of the book
			b := item

			//Delete book from databse
			Books = append(Books[:index], Books[index+1:]...)

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
