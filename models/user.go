package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique;not null" validate:"required"`
	Password     string `json:"password" gorm:"not null"`
	RoleAccessID uint   `json:"role_access_id"` // maps to Role.AccessID
	IsDev        bool   `json:"isDev" gorm:"default:false"`

	Role Role `gorm:"foreignKey:RoleAccessID;references:AccessID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
