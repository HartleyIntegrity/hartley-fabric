package main

import (
	"hartley-fabric/backend/api"
	"hartley-fabric/backend/blockchain"
	"hartley-fabric/backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create a new CORS configuration.
	config := cors.DefaultConfig()

	// Set the allowed origins.
	config.AllowOrigins = []string{"http://localhost:3000"}

	// Set the allowed methods.
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	// Set the allowed headers.
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	// Use the CORS middleware.
	r.Use(cors.New(config))

	// Create a new database instance.
	db := database.NewDatabase()

	// Create a new blockchain instance.
	bc := blockchain.NewBlockchain()

	// Register the API handlers.
	api.RegisterHandlers(r, db, bc)

	// Start the server.
	r.Run(":8080")
}
