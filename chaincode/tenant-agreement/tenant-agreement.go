package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TenancyAgreement struct {
	ID                  string  `json:"id"`
	StartDate           string  `json:"startDate"`
	EndDate             string  `json:"endDate"`
	RentAmount          float64 `json:"rentAmount"`
	SecurityDeposit     float64 `json:"securityDeposit"`
	SecurityDepositPaid bool    `json:"securityDepositPaid"`
	Landlord            string  `json:"landlord"`
	Tenant              string  `json:"tenant"`
}

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) CreateTenancyAgreement(ctx contractapi.TransactionContextInterface, id string, startDate string, endDate string, rentAmount float64, securityDeposit float64, landlord string, tenant string) error {
	tenancyAgreement := TenancyAgreement{
		ID:                  id,
		StartDate:           startDate,
		EndDate:             endDate,
		RentAmount:          rentAmount,
		SecurityDeposit:     securityDeposit,
		SecurityDepositPaid: false,
		Landlord:            landlord,
		Tenant:              tenant,
	}

	tenancyAgreementJSON, err := json.Marshal(tenancyAgreement)
	if err != nil {
		return fmt.Errorf("failed to marshal tenancy agreement JSON: %v", err)
	}

	err = ctx.GetStub().PutState(id, tenancyAgreementJSON)
	if err != nil {
		return fmt.Errorf("failed to put tenancy agreement JSON on ledger: %v", err)
	}

	return nil
}

func (s *SmartContract) PaySecurityDeposit(ctx contractapi.TransactionContextInterface, id string) error {
	tenancyAgreementJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read tenancy agreement from ledger: %v", err)
	}
	if tenancyAgreementJSON == nil {
		return fmt.Errorf("tenancy agreement does not exist")
	}

	var tenancyAgreement TenancyAgreement
	err = json.Unmarshal(tenancyAgreementJSON, &tenancyAgreement)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tenancy agreement JSON: %v", err)
	}

	if tenancyAgreement.SecurityDepositPaid {
		return fmt.Errorf("security deposit has already been paid")
	}

	err = ctx.GetStub().SetEvent("SecurityDepositPaid", []byte(id))
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	tenancyAgreement.SecurityDepositPaid = true

	tenancyAgreementJSON, err = json.Marshal(tenancyAgreement)
	if err != nil {
		return fmt.Errorf("failed to marshal tenancy agreement JSON: %v", err)
	}

	err = ctx.GetStub().PutState(id, tenancyAgreementJSON)
	if err != nil {
		return fmt.Errorf("failed to put tenancy agreement JSON on ledger: %v", err)
	}

	return nil
}

func (s *SmartContract) RefundSecurityDeposit(ctx contractapi.TransactionContextInterface, id string) error {
	tenancyAgreementJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read tenancy agreement from ledger: %v", err)
	}
	if tenancyAgreementJSON == nil {
		return fmt.Errorf("tenancy agreement does not exist")
	}

	var tenancyAgreement TenancyAgreement
	err = json.Unmarshal(tenancyAgreementJSON, &tenancyAgreement)
	if err != nil {
		return fmt.Errorf("failed to unmarshal tenancy agreement JSON: %v", err)
	}

	if !tenancyAgreement.SecurityDepositPaid {
		return fmt.Errorf("security deposit has not been paid")
	}

	err = ctx.GetStub().SetEvent("SecurityDepositRefunded", []byte(id))
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	tenancyAgreement.SecurityDepositPaid = false

	tenancyAgreementJSON, err = json.Marshal(tenancyAgreement)
	if err != nil {
		return fmt.Errorf("failed to marshal tenancy agreement JSON: %v", err)
	}

	err = ctx.GetStub().PutState(id, tenancyAgreementJSON)
	if err != nil {
		return fmt.Errorf("failed to put tenancy agreement JSON on ledger: %v", err)
	}

	return nil
}
