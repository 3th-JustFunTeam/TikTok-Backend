package dao

import "github.com/3th-JustFunTeam/Tiktok-Backend/model/po"

// SaveComment 保存评论
func SaveComment(common po.VideoCommon) error {

	tx := DB.Create(&common)
	return tx.Error
}

// CommentCount 视频的评论总数
func CommentCount(id uint) uint64 {

	var count int64
	DB.Model(&po.VideoCommon{}).Where("video_id = ?", id).Count(&count)

	return uint64(count)

}
