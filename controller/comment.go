package controller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/3th-JustFunTeam/Tiktok-Backend/service"
	"github.com/gin-gonic/gin"
)

// CommentActionHandler 添加评论
func CommentActionHandler(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	videoId, _ := strconv.Atoi(video_id)
	actionType, _ := strconv.Atoi(action_type)
	// 添加评论
	if actionType == 1 {
		commentText := c.Query("comment_text")
		commentText, _ = url.QueryUnescape(commentText)
		commentVo, err := service.AddComment(videoId, token, commentText)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "评论成功",
				"comment":     commentVo,
			})
		}

	} else { // 删除评论
		comment_id := c.Query("comment_id")
		commentId, _ := strconv.Atoi(comment_id)
		err := service.RemoveCommentById(commentId, videoId)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "删除评论成功",
			})
		}
	}
}

// CommentListHandler 视频列表
func CommentListHandler(c *gin.Context) {

	token := c.Query("token")
	video_id := c.Query("video_id")

	videoId, _ := strconv.Atoi(video_id)

	commentList, err := service.GetCommentList(videoId, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code":  0,
			"status_msg":   "获取评论列表成功",
			"comment_list": commentList,
		})
	}
}
