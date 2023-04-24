package api

import (
	"github.com/gorilla/mux"
	"github.com/hartley-fabric/hartley-fabric/backend/database"
)

// NewRouter returns a new router for our API.
func NewRouter(db database.Models) *mux.Router {
	// create a new router
	r := mux.NewRouter()

	// define the routes
	r.HandleFunc("/api/properties", PropertiesHandler(db)).Methods("GET")
	r.HandleFunc("/api/properties/{id}", PropertyHandler(db)).Methods("GET")
	r.HandleFunc("/api/contracts", ContractsHandler(db)).Methods("GET", "POST")
	r.HandleFunc("/api/contracts/{id}", ContractHandler(db)).Methods("DELETE")

	// add the authentication middleware to all routes
	r.Use(AuthMiddleware)

	return r
}
