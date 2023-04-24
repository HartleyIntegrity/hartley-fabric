package main

import (
	"log"
	"net/http"

	"github.com/yourusername/hartley-fabric/backend/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
