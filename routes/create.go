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
	_, err := Db.Query("INSERT INTO books(title, author) VALUES ('" + book.Title + "', '" + book.Author + "');")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"Response": "Ay Blin. An error occured.",
		})
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//Returns the created index
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
