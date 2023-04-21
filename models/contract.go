package models

import "github.com/jinzhu/gorm"

type Contract struct {
	gorm.Model
	Address         string `json:"address"`
	LandlordEmail   string `json:"landlord_email"`
	TenantEmail     string `json:"tenant_email"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	RentPerMonth    uint   `json:"rent_per_month"`
	Deposit         uint   `json:"deposit"`
	IsTerminated    bool   `json:"is_terminated"`
	TerminationDate string `json:"termination_date"`
}
