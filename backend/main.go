package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create a new router
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// define the routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/api/properties", getPropertiesHandler).Methods("GET")
	r.HandleFunc("/api/properties/{id}", getPropertyHandler).Methods("GET")
	r.HandleFunc("/api/contracts", getContractsHandler).Methods("GET")
	r.HandleFunc("/api/contracts", createContractHandler).Methods("POST")
	r.HandleFunc("/api/contracts/{id}", deleteContractHandler).Methods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// return a simple message
	w.Write([]byte("Welcome to Hartley-Fabric"))
}

func getPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this handler
}

func getPropertyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this handler
}

func getContractsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this handler
}

func createContractHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this handler
}

func deleteContractHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this handler
}
