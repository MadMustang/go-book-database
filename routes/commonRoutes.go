package routes

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// NewRouter initiates a new router
func NewRouter() *mux.Router {

	// Start new router
	router := mux.NewRouter()

	// Define sub-routers base on path
	baseRoute := router.PathPrefix(baseEndpoint).Subrouter()
	booksRoute := router.PathPrefix(bookEndpoint).Subrouter()

	// Define handlers for each endpoint path
	baseRoute.HandleFunc("", hello).Methods(http.MethodGet)
	booksRoute.HandleFunc("", GetBooks).Methods(http.MethodGet)
	booksRoute.HandleFunc("/{id}", GetBook).Methods(http.MethodGet)
	booksRoute.HandleFunc("/{id}", CreateBook).Methods(http.MethodPost)
	booksRoute.HandleFunc("/{id}", UpdateBook).Methods(http.MethodPut)
	booksRoute.HandleFunc("/{id}", DeleteBook).Methods(http.MethodDelete)

	//Return the router
	return router
}
