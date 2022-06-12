package service

import (
	"errors"
	"time"

	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/po"
	"github.com/3th-JustFunTeam/Tiktok-Backend/model/vo"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
)

// AddComment 添加视频评论
func AddComment(videoId int, token, commentText string) (vo.CommentVo, error) {

	var comment = vo.CommentVo{}
	claims, err := utils.ParseToken(token)
	if err != nil {
		return comment, errors.New("token 解析错误")
	}
	now := time.Now()
	// 封装信息
	temp := po.VideoCommon{
		UserId:    uint(claims.UserId),
		VideoId:   uint(videoId),
		Comment:   commentText,
		CreatedAt: now,
		UpdatedAt: now,
		IsDeleted: false,
	}
	// 保存到 数据库
	err = dao.SaveComment(temp)
	if err != nil {
		return comment, errors.New("sql 运行错误")
	}
	var user = po.User{}
	dao.DB.Where("id = ?", claims.UserId).Find(&user)

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
	// 时间格式化
	formatTime := now.Format("2006-01-02 15:04:05")
	comment.CreateDate = formatTime

	comment.Id = temp.Id
	comment.Content = commentText
	comment.User = info

	return comment, nil
}

// RemoveCommentById 删除评论
func RemoveCommentById(commentId, videoId int) error {

	var common = po.VideoCommon{}
	tx := dao.DB.Where("id = ? and video_id = ?", commentId, videoId).Delete(&common)
	return tx.Error
}

// GetCommentList 获取评论列表
func GetCommentList(videoId int, token string) ([]vo.CommentVo, error) {
	var commentVo []vo.CommentVo
	var comment []po.VideoCommon
	_, err := utils.ParseToken(token)
	if err != nil {
		return commentVo, errors.New("token 解析失败")
	}

	// 封装前端需要的信息
	tx := dao.DB.Where("video_id = ?", videoId).Find(&comment)
	if tx.Error != nil {
		return commentVo, errors.New("sql 错误")
	}
	var ids []int
	for _, value := range comment {
		ids = append(ids, int(value.UserId))
	}
	setIds := utils.RemoveRepeatedElement(ids)
	infos, err := GetUserInfoByIds(setIds, token)

	for _, common := range comment {

		for _, info := range infos {
			if common.UserId == info.UserId {
				formatTime := common.CreatedAt.Format("2006-01-02 15:04:05")
				vo := vo.CommentVo{
					Id:         common.Id,
					User:       info,
					Content:    common.Comment,
					CreateDate: formatTime,
				}
				commentVo = append(commentVo, vo)
			}
		}
	}
	return commentVo, err
}
