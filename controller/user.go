package controller

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/api"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(context *gin.Context) {
	// need to finish
	var userRegReq api.DouyinUserRegisterRequest
	var userRegResp api.DouyinUserRegisterResponse
	err := context.Bind(&userRegReq)
	if err != nil {
		context.JSON(http.StatusOK, userRegResp)
	}
	var user model.User
	user.AuthName = *userRegReq.Username
	user.Password = *userRegReq.Password
	//db := config.DB

}

func UserLoginHandler(context *gin.Context) {
}

func UserInfoHandler(context *gin.Context) {
}
