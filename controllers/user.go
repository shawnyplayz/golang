package controllers

import (
	"fmt"
	"gin/config"

	// "gin/helper"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user models.User
	if err:= c.ShouldBindJSON(&user); err !=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	return
	}
	// errors  := helper.ValidationHelper(&user)
	// if len(errors) > 0{
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": errors[0]})
	// 	return
	// }
	var userFound models.User
	config.DB.Where("email=?", user.Email).Find(&userFound);
	
	if userFound.ID !=0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"User with this email address already exists"})
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("It reached here!");

	user.Password = string(passwordHash)
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"User Registered Successfully!"})
}