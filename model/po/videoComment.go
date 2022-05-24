package po

import "time"

type VideoCommon struct {
	Id        uint      `gorm:"primarykey;comment:主键id"`
	UserId    uint      `gorm:"type:bigint;comment:用户id"`
	VideoId   uint      `gorm:"type:bigint;comment:视频id"`
	Comment   string    `gorm:"type:text;comment:评论内容"`
	CreatedAt time.Time `gorm:"column:ts;type:timestamp;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:update_ts;type:timestamp;not null;comment:修改时间"`
	IsDeleted bool      `gorm:"column:is_deleted;type:bool;not null;comment:是否删除"`
}
