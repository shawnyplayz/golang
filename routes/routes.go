package routes

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	//User routes
	api.GET("/users", controllers.GetUsers)
	api.POST("/signup", controllers.Signup)
	api.POST("/login", controllers.Login)
	api.PUT("/user/:id", controllers.UpdateUser)
	api.DELETE("/user/:id", controllers.DeleteUser)
	api.POST("/api/send-reset-link", controllers.SendResetLink)
	api.POST("/api/reset-password", controllers.ResetPassword)

	// Customer routes
	api.GET("/customers", controllers.GetCustomers)
	api.GET("/customers/:id", controllers.GetCustomerByID)
	api.POST("/customers", controllers.CreateCustomer)
	api.PUT("/customers/:id", controllers.UpdateCustomer)
	api.DELETE("/customers/:id", controllers.DeleteCustomer)

	// Vehicle routes (optional, depends on how tightly they're tied to customers)
	api.GET("/vehicles", controllers.GetVehicles)
	api.GET("/vehicles:id", controllers.GetVehicles)
	api.POST("/vehicles", controllers.CreateVehicle)
	api.PUT("/vehicles/:id", controllers.UpdateVehicle)
	api.DELETE("/vehicles/:id", controllers.DeleteVehicle)

}
