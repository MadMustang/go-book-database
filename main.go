package main

import (
	"fmt"
	"go-book-database/routes"
	"log"
	"net/http"
	"os"
)

//Main Server Function
func main() {

	//Initialize router
	r := routes.NewRouter()

	//Specify server port
	port := os.Getenv("Server_Port") //Grabs the port number from the dotenv

	//Adjust to 3000 if Server_Port not defined
	if port == "" {
		port = "3000"
	}

	//Running the server
	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
