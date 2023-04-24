package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAndGetTenancyAgreement(t *testing.T) {
	// Create a new HTTP POST request with JSON body
	agreement := TenancyAgreement{
		Property:   "123 Main St",
		Landlord:   "John Smith",
		Tenant:     "Jane Doe",
		StartDate:  "2022-01-01",
		EndDate:    "2023-01-01",
		SigningKey: "secret",
	}
	payload, _ := json.Marshal(agreement)
	request := httptest.NewRequest(http.MethodPost, "/api/tenancy-agreements", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder and handle the request
	recorder := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(recorder, request)

	// Check the response status code and body
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response := recorder.Body.Bytes()
	if string(response) != "1" {
		t.Errorf("handler returned unexpected body: got %v want %v", string(response), "1")
	}

	// Create a new HTTP GET request to retrieve the agreement
	request = httptest.NewRequest(http.MethodGet, "/api/tenancy-agreements/1", nil)

	// Create a new HTTP recorder and handle the request
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Check the response status code and body
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = recorder.Body.Bytes()
	var retrievedAgreement TenancyAgreement
	json.Unmarshal(response, &retrievedAgreement)
	if retrievedAgreement.Property != "123 Main St" {
		t.Errorf("handler returned unexpected body: got %v want %v", retrievedAgreement.Property, "123 Main St")
	}
}

func TestUpdateAndDeleteTenancyAgreement(t *testing.T) {
	// Create a new HTTP PUT request with JSON body
	agreement := TenancyAgreement{
		Property:   "456 Oak St",
		Landlord:   "Bob Johnson",
		Tenant:     "Alice Brown",
		StartDate:  "2022-02-01",
		EndDate:    "2023-02-01",
		SigningKey: "secret",
	}
	payload, _ := json.Marshal(agreement)
	request := httptest.NewRequest(http.MethodPut, "/api/tenancy-agreements/1", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder and handle the request
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Check the response status code and body
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = recorder.Body.Bytes()
	var retrievedAgreement TenancyAgreement
	json.Unmarshal(response, &retrievedAgreement)
	if retrievedAgreement.Property != "456 Oak St" {
		t.Errorf("handler returned unexpected body: got %v want %v", retrievedAgreement.Property, "456 Oak St")
	}

	// Create a new HTTP DELETE request to delete the agreement
	request = httptest.NewRequest(http.MethodDelete, "/api/tenancy-agreements/1", nil)

	// Create a new HTTP recorder and handle the request
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Check the response status code and body
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = recorder.Body.Bytes()
	if string(response) != "1" {
		t.Errorf("handler returned unexpected body: got %v want %v", string(response), "1")
	}

	// Create a new HTTP GET request to retrieve the deleted agreement
	request = httptest.NewRequest(http.MethodGet, "/api/tenancy-agreements/1", nil)

	// Create a new HTTP recorder and handle the request
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Check the response status code and body
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetAllTenancyAgreements(t *testing.T) {
	// Create a new HTTP GET request to retrieve all agreements
	request := httptest.NewRequest(http.MethodGet, "/api/tenancy-agreements", nil)

	// Create a new HTTP recorder and handle the request
	recorder := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(recorder, request)

	// Check the response status code and body
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response := recorder.Body.Bytes()
	var agreements []TenancyAgreement
	json.Unmarshal(response, &agreements)
	if len(agreements) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", len(agreements), 0)
	}
}
