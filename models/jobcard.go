package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentType string

const (
	Cash         PaymentType = "Cash"
	KNet         PaymentType = "K Net"
	LinkTransfer PaymentType = "Link Transfer"
)

type JobCard struct {
	gorm.Model
	WIPNumber      uint        `json:"wip_number" gorm:"unique"`
	VehicleID      uint        `json:"vehicle_id"`
	Vehicle        Vehicle     `json:"vehicle" gorm:"foreignKey:VehicleID;references:ID"`
	CustomerID     uint        `json:"customer_id"`
	Customer       Customer    `json:"customer" gorm:"foreignKey:CustomerID;references:ID"`
	ServiceType    string      `json:"service_type"`
	RepairDate     time.Time   `json:"repair_date"`
	TimeIn         string      `json:"time_in"`
	Department     string      `json:"department"`
	Team           string      `json:"team"`
	Branch         string      `json:"branch"`
	InvoiceNumber  string      `json:"invoice_number"`
	ServiceAdvisor string      `json:"service_advisor"`
	Remarks        string      `json:"remarks"`
	Amount         uint        `json:"amount" gorm:"not null"`
	PaymentType    PaymentType `json:"payment_type" gorm:"type:varchar(20);not null"`
	HasPayed       bool        `json:"has_payed"`
	UserID         uint        `json:"user_id"`
	User           User        `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
