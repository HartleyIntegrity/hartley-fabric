package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HartleyIntegrity/hartley-fabric/api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/properties", handlers.GetProperties).Methods("GET")
	router.HandleFunc("/properties/{id}", handlers.GetProperty).Methods("GET")
	router.HandleFunc("/properties", handlers.CreateProperty).Methods("POST")
	router.HandleFunc("/properties/{id}", handlers.UpdateProperty).Methods("PUT")
	router.HandleFunc("/properties/{id}/terminate", handlers.TerminateTenancy).Methods("PUT")

	fmt.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
