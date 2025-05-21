package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string    `json:"first_name" validate:"required,min=2"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone" gorm:"unique" validate:"required,min=2"`
	Email     string    `json:"email" gorm:"unique" validate:"required,min=2"`
	Vehicles   []Vehicle `json:"vehicles" gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
