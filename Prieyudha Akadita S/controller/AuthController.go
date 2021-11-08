package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// *AuthController inteerface is  a contracr what controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// *Disini masukan service yanng kalian butuhkan
	// *this is where put your service
}

// *NewAuthController creates a  new instance of AuthController
func NewAuthController() AuthController {

	return &authController{}

}

func (c *authController) Login(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello i'm login",
	})

}

func (c *authController) Register(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello i'm Register",
	})

}
