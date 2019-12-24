package routes

import (
	"encoding/json"
	"github.com/MadMustang/go-book-database/models"
	"net/http"
)

// CreateBook : inserting a new book into the library
func CreateBook(w http.ResponseWriter, r *http.Request) {

	// Database
	db := initDat()
	defer db.Close()

	//Making the insertion from the request body
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	_, err := db.Query("INSERT INTO books(title, author) VALUES ('" + book.Title + "', '" + book.Author + "');")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"Response": "Ay Blin. An error occured.",
		})
	}

	//Returns the created index
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
