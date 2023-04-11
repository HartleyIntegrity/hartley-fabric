// chaincode/tenancy/tenancy.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type TenancyContract struct {
}

type Agreement struct {
	ID            string `json:"id"`
	PropertyID    string `json:"property_id"`
	LandlordID    string `json:"landlord_id"`
	TenantID      string `json:"tenant_id"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	RentAmount    int    `json:"rent_amount"`
	PaymentPeriod string `json:"payment_period"`
}

func (t *TenancyContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *TenancyContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "createAgreement" {
		return t.createAgreement(stub, args)
	} else if function == "queryAgreement" {
		return t.queryAgreement(stub, args)
	} else if function == "queryAllAgreements" {
		return t.queryAllAgreements(stub)
	} else if function == "updateAgreement" {
		return t.updateAgreement(stub, args)
	} else if function == "deleteAgreement" {
		return t.deleteAgreement(stub, args)
	}

	return shim.Error("Invalid function name.")
}

func (t *TenancyContract) createAgreement(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	agreement := Agreement{
		ID:            args[0],
		PropertyID:    args[1],
		LandlordID:    args[2],
		TenantID:      args[3],
		StartDate:     args[4],
		EndDate:       args[5],
		RentAmount:    strToInt(args[6]),
		PaymentPeriod: args[7],
	}

	agreementAsBytes, _ := json.Marshal(agreement)
	err := stub.PutState(agreement.ID, agreementAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create agreement: %s", args[0]))
	}

	return shim.Success(nil)
}

func (t *TenancyContract) queryAgreement(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	agreementAsBytes, _ := stub.GetState(args[0])
	if agreementAsBytes == nil {
		return shim.Error(fmt.Sprintf("Agreement not found: %s", args[0]))
	}

	return shim.Success(agreementAsBytes)
}

func (t *TenancyContract) queryAllAgreements(stub shim.ChaincodeStubInterface) peer.Response {
	startKey := ""
	endKey := ""
	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error("Failed to get agreements.")
	}
	defer resultsIterator.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Failed to iterate results.")
		}
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (t *TenancyContract) updateAgreement(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	agreementAsBytes, err := stub.GetState(args[0])
	if err != nil || agreementAsBytes == nil {
		return shim.Error(fmt.Sprintf("Agreement not found: %s", args[0]))
	}

	updatedAgreement := Agreement{
		ID:            args[0],
		PropertyID:    args[1],
		LandlordID:    args[2],
		TenantID:      args[3],
		StartDate:     args[4],
		EndDate:       args[5],
		RentAmount:    strToInt(args[6]),
		PaymentPeriod: args[7],
	}

	updatedAgreementAsBytes, _ := json.Marshal(updatedAgreement)
	err = stub.PutState(updatedAgreement.ID, updatedAgreementAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update agreement: %s", args[0]))
	}

	return shim.Success(nil)
}

func (t *TenancyContract) deleteAgreement(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to delete agreement: %s", args[0]))
	}

	return shim.Success(nil)
}

func strToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}

func main() {
	err := shim.Start(new(TenancyContract))
	if err != nil {
		fmt.Printf("Error starting TenancyContract chaincode: %s", err)
	}
}
