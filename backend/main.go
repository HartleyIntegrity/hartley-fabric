package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Hartley-Fabric/backend/api"

	"github.com/HartleyIntegrity/hartley-fabric/backend/api"
	"github.com/HartleyIntegrity/hartley-fabric/backend/blockchain"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/signin", api.signInHandler).Methods("POST")

	genesisBlock := blockchain.Block{
		Index:        0,
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: []blockchain.Transaction{},
		PrevHash:     "",
		Hash:         "",
	}
	genesisBlock.Hash = blockchain.CreateHash(genesisBlock)
	api.Blockchain = append(api.Blockchain, genesisBlock)

	apiHandler := api.NewAPI()

	port := "8000"
	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
