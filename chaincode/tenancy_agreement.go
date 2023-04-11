// chaincode/tenancy_agreement.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

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

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func (s *SmartContract) CreateTenancyAgreement(ctx contractapi.TransactionContextInterface, id string, propertyID string, tenantID string, landlordID string, startDate string, endDate string, rentAmount int, rentPaymentTerms string) error {
	exists, err := s.TenancyAgreementExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the tenancy agreement %s already exists", id)
	}

	agreement := TenancyAgreement{
		ID:               id,
		PropertyID:       propertyID,
		TenantID:         tenantID,
		LandlordID:       landlordID,
		StartDate:        startDate,
		EndDate:          endDate,
		RentAmount:       rentAmount,
		RentPaymentTerms: rentPaymentTerms,
	}
	agreementJSON, err := json.Marshal(agreement)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, agreementJSON)
}

func (s *SmartContract) TenancyAgreementExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	agreementJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreementJSON != nil, nil
}

func (s *SmartContract) GetAllTenancyAgreements(ctx contractapi.TransactionContextInterface) ([]*TenancyAgreement, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var agreements []*TenancyAgreement
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err

		}

		var agreement TenancyAgreement
		err = json.Unmarshal(queryResponse.Value, &agreement)
		if err != nil {
			return nil, err
		}
		agreements = append(agreements, &agreement)
	}

	return agreements, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error create tenancy_agreement chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting tenancy_agreement chaincode: %s", err.Error())
	}
}
