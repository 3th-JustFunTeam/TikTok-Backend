package controller

import (
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/3th-JustFunTeam/Tiktok-Backend/service"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"github.com/gin-gonic/gin"
)

// VideoFeedHandler 视频流接口
func VideoFeedHandler(context *gin.Context) {
	token := context.Query("token")
	video, err := service.GetAllVideo(token)

	if err != nil {
		context.JSON(404, gin.H{
			"msg": err,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "success",
			"next_time":   123123124,
			"video_list":  video,
		})
	}
}

// VideoPublishHandler 投稿接口
func VideoPublishHandler(context *gin.Context) {
	token := context.PostForm("token")
	title := context.PostForm("title")
	file, err := context.FormFile("data")

	// 解析token
	claims, err2 := utils.ParseToken(token)

	if err2 != nil {
		context.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "token error",
		})
	}
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}
	fileName := filepath.Base(file.Filename)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000) + 1
	newFileName := strconv.Itoa(r) + fileName
	saveFile := filepath.Join("./public/", newFileName)
	if err := context.SaveUploadedFile(file, saveFile); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}

	err = service.AddVideo(claims.UserId, title, newFileName)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"StatusCode": 0,
			"StatusMsg":  "发布成功",
		})
	}

}

// VideoPublishListHandler 发布列表
func VideoPublishListHandler(c *gin.Context) {

	token := c.Query("token")
	user_id := c.Query("user_id")

	videoList, err := service.GetUserAllVideo(user_id, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "success",
			"video_list":  videoList,
		})

	}
}
