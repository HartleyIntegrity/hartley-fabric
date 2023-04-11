package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
)

func TestCreateTenancyAgreement(t *testing.T) {
	cc := new(SmartContract)
	ctx := contractapi.NewTransactionContext(nil, nil)
	err := cc.CreateTenancyAgreement(ctx, "1", "2023-05-01", "2024-04-30", 1000.0, 1000.0, "Landlord", "Tenant")
	assert.NoError(t, err)

	tenancyAgreementJSON, err := ctx.GetStub().GetState("1")
	assert.NoError(t, err)

	var tenancyAgreement TenancyAgreement
	err = json.Unmarshal(tenancyAgreementJSON, &tenancyAgreement)
	assert.NoError(t, err)

	assert.Equal(t, "1", tenancyAgreement.ID)
	assert.Equal(t, "2023-05-01", tenancyAgreement.StartDate)
	assert.Equal(t, "2024-04-30", tenancyAgreement.EndDate)
	assert.Equal(t, 1000.0, tenancyAgreement.RentAmount)
	assert.Equal(t, 1000.0, tenancyAgreement.SecurityDeposit)
	assert.False(t, tenancyAgreement.SecurityDepositPaid)
	assert.Equal(t, "Landlord", tenancyAgreement.Landlord)
	assert.Equal(t, "Tenant", tenancyAgreement.Tenant)
}
