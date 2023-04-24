package database

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// Property represents a property that can be rented.
type Property struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// Contract represents a smart tenancy agreement between a landlord and a tenant.
type Contract struct {
	ID           string `json:"id"`
	PropertyID   string `json:"property_id"`
	LandlordName string `json:"landlord_name"`
	TenantName   string `json:"tenant_name"`
	RentAmount   int    `json:"rent_amount"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

// Database is a thread-safe in-memory database.
type Database struct {
	properties map[string]*Property
	contracts  map[string]*Contract
	mutex      sync.RWMutex
}

// NewDatabase creates a new in-memory database instance.
func NewDatabase() *Database {
	return &Database{
		properties: make(map[string]*Property),
		contracts:  make(map[string]*Contract),
	}
}

// GetProperties returns a list of all properties.
func (d *Database) GetProperties() []*Property {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	properties := make([]*Property, 0, len(d.properties))
	for _, property := range d.properties {
		properties = append(properties, property)
	}
	return properties
}

// GetProperty returns the property with the specified ID.
func (d *Database) GetProperty(id string) (*Property, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	property, ok := d.properties[id]
	if !ok {
		return nil, errors.New("property not found")
	}
	return property, nil
}

// GetContracts returns a list of all contracts.
func (d *Database) GetContracts() []*Contract {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	contracts := make([]*Contract, 0, len(d.contracts))
	for _, contract := range d.contracts {
		contracts = append(contracts, contract)
	}
	return contracts
}

// CreateContract creates a new contract and returns its ID.
func (d *Database) CreateContract(contract *Contract) (string, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// check if the property exists
	if _, ok := d.properties[contract.PropertyID]; !ok {
		return "", errors.New("property not found")
	}

	// generate a new ID for the contract
	contract.ID = uuid.New().String()

	// add the contract to the database
	d.contracts[contract.ID] = contract
	return contract.ID, nil
}

// DeleteContract deletes the contract with the specified ID.
func (d *Database) DeleteContract(id string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.contracts[id]
	if !ok {
		return errors.New("contract not found")
	}

	delete(d.contracts, id)
	return nil
}
