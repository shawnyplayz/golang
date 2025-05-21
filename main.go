package main

import (
	"gin/config"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDatabase()
    r := gin.Default()
    routes.RegisterRoutes(r);
    r.Run(":8080")
}
