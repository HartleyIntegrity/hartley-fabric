package main

import (
	"github.com/HartleyIntegrity/hartley-fabric/models"
	"github.com/HartleyIntegrity/hartley-fabric/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	auth := r.Group("/auth")
	{
		auth.POST("/signup", routes.Signup)
		auth.POST("/login", routes.Login)
	}

	api := r.Group("/api")
	api.Use(routes.AuthMiddleware())
	{
		api.POST("/contracts", routes.CreateContract)
		api.GET("/contracts", routes.GetContracts)
		api.PUT("/contracts/:id", routes.UpdateContract)
		api.DELETE("/contracts/:id", routes.DeleteContract)
	}

	r.Run(":8080")
}
