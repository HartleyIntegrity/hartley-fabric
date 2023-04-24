package api

import (
	"net/http"

	"github.com/HartleyIntegrity/hartley-fabric/backend/blockchain"
)

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	chain := blockchain.GetBlockchain()
	sendJSON(w, http.StatusOK, chain)
}

func AddBlock(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the function to add a block
}

func UpdateBlock(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the function to update a block
}

func DeleteBlock(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the function to delete a block
}
