package controllers

import (
	"gin/config"
	"gin/helper"
	"gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	
	if err := config.DB.Find(&vehicles).Error; err != nil {
		log.Printf("DB error==>: %v",err);
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}
func GetVehicleByID(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}
func CreateVehicle(c *gin.Context){
	var vehicle models.Vehicle;
	if err:=c.ShouldBind(&vehicle); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
	}
	errors  := helper.ValidationHelper(&vehicle)
	if len(errors) > 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": errors[0]})
		return
	}
	if err := config.DB.Create(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func UpdateVehicle(c *gin.Context) {
	id := c.Param("id")

	var existingVehicle models.Vehicle
	if err := config.DB.First(&existingVehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	var updateData models.Vehicle
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errors :=helper.ValidationHelper(&updateData)
	if len(errors) > 0{
		c.JSON(http.StatusBadRequest, gin.H{"error": errors[0]})
		return
	}
	
	if err := config.DB.Model(&existingVehicle).Updates(models.Vehicle{
		VehicleName:        updateData.VehicleName,
		PlateNumber:        updateData.PlateNumber,
		VehicleDescription: updateData.VehicleDescription,
		CustomerID:         updateData.CustomerID,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, existingVehicle)
}


func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	config.DB.Unscoped().Delete(&vehicle)
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted"})
}