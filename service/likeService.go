package service

import (
	"errors"
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"strconv"
	"time"
)

func Like(userId, videoId int) error {

	now := time.Now()
	var love = po.Love{
		UserId:    uint(userId),
		VideoId:   uint(videoId),
		CreatedAt: now,
		UpdatedAt: now,
	}
	return dao.SaveLike(love)
}

func DownLike(userId, videoId int) error {

	err, row := dao.DeleteLike(userId, videoId)
	if row != 1 {
		return errors.New("不能再删除了")
	}
	return err
}
func FindLikeList(userId, token string, currentUserId int) []vo.VideoVo {
	var res []vo.VideoVo
	loves, err := dao.QueryLikeByUserId(userId)
	if err != nil {
		return res
	}
	for _, love := range loves {
		video, _ := dao.QueryVideoByVideoId(love.VideoId)
		uId, _ := strconv.Atoi(userId)
		userInfo, _ := GetUserInfoById(uId, currentUserId)
		count := dao.CommentCount(video.VideoId)
		likeCount := dao.VideoLikeCount(video.VideoId)
		isLike := CheckIsLike(token, video.VideoId)
		var videoVo = vo.VideoVo{
			Id:            video.VideoId,
			Author:        userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: uint64(likeCount),
			CommentCount:  count,
			IsFavorite:    isLike,
			Title:         video.Title,
		}
		res = append(res, videoVo)
	}
	return res
}
