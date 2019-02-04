package routes

import (
	"encoding/json"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//I used a map
	greet := map[string]string{
		"salutations": "Hello there",
		"statement":   "You seemed to be lost",
		"instruction": "Read the docs please",
	}

	//Return response
	json.NewEncoder(w).Encode(greet)
}
