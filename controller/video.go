package controller

import (
	"errors"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/service"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"github.com/gin-gonic/gin"
)

// VideoFeedHandler 视频流接口
func VideoFeedHandler(context *gin.Context) {
	video, err := service.GetAllVideo()

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

	err = service.AddVideo(token, title, newFileName)

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
func VideoPublishListHandler(ctx *gin.Context) {
	/*
		思路一：我的逻辑是先判断user_id，如果user_id是否存在，再判断token是否已解析，全部完成可得到发布列表
		思路二*：先判断token是否已解析，通过接收到的错误，再查找id，全部完成可得到发布列表
	*/
	token := ctx.Query("token")
	user_id := ctx.Query("user_id")

	claims, err := utils.ParseToken(token)
	if err != nil || claims.Id != user_id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": 1,
			"StatusMsg":  errors.New("信息错误"),
		})
	} else {
		var user po.User
		dao.DB.Where("id = ?", claims.UserId).Find(&user)

		userInfo, err := service.GetUserInfoById(uint(claims.UserId))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"StatusCode": 1,
				"StatusMsg":  err.Error(),
			})
		} else {
			videoList, err := service.GetUserAllVideo(uint(claims.UserId))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"StatusCode": 1,
					"StatusMsg":  err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"StatusCode": 0,
					"StatusMsg":  err.Error(),
					"userInfo":   userInfo,
					"videoList":  videoList,
				})
			}
		}

	}
}
