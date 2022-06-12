package service

import (
	"errors"
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/common"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"time"
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
func AddVideo(userId int, title, fileName string) error {

	video := po.Video{
		UserId:      uint(userId),
		PlayUrl:     fileName,
		CoverUrl:    common.DEFAULT_COVERURL,
		Title:       title,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err := dao.SaveVideo(video)
	return err
}
