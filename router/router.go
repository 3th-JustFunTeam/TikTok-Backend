package router

import (
	controller "github.com/3th-JustFunTeam/Tiktok-Backend/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/douyin/")
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	// 基础接口
	g.POST("/user/register/", controller.UserRegisterHandler)
	g.POST("/user/login/", controller.UserLoginHandler)
	g.GET("/user/", controller.UserInfoHandler)
	g.GET("/feed/", controller.VideoFeedHandler)
	g.POST("/publish/action/", controller.VideoPublishHandler)
	g.GET("/publish/list/", controller.VideoPublishListHandler)

	// 拓展接口 - I
	//g.POST("/favorite/action/", controller.FavoriteActionHandler)
	g.GET("/favorite/list/", controller.FavoriteListHandler)
	g.POST("/comment/action/", controller.CommentActionHandler)
	g.GET("/comment/list/", controller.CommentListHandler)

	// 拓展接口 - II
	g.POST("/relation/action/", controller.RelationActionHandler)
	g.GET("/relation/follow/list/", controller.FollowListHandler)
	g.GET("/relation/follower/list/", controller.FollowerListHandler)

}
