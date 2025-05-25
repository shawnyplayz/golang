package config

import (
	"fmt"
	"gin/models"
	"gin/seeders"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	connectionString := "root:root@tcp(127.0.0.1:3306)/autoservice?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	fmt.Println("Connected to MySQL database")

	// Auto-migrate your models
	err = DB.AutoMigrate(&models.Customer{}, &models.Vehicle{}, &models.Role{}, &models.User{}, models.JobCard{})
	//Seeding the needful
	seeders.SeedRoles(DB)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
