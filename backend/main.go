package main

import (
	"github.com/HartleyIntegrity/hartley-fabric/backend/api"
	"github.com/HartleyIntegrity/hartley-fabric/backend/blockchain"
	"github.com/HartleyIntegrity/hartley-fabric/backend/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize the blockchain and database
	bc := blockchain.NewBlockchain()
	db := database.NewDatabase()

	// Register the API handlers
	api.RegisterHandlers(router)

	// Set up JWT middleware
	router.Use(api.JWTMiddleware("your-secret-key"))

	// Start the server
	router.Run(":8080")
}
