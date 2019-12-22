package main

import (
	"github.com/MadMustang/go-book-database/server"
	"go.uber.org/fx"
)

//Main Server Function
func main() {

	// Create and run instance
	app := fx.New(
			server.Module,
		)
	app.Run()
}
