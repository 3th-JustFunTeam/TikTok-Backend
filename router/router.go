package router

import (
	controller "github.com/3th-JustFunTeam/Tiktok-Backend/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/douyin/")
	g.POST("/user/register/", controller.UserRegisterHandler)
	g.POST("/user/login/", controller.UserLoginHandler)
	g.GET("/user/", controller.UserInfoHandler)
	g.GET("/feed/", controller.VideoFeedHandler)
	g.POST("/publish/action/", controller.VideoPublishHandler)
	g.GET("/publish/list/", controller.VideoPublishListHandler)

}
