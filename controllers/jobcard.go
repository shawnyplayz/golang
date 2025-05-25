package controllers

import (
	"gin/config"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJobCards(c *gin.Context) {
	var jobCard []models.JobCard
	if err := config.DB.Preload("Vehicle").Preload("Customer").Preload("User.Role").Find(&jobCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, jobCard)
}
func CreateJobCards(c *gin.Context) {
	var jobCard models.JobCard
	if err := c.ShouldBindJSON(&jobCard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var vehicle models.Vehicle
	if err := config.DB.Where("id = ? AND customer_id = ?", jobCard.VehicleID, jobCard.CustomerID).First(&vehicle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vehicle does not belong to the specified customer"})
		return
	}

	if err := config.DB.Create(&jobCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully added a job card!"})
}
func UpdateJobCard(c *gin.Context) {
	id := c.Param("id")
	var jobCard models.JobCard
	if err := config.DB.First(&jobCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job Card not found"})
		return
	}
	var updatedJobCard models.JobCard
	if err := c.ShouldBindJSON(&updatedJobCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var vehicle models.Vehicle
	if err := config.DB.Where("id = ? AND customer_id = ?", updatedJobCard.VehicleID, updatedJobCard.CustomerID).First(&vehicle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vehicle does not belong to the specified customer"})
		return
	}

	if err := config.DB.Model(&jobCard).Updates(models.JobCard{
		WIPNumber:      updatedJobCard.WIPNumber,
		VehicleID:      updatedJobCard.VehicleID,
		Vehicle:        updatedJobCard.Vehicle,
		CustomerID:     updatedJobCard.CustomerID,
		Customer:       updatedJobCard.Customer,
		ServiceType:    updatedJobCard.ServiceType,
		RepairDate:     updatedJobCard.RepairDate,
		TimeIn:         updatedJobCard.TimeIn,
		Department:     updatedJobCard.Department,
		Team:           updatedJobCard.Team,
		Branch:         updatedJobCard.Branch,
		InvoiceNumber:  updatedJobCard.InvoiceNumber,
		ServiceAdvisor: updatedJobCard.ServiceAdvisor,
		Remarks:        updatedJobCard.Remarks,
		Amount:         updatedJobCard.Amount,
		PaymentType:    updatedJobCard.PaymentType,
		HasPayed:       updatedJobCard.HasPayed,
		UserID:         updatedJobCard.UserID,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobCard)
}

func DeleteJobCard(c *gin.Context) {
	id := c.Param("id")
	var jobCard models.JobCard
	if err := config.DB.First(&jobCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job Card not found"})
		return
	}
	if err := config.DB.Unscoped().Delete(&jobCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Job Card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job Card Deleted Successfully!"})
}
