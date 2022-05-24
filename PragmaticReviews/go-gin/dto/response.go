package dto

type Repsonse struct {
	Message string `json:"message"`
}

func (api *VideoApi) Authentication(ctx *gin.Context {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token:token,
		})
	}else{
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Not Authorized",
		})
	}
})

func (api *VideoApi) GetVideo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.videoController.FindAll())
} 