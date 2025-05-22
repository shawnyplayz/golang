package seeders

import (
	"gin/models"

	"gorm.io/gorm"
)

func SeedRoles(DB *gorm.DB) {
	roles := []models.Role{
		{AccessID: 5, AccessName:"SuperAdmin"},
		{AccessID: 4, AccessName: "Admin"},
		{AccessID: 3, AccessName: "User"},
	}

	for _, role := range roles {
		var existing models.Role
		if err := DB.Where("access = ?", role.AccessName).First(&existing).Error; err != nil {
			DB.Create(&role)
		}
	}
}
