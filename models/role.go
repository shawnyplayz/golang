package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	AccessID   uint   `json:"access_id" gorm:"uniqueIndex;not null"`
	AccessName string `json:"access_name" gorm:"unique;not null" validate:"required"`
}
