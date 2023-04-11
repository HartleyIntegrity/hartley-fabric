// api/api.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

// TenancyAgreement represents a tenancy agreement on the blockchain
type TenancyAgreement struct {
	ID               string `json:"id"`
	PropertyID       string `json:"property_id"`
	TenantID         string `json:"tenant_id"`
	LandlordID       string `json:"landlord_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	RentAmount       int    `json:"rent_amount"`
	RentPaymentTerms string `json:"rent_payment_terms"`
}

var tenancyAgreements []TenancyAgreement

// GetTenancyAgreements retrieves all tenancy agreements
// GetTenancyAgreements retrieves all tenancy agreements
func GetTenancyAgreements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contract, err := connectToNetwork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := contract.EvaluateTransaction("getAllTenancyAgreements")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var agreements []TenancyAgreement
	json.Unmarshal(result, &agreements)

	json.NewEncoder(w).Encode(agreements)
}

// CreateTenancyAgreement creates a new tenancy agreement
func CreateTenancyAgreement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAgreement TenancyAgreement
	json.NewDecoder(r.Body).Decode(&newAgreement)

	contract, err := connectToNetwork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	args := []string{
		newAgreement.ID,
		newAgreement.PropertyID,
		newAgreement.TenantID,
		newAgreement.LandlordID,
		newAgreement.StartDate,
		newAgreement.EndDate,
		strconv.Itoa(newAgreement.RentAmount),
		newAgreement.RentPaymentTerms,
	}

	_, err = contract.SubmitTransaction("createTenancyAgreement", args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newAgreement)
}

func connectToNetwork() (*gateway.Contract, error) {
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile("connection-profile.yaml")),
		gateway.WithIdentity(wallet, "user1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		return nil, fmt.Errorf("failed to get network: %v", err)
	}

	contract := network.GetContract("tenancy_agreement")

	return contract, nil
}

func main() {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/tenancy-agreements", GetTenancyAgreements).Methods("GET")
	router.HandleFunc("/tenancy-agreements", CreateTenancyAgreement).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
