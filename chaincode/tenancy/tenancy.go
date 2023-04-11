package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TenancyContract struct {
	contractapi.Contract
}

type Property struct {
	ID          string `json:"id"`
	OwnerID     string `json:"ownerID"`
	Address     string `json:"address"`
	RentalPrice int    `json:"rentalPrice"`
	Status      string `json:"status"`
}

type TenancyAgreement struct {
	ID               string `json:"id"`
	PropertyID       string `json:"propertyID"`
	TenantID         string `json:"tenantID"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	RentAmount       int    `json:"rentAmount"`
	SecurityDeposit  int    `json:"securityDeposit"`
	PaymentFrequency string `json:"paymentFrequency"`
}

// CreateProperty stores a new property on the ledger
func (t *TenancyContract) CreateProperty(ctx contractapi.TransactionContextInterface, propertyID string, ownerID string, address string, rentalPrice int) error {
	property := Property{
		ID:          propertyID,
		OwnerID:     ownerID,
		Address:     address,
		RentalPrice: rentalPrice,
		Status:      "available",
	}

	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(propertyID, propertyJSON)
}

// GetProperty returns a property by its ID
func (t *TenancyContract) GetProperty(ctx contractapi.TransactionContextInterface, propertyID string) (*Property, error) {
	propertyJSON, err := ctx.GetStub().GetState(propertyID)
	if err != nil {
		return nil, fmt.Errorf("failed to read property: %v", err)
	}
	if propertyJSON == nil {
		return nil, fmt.Errorf("property not found: %s", propertyID)
	}

	var property Property
	err = json.Unmarshal(propertyJSON, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

// CreateTenancyAgreement stores a new tenancy agreement on the ledger
func (t *TenancyContract) CreateTenancyAgreement(ctx contractapi.TransactionContextInterface, agreementID string, propertyID string, tenantID string, startDate string, endDate string, rentAmount int, securityDeposit int, paymentFrequency string) error {
	agreement := TenancyAgreement{
		ID:               agreementID,
		PropertyID:       propertyID,
		TenantID:         tenantID,
		StartDate:        startDate,
		EndDate:          endDate,
		RentAmount:       rentAmount,
		SecurityDeposit:  securityDeposit,
		PaymentFrequency: paymentFrequency,
	}

	agreementJSON, err := json.Marshal(agreement)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(agreementID, agreementJSON)
}

// GetTenancyAgreement returns a tenancy agreement by its ID
func (t *TenancyContract) GetTenancyAgreement(ctx contractapi.TransactionContextInterface, agreementID string) (*TenancyAgreement, error) {
	agreementJSON, err := ctx.GetStub().GetState(agreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to read tenancy agreement: %v", err)
	}
	if agreementJSON == nil {
		return nil, fmt.Errorf("tenancy agreement not found: %s", agreementID)
	}

	var agreement TenancyAgreement
	err = json.Unmarshal(agreementJSON, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}
