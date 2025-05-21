package controllers

import (
	"gin/config"
	"gin/helper"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	if err:=config.DB.Preload("Vehicles").Find(&customers).Error; err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.Preload("Vehicles").First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errors  := helper.ValidationHelper(&customer)
	if len(errors) > 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": errors[0]})
		return
	}
	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	var updatedCustomer models.Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errors :=helper.ValidationHelper(&updatedCustomer)
	if len(errors) > 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": errors[0]})
		return
	}

	if err := config.DB.Model(&customer).Updates(models.Customer{
		FirstName:        updatedCustomer.FirstName,
		LastName:        updatedCustomer.LastName,
		Phone: 			updatedCustomer.Phone,
		Email:         updatedCustomer.Email,
		Vehicles : updatedCustomer.Vehicles,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	config.DB.Unscoped().Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
