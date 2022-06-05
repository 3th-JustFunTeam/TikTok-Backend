package controller

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/service"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationActionHandler 关注
func RelationActionHandler(c *gin.Context) {

	token := c.Query("token")
	followingId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	claims, err := utils.ParseToken(token)
	id := strconv.Itoa(claims.UserId)
	if id == followingId {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "亲! 不能自己关注自己哦",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err,
		})
		return
	}
	// 关注
	if actionType == "1" {

		err := service.AddFollowing(followingId, claims.UserId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "不合法的操作",
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "关注成功",
			})
		}
		// 取关
	} else {
		err := service.RemoveFollowing(followingId, claims.UserId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "取关成功",
			})
		}
	}

}

// RelationActionHandler 关注列表
func FollowListHandler(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	_, err := utils.ParseToken(token)
	id, _ := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
		return
	}

	res, err := service.GetFollowingListByUserId(id, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "获取关注列表成功",
			"user_list":   res,
		})
	}

}

// RelationActionHandler 粉丝列表
func FollowerListHandler(c *gin.Context) {

	token := c.Query("token")
	userId := c.Query("user_id")
	_, err := utils.ParseToken(token)
	id, _ := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}

	res, err := service.GetFenshiListByUserId(id, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "获取粉丝列表成功",
			"user_list":   res,
		})
	}
}
