package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HartleyIntegrity/hartley-fabric/api/fabric"
	"github.com/gorilla/mux"
)

func GetProperties(w http.ResponseWriter, r *http.Request) {
	properties, err := fabric.GetAllProperties()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting properties: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(properties)
}

func GetProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	property, err := fabric.GetProperty(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting property: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(property)
}

// ...existing code...

func CreateProperty(w http.ResponseWriter, r *http.Request) {
	var property fabric.Property
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&property)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = fabric.CreateProperty(property)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating property: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var property fabric.Property
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&property)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = fabric.UpdateProperty(id, property)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating property: %v", err), http.StatusInternalServerError)
		return
	}
}

func TerminateTenancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := fabric.TerminateTenancy(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error terminating tenancy: %v", err), http.StatusInternalServerError)
		return
	}
}
