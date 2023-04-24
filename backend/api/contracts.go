package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hartley-fabric/hartley-fabric/backend/database"
)

// ContractsHandler handles requests to get all contracts and create a new contract.
func ContractsHandler(db database.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			contracts := db.GetContracts()
			json.NewEncoder(w).Encode(contracts)
		case "POST":
			var contract database.Contract
			err := json.NewDecoder(r.Body).Decode(&contract)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// create a new contract and add it to the database
			contract.ID = uuid.New().String()
			_, err = db.CreateContract(&contract)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(contract)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// ContractHandler handles requests to delete a contract by its ID.
func ContractHandler(db database.Models) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		err := db.DeleteContract(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
