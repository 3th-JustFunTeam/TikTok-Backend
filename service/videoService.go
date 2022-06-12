package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/common"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
)

// GetAllVideo 获取所有视频(最前发布20个)
func GetAllVideo(token string) ([]vo.VideoVo, error) {

	video, err := dao.QueryAllVideo()
	var videoVoList []vo.VideoVo
	if err != nil {
		return videoVoList, errors.New("获取视频失败")
	}
	// 获取用户用户id集合
	var ids []int
	for _, value := range video {
		ids = append(ids, int(value.UserId))
	}
	setIds := utils.RemoveRepeatedElement(ids)

	infos, err := GetUserInfoByIds(setIds, token)

	// 封装数据
	for _, v := range video {
		for _, info := range infos {

			if v.UserId == info.UserId {
				count := dao.CommentCount(v.VideoId)
				videoVo := vo.VideoVo{
					Id:           v.VideoId,
					Author:       info,
					PlayUrl:      common.VIDEO_RUL_PREFIX + common.VIDEO_RUL_MID + common.VIDEO_RUL_SUFFIX + v.PlayUrl,
					CoverUrl:     v.CoverUrl,
					CommentCount: count,
					IsFavorite:   false,
					Title:        v.Title,
				}
				videoVoList = append(videoVoList, videoVo)
			}
		}
	}
	return videoVoList, err
}

// GetUserInfoByIds 根据用户的id集合获取用户的信息
func GetUserInfoByIds(ids []int, token string) ([]vo.UserInfo, error) {

	var users []po.User
	var userInfos []vo.UserInfo
	tx := dao.DB.Where("id in ?", ids).Find(&users)

	c, err := utils.ParseToken(token)
	if err == nil { // 登录状态
		for _, user := range users {
			// 当前用户是否关注了视频发布者
			count := dao.QueryIsFollow(c.UserId, int(user.ID))
			var IsFollow = false
			if count == 1 {
				IsFollow = true
			}
			// 获取关注和粉丝数
			followCount, followerCount := dao.QueryFollowingListByUserIdCount(int(user.ID))
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

	} else { // 未登录状态
		for _, user := range users {
			followCount, followerCount := dao.QueryFollowingListByUserIdCount(int(user.ID))
			info := vo.UserInfo{
				UserId:          uint(user.ID),
				NickName:        user.NickName,
				FollowCount:     followCount,
				FollowerCount:   followerCount,
				IsFollow:        false, // 未登录状态  都是未关注
				Avatar:          user.Avatar,
				Signature:       user.Signature,
				BackgroundImage: user.BackgroundImage,
			}
			userInfos = append(userInfos, info)
		}
	}

	return userInfos, tx.Error

}

// AddVideo 添加视频
func AddVideo(token, title, fileName string) error {

	// 解析token
	claims, err2 := utils.ParseToken(token)

	if err2 != nil {
		return errors.New("token 解析错误")
	}

	video := po.Video{
		UserId:      uint(claims.UserId),
		PlayUrl:     fileName,
		CoverUrl:    common.DEFAULT_COVERURL,
		Title:       title,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err := dao.SaveVideo(video)
	return err
}

// 根据单个id获取单个视频作者信息（用于发布列表）
func GetUserInfoById(uId, userid int) (vo.UserInfo, error) {
	var user po.User
	var userInfo vo.UserInfo
	tx := dao.DB.Where("id = ?", uId).Find(&user)
	followCount, followerCount := dao.QueryFollowingListByUserIdCount(int(user.ID))
	var isFollow = false
	// 判断是否其他用户还是自己
	if userid != uId {
		count := dao.QueryIsFollow(userid, uId)
		if count < 0 {
			isFollow = true
		}
	}
	userInfo = vo.UserInfo{
		UserId:        uint(user.ID),
		NickName:      user.NickName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}

	return userInfo, tx.Error
}

// 根据单个id获取用户的全部Video（用于发布列表）
func GetUserAllVideo(id, token string) ([]vo.VideoVo, error) {
	claims, err2 := utils.ParseToken(token)

	if err2 != nil {
		return []vo.VideoVo{}, errors.New("error")
	}
	var err = errors.New("")
	var VoList []vo.VideoVo
	var userInfo vo.UserInfo

	uId, _ := strconv.Atoi(id)
	userInfo, err = GetUserInfoById(uId, claims.UserId)
	videoList, _ := dao.QueryUserAllVideo(id)
	for _, video := range videoList {
		var videoVo = vo.VideoVo{
			Id:       video.VideoId,
			Author:   userInfo,
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			// 点赞
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         video.Title,
		}
		VoList = append(VoList, videoVo)
	}
	return VoList, err
}
