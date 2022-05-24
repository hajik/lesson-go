package main

import (
	"go/gin/controllers"
	"go/gin/middleware"
	"go/gin/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService        = service.New()
	videoController controllers.VideoController = controllers.New(videoService)

	// * login
	// loginService    service.LoginService        = service.NewLoginService()
	// loginController controllers.LoginController = controllers.New(loginService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())
	server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())

	// login endpoint: Authentication + token Creation
	// server.POST("/login", func(ctx *gin.Context) {
	// 	token := loginController.Login(ctx)
	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })

	apiRoutes := server.Group("/api", middleware.BasicAuth())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			// ctx.JSON(http.StatusOK, videoController.Save(c))
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Videos is valid",
				})
			}

		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	err := godotenv.Load()
	if err != nil {
		panic("Failed load a file .env")
	}
	port := os.Getenv("PORT")
	server.Run(":" + port)

}
