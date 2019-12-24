package routes

import (
	"encoding/json"
	"github.com/MadMustang/go-book-database/models"
	"net/http"

	"github.com/gorilla/mux"
)

// UpdateBook : Update the book's information
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	// Database
	db := initDat()
	defer db.Close()

	//Get Parametters from request
	params := mux.Vars(r)

	//Making the insertion from the request body
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Check for the book in the database
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

		_, errChange := db.Query("UPDATE books SET title='" + book.Title + "', author='" + book.Author + "' WHERE id=" + params["id"])
		if errChange != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"Response": "Ay Blin. An error occured.",
			})
		}

		//Return the data of the updated book
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(book)
		return

	}

	//If the book is not found, send error response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"Response": "No such book in this library",
	})
}
