package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hartley-fabric/hartley-fabric/backend/database"
)

// PropertiesHandler handles requests to get all properties.
func PropertiesHandler(db database.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		properties := db.GetProperties()
		json.NewEncoder(w).Encode(properties)
	}
}

// PropertyHandler handles requests to get a specific property by its ID.
func PropertyHandler(db database.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		property, err := db.GetProperty(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(property)
	}
}
