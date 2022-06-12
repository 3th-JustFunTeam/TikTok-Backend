package dao

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
)

// 投稿后的视频点赞数存入数据库
func SaveLike(favorite po.Favorite) error {
	tx := DB.Create(&favorite)
	return tx.Error
}

// 点赞数+1
func LikeUp(video_id uint) error {
	tx := DB.Exec("UPDATE favorites SET like = like + 1 WHERE video_id = ?", video_id)
	return tx.Error
}

// 点赞数-1
func LikeDown(video_id uint) error {
	tx := DB.Exec("UPDATE favorites SET `like` = `like` - 1 WHERE video_id = ?", video_id)
	return tx.Error
}
