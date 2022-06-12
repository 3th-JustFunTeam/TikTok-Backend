package controller

import (
	"net/http"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/common"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(ctx *gin.Context) {
	name := ctx.Query("username")
	pwd := ctx.Query("password")
	if len(name) > 32 || len(pwd) > 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "用户名或密码格式错误",
		})
	}
	user := po.User{
		AuthName:        name,
		Password:        pwd,
		NickName:        common.DEFAULT_NICK_NAME,
		Avatar:          common.DEFAULT_AVATAR,
		Signature:       common.DEFAULT_SIGNATURE,
		BackgroundImage: common.DEFAULT_BACKGROUND_IMAGE,
	}
	dao.DB.Where("auth_name = ?", name).Find(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "该用户名已被注册",
		})
	} else {
		dao.DB.Create(&user)
		dao.DB.Where("auth_name = ?", name).Find(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"StatusCode": 1,
				"StatusMsg":  "注册失败",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"StatusCode": 0,
				"StatusMsg":  "注册成功",
				"user_id":    user.ID,
			})
		}
	}
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
		followCount, followerCount := dao.QueryFollowingListByUserIdCount(int(user.ID))
		userInfo := vo.UserInfo{
			UserId:          uint(user.ID),
			NickName:        user.NickName,
			FollowCount:     followCount,
			FollowerCount:   followerCount,
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
