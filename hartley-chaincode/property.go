package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type Property struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tenancies   []Tenancy `json:"tenancies"`
	// ...other fields...
}

type Tenancy struct {
	ID          string `json:"id"`
	TenantID    string `json:"tenantId"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Terminated  bool   `json:"terminated"`
	// ...other fields...
}

type Property struct {
	ID          string    `json:"id"`
	Title       string   


type PropertyChaincode struct {
}

func (t *PropertyChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *PropertyChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "createProperty":
		return t.createProperty(stub, args)
	case "updateProperty":
		return t.updateProperty(stub, args)
	case "terminateTenancy":
		return t.terminateTenancy(stub, args)
	// ...Add more functions as needed...
	default:
		return shim.Error(fmt.Sprintf("Unsupported function: %s", function))
	}
}

func (t *PropertyChaincode) createProperty(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 (JSON object of the property).")
	}

	var property Property
	err := json.Unmarshal([]byte(args[0]), &property)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to parse JSON: %s", err.Error()))
	}

	err = stub.PutState(property.ID, []byte(args[0]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create property: %s", err.Error()))
	}

	return shim.Success(nil)
}

func (t *PropertyChaincode) updateProperty(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 (ID and JSON object of the property).")
	}

	propertyID := args[0]
	var newProperty Property
	err := json.Unmarshal([]byte(args[1]), &newProperty)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to parse JSON: %s", err.Error()))
	}

	oldPropertyBytes, err := stub.GetState(propertyID)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get property: %s", err.Error()))
	}

	var oldProperty Property
	err = json.Unmarshal(oldPropertyBytes, &oldProperty)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to parse JSON: %s", err.Error()))
	}

	oldProperty.Title = newProperty.Title
	oldProperty.Description = newProperty.Description
	// ...Update other fields...

	newPropertyBytes, err := json.Marshal(oldProperty)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal JSON: %s", err.Error()))
	}

	err = stub.PutState(propertyID, newPropertyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update property: %s", err.Error()))
	}

	return shim.Success(nil)
}

func (t *PropertyChaincode) terminateTenancy(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 (property ID and tenancy ID).")
	}

	propertyID := args[0]
	tenancyID := args[1]

	propertyBytes, err := stub.GetState(propertyID)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get property: %s", err.Error()))
	}

	var property Property
	err = json.Unmarshal(propertyBytes, &property)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to parse JSON: %s", err.Error()))
	}

	tenancyFound := false
	for i, tenancy := range property.Tenancies {
		if tenancy.ID == tenancyID {
			property.Tenancies[i].Terminated = true
			tenancyFound = true
			break
		}
	}

	if !tenancyFound {
		return shim.Error(fmt.Sprintf("Tenancy with ID %s not found in property %s", tenancyID, propertyID))
	}

	updatedPropertyBytes, err := json.Marshal(property)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal JSON: %s", err.Error()))
	}

	err = stub.PutState(propertyID, updatedPropertyBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update property: %s", err.Error()))
	}

	return shim.Success(nil)
}


func main() {
	err := shim.Start(new(PropertyChaincode))
	if err != nil {
		fmt.Printf("Error starting Property chaincode: %s", err)
	}
}
