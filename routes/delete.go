package routes

import (
	"encoding/json"
	"github.com/MadMustang/go-book-database/models"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteBook : Deleting an existing book from the database/library
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	// Database
	db := initDat()
	defer db.Close()

	//Get Parameters from request
	params := mux.Vars(r)

	// Query Database
	results, err := db.Query("SELECT id, title, author FROM books WHERE id=" + params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"Response": "Ay Blin. An error occured.",
		})
	}

	// Initiate local variable
	var qu models.Book

	// Process data
	for results.Next() {

		err := results.Scan(&qu.ID, &qu.Title, &qu.Author)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"Response": "Ay Blin. An error occured.",
			})
		}
	}

	// Modify entry in database
	if qu.ID != "" {

		_, errChange := db.Query("DELETE FROM books WHERE id=" + params["id"])
		if errChange != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"Response": "Ay Blin. An error occured.",
			})
		}

		//Return the data of the deleted book
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(qu)
		return

	}

	//If the book is not found, send error response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"Response": "No such book in this library",
	})
}
