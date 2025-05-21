package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	VehicleName        string `json:"vehicle_name" validate:"required"`
	PlateNumber        string `json:"plate_number" gorm:"unique"`
	VehicleDescription string `json:"vehicle_description"`
	CustomerID         uint   `json:"customer_id"` // must be uint, not Customer
}
