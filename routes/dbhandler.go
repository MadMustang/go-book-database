/*

   Insert desc.

   Project: go-book-database
   Package: routes

   Created by rionaldytriananto on December 22, 2019
           at 12.09
*/
package routes

import (
	"database/sql"
)

// Database object
var Db *sql.DB

// Init the database
func initDat() *sql.DB {

	// Open up our database connection.
	// The database is called acme
	db, err := sql.Open("mysql", "rionaldy:password@tcp(127.0.0.1:3306)/acme")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// return db object
	return db
}