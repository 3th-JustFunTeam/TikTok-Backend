package service

import (
	"errors"
	"time"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/common"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
)

// GetAllVideo 获取所有视频(最前发布20个)
func GetAllVideo() ([]vo.VideoVo, error) {

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
	infos, err := GetUserInfoByIds(setIds)

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
				}
				videoVoList = append(videoVoList, videoVo)
			}
		}
	}
	return videoVoList, err
}

// GetUserInfoByIds 根据用户的id集合获取用户的信息
func GetUserInfoByIds(ids []int) ([]vo.UserInfo, error) {
	var users []po.User
	var userInfos []vo.UserInfo
	tx := dao.DB.Where("id in ?", ids).Find(&users)

	for _, user := range users {
		// todo 用户的关注和粉丝没写
		info := vo.UserInfo{
			UserId:          uint(user.ID),
			NickName:        user.NickName,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        false,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			BackgroundImage: user.BackgroundImage,
		}
		userInfos = append(userInfos, info)
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
func GetUserInfoById(id uint) (vo.UserInfo, error) {
	var user po.User
	tx := dao.DB.Where("id = ?", id).Find(&user)
	userInfo := vo.UserInfo{
		// 关注和粉丝没写
		UserId:        uint(user.ID),
		NickName:      user.NickName,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	return userInfo, tx.Error
}

// 根据单个id获取用户的全部Video（用于发布列表）
func GetUserAllVideo(id uint) ([]vo.VideoVo, error) {
	videoList, err := dao.QueryUserAllVideo(id)
	return videoList, err
}
