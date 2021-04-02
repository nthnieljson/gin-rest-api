package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nthnieljson/gin-rest-api/config"
	"github.com/nthnieljson/gin-rest-api/controller"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)
	

func main() {
	port := "8080"

	defer config.CloseDatabaseConnection(db)
	router := gin.Default()
	
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login")
		authRoutes.POST("/register")
	}

	router.Run(":" + port) 
}