package controller

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/api"
	"github.com/gin-gonic/gin"
	"log"
)

func UserRegisterHandler(context *gin.Context) {
	var user api.DouyinUserRegisterRequest
	err := context.Bind(&user)
	if err != nil {
		return
	}
	log.Printf("user: %s\n", user.GetUsername())
}

func UserLoginHandler(context *gin.Context) {

}

func UserInfoHandler(context *gin.Context) {

}
