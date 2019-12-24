/*

   Insert desc.

   Project: go-book-database
   Package: server

   Created by rionaldytriananto on December 22, 2019
           at 10.55
*/
package server

import (
	"context"
	"github.com/MadMustang/go-book-database/routes"
	"go.uber.org/fx"
	"net/http"
)

// Exportable server module
var Module = fx.Options(
		fx.Invoke(initServer),
	)

// Initialize server
func initServer(lifecycle fx.Lifecycle) {

	// Server endpoints
	endpoints := routes.NewRouter()

	// Server object
	server := &http.Server{
		Handler: endpoints,
		Addr: ":3000",
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop:  func(ctx context.Context) error {
			return server.Close()
		},
	})

}