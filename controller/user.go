package controller

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/api"
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(context *gin.Context) {
	// need to finish
	var userRegReq api.DouyinUserRegisterRequest
	var userRegResp api.DouyinUserRegisterResponse

	err := context.Bind(&userRegReq)
	if err != nil {
		context.JSON(http.StatusOK, userRegResp)
	}
	var user po.User
	user.AuthName = *userRegReq.Username
	user.Password = *userRegReq.Password
	//db := config.DB

}

func UserLoginHandler(c *gin.Context) {

	name := c.Query("username")
	pwd := c.Query("password")

	user := po.User{
		AuthName: name,
		Password: pwd,
	}

	dao.DB.Where("auth_name = ? and password = ?", name, pwd).Find(&user)

	if user.ID != 0 {
		token, err := utils.GenerateToken(int(user.ID))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"StatusCode": 1,
				"StatusMsg":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "登录成功",
				"user_id":     user.ID,
				"token":       token,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "登录错误",
		})
	}

}

func UserInfoHandler(c *gin.Context) {

	token := c.Query("token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
	} else {

		var user po.User
		dao.DB.Where("id = ?", claims.UserId).Find(&user)

		userInfo := vo.UserInfo{
			UserId:          uint(user.ID),
			NickName:        user.NickName,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        false,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			BackgroundImage: user.BackgroundImage,
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "成功",
			"user":        userInfo,
		})

	}

}
