package dao

import "github.com/3th-JustFunTeam/Tiktok-Backend/model/po"

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
