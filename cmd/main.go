package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulii/Library/controller/books"
	"github.com/rahulii/Library/db"
	"github.com/rahulii/Library/pkg/config"
)


func main() {
	
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Initialize the database
	dbHandler, err := db.Init(config.DBUrl)
	if err != nil {
		panic(err)
	}

	// Initialize the books handler
	booksHandler := books.NewHandler(dbHandler)

	// Register the routes
	booksHandler.RegisterRoutes(router)

	// Start the server
	router.Run(config.Port)
}