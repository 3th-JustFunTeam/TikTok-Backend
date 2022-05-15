package router

import (
	controller2 "github.com/3th-JustFunTeam/Tiktok-Backend/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/douyin/user/register/", controller2.UserRegisterHandler)
	r.POST("/douyin/user/login/", controller2.UserLoginHandler)
	r.GET("/douyin/user/", controller2.UserInfoHandler)
	r.GET("/douyin/feed/", controller2.VideoFeedHandler)
	r.POST("/douyin/publish/action/", controller2.VideoPublishHandler)
	r.GET("/douyin/publish/list/", controller2.VideoPublishListHandler)
}
