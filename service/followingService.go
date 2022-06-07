package service

import (
	"errors"
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"strconv"
	"time"
)

// AddFollowing 关注
func AddFollowing(followingId string, userId int) error {

	now := time.Now()
	id, _ := strconv.Atoi(followingId)
	count := dao.QueryIsFollow(userId, id)
	if count > 0 {
		return errors.New("已经关注过了")
	}

	follow := po.Follow{
		UserId:      uint(userId),
		FollowingId: uint(id),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := dao.CreateFollowing(follow)

	return err

}

// RemoveFollowing 取关
func RemoveFollowing(followingId string, userId int) error {

	err := dao.DeleteFollowing(followingId, userId)
	return err
}

// GetFollowingListByUserId 根据用户id 获取 用户信息
func GetFollowingListByUserId(userId int, token string) ([]vo.UserInfo, error) {

	var userInfoList []vo.UserInfo
	follows, err := dao.QueryFollowingListByUserId(userId)
	if err != nil {
		return userInfoList, errors.New("db error")
	}

	var ids []int
	for _, value := range follows {
		ids = append(ids, int(value.FollowingId))
	}
	// id 集合去重
	setIds := utils.RemoveRepeatedElement(ids)
	infos, err := GetUsersByIds(setIds, token, userId)
	return infos, err

}

// GetFenshiListByUserId 根据用户id 获取 粉丝信息
func GetFenshiListByUserId(userId int, token string) ([]vo.UserInfo, error) {

	var userInfoList []vo.UserInfo
	// fenshi 数据
	fenshi, err := dao.QueryFenshiListByUserId(userId)
	if err != nil {
		return userInfoList, errors.New("db error")
	}

	var ids []int
	for _, value := range fenshi {
		ids = append(ids, int(value.UserId))
	}
	setIds := utils.RemoveRepeatedElement(ids)
	infos, err := GetUsersByIds(setIds, token, userId)
	return infos, err
}

// GetUsersByIds 获取用户信息
func GetUsersByIds(ids []int, token string, userId int) ([]vo.UserInfo, error) {

	claims, _ := utils.ParseToken(token)

	var users []po.User
	var userInfos []vo.UserInfo
	tx := dao.DB.Where("id in ?", ids).Find(&users)

	for _, user := range users {
		// 当前用户是否关注了视频发布者
		count := dao.QueryIsFollow(claims.UserId, int(user.ID))
		var IsFollow = false
		if count == 1 {
			IsFollow = true
		}
		// 获取关注和粉丝数
		followCount, followerCount := dao.QueryFollowingListByUserIdCount(userId)
		info := vo.UserInfo{
			UserId:          uint(user.ID),
			NickName:        user.NickName,
			FollowCount:     followCount,
			FollowerCount:   followerCount,
			IsFollow:        IsFollow,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			BackgroundImage: user.BackgroundImage,
		}
		userInfos = append(userInfos, info)
	}

	return userInfos, tx.Error
}
