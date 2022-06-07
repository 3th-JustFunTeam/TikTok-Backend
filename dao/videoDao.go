package dao

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
)

// QueryAllVideo 查询全部视频
func QueryAllVideo() ([]po.Video, error) {
	var video []po.Video
	tx := DB.Limit(30).Order("created_time desc").Find(&video)
	return video, tx.Error

}

// SaveVideo 保存视频
func SaveVideo(video po.Video) error {

	tx := DB.Create(&video)
	return tx.Error
}

// 根据用户id查询用户的视频列表，用于（发布列表）
func QueryUserAllVideo(id uint) ([]vo.VideoVo, error) {
	var video []vo.VideoVo
	// 以创建时间降序获取视频
	tx := DB.Where("id = ?", id).Order("created_time desc").Find(&video)
	return video, tx.Error
}
