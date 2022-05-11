package server

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/server/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/douyin/user/register/", controller.UserRegisterHandler)
	r.POST("/douyin/user/login/", controller.UserLoginHandler)
	r.GET("/douyin/user/", controller.UserInfoHandler)
	r.GET("/douyin/feed/", controller.VideoFeedHandler)
	r.POST("/douyin/publish/action/", controller.VideoPublishHandler)
	r.GET("/douyin/publish/list/", controller.VideoPublishListHandler)
}
