package controller

import (
	"net/http"
	"strconv"

	"github.com/3th-JustFunTeam/Tiktok-Backend/service"
	"github.com/gin-gonic/gin"
)

func FavoriteActionHandler(ctx *gin.Context) {
	// 不支持未登录用户
	_, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "用户未登录",
		})
		return
	}

	token := ctx.Query("token")
	video_id, _ := strconv.Atoi(ctx.Query("video_id"))
	action_type, _ := strconv.Atoi(ctx.Query("action_type"))
	if action_type == 1 {
		err := service.Like(token, video_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"StatusCode": 1,
				"StatusMsg":  err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"StatusCode": 0,
				"StatusMsg":  "点赞成功",
			})
		}
	} else if action_type == 2 {
		err := service.DisLike(video_id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"StatusCode": 1,
				"StatusMsg":  err,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"StatusCode": 0,
				"StatusMsg":  "取消点赞成功",
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "action_type数据错误",
		})
	}
}
