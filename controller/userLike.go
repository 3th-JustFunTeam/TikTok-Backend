package controller

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/service"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteActionHandler 点赞
func FavoriteActionHandler(c *gin.Context) {

	token := c.Query("token")

	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err,
		})
	}

	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	videoId, _ := strconv.Atoi(video_id)
	actionType, _ := strconv.Atoi(action_type)
	// 点赞
	if actionType == 1 {

		err := service.Like(claims.UserId, videoId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "点赞成功",
			})
		}

	} else { // 取消点赞
		err := service.DownLike(claims.UserId, videoId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "取消点赞成功",
			})
		}
	}
}

// FavoriteListHandler 点赞列表
func FavoriteListHandler(c *gin.Context) {
	token := c.Query("token")

	userId := c.Query("user_id")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err,
		})
	}
	likeList := service.FindLikeList(userId, token, claims.UserId)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"video_list":  likeList,
	})

}
