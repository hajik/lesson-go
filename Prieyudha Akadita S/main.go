package main

import (
	"lessongin/config"
	"lessongin/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {

	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login")
		authRoutes.POST("/register")
	}

	r.Run()

}
