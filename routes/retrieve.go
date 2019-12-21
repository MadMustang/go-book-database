package routes

import (
	"encoding/json"
	"net/http"

	"github.com/MadMustang/go-book-database/models"

	"github.com/gorilla/mux"
)

// GetBooks : Retrieving all the books in the database
func GetBooks(w http.ResponseWriter, r *http.Request) {

	// Initiate local variable for data
	var retrievedData []models.Book

	// Retrieve data from database
	results, err := Db.Query("SELECT id, title, author FROM books")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Gather the data
	for results.Next() {

		var qu models.Book

		err := results.Scan(&qu.ID, &qu.Title, &qu.Author)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"Response": "Ay Blin. An error occured.",
			})
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		retrievedData = append(retrievedData, qu)

	}

	//Send data
	w.Header().Set("Content-Type", "application/json") //Response content header
	w.WriteHeader(http.StatusOK)                       //Http status code
	json.NewEncoder(w).Encode(retrievedData)
}

// GetBook : Getting a specific book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {

	//Get parameters from request
	params := mux.Vars(r)

	// Query Database
	results, err := Db.Query("SELECT id, title, author FROM books WHERE id=" + params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"Response": "Ay Blin. An error occured.",
		})
		panic(err.Error()) // proper error handling instead of panic in your app
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
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	// If the specific book exists
	if qu.ID != "" {

		//Return the found book
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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
