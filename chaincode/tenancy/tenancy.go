// chaincode/tenancy/tenancy.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TenancyContract struct {
	contractapi.Contract
}

type TenancyAgreement struct {
	ID         string `json:"id"`
	Landlord   string `json:"landlord"`
	Tenant     string `json:"tenant"`
	Property   string `json:"property"`
	RentAmount int    `json:"rentAmount"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

func (t *TenancyContract) CreateTenancy(ctx contractapi.TransactionContextInterface, id string, landlord string, tenant string, property string, rentAmount int, startDate string, endDate string) error {
	tenancy := TenancyAgreement{
		ID:         id,
		Landlord:   landlord,
		Tenant:     tenant,
		Property:   property,
		RentAmount: rentAmount,
		StartDate:  startDate,
		EndDate:    endDate,
	}

	tenancyBytes, err := json.Marshal(tenancy)
	if err != nil {
		return fmt.Errorf("failed to marshal tenancy: %v", err)
	}

	err = ctx.GetStub().PutState(id, tenancyBytes)
	if err != nil {
		return fmt.Errorf("failed to put tenancy to state: %v", err)
	}

	return nil
}

func (t *TenancyContract) GetTenancy(ctx contractapi.TransactionContextInterface, id string) (*TenancyAgreement, error) {
	tenancyBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenancy from state: %v", err)
	}

	if tenancyBytes == nil {
		return nil, fmt.Errorf("tenancy not found: %s", id)
	}

	var tenancy TenancyAgreement
	err = json.Unmarshal(tenancyBytes, &tenancy)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tenancy: %v", err)
	}

	return &tenancy, nil
}
