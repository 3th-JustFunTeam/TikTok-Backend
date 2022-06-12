package po

import "time"

// Video 视频表
type Video struct {
	VideoId     uint      `gorm:"primarykey;comment:视频id"`
	UserId      uint      `gorm:"type:bigint;comment:用户id"`
	PlayUrl     string    `gorm:"type:varchar(500);comment:视频播放地址"`
	CoverUrl    string    `gorm:"type:varchar(500);comment:视频封面地址"`
	Title       string    `gorm:"type:varchar(250);comment:视频标题"`
	CreatedTime time.Time `gorm:"comment:创建时间"`
	UpdatedTime time.Time `gorm:"comment:修改时间"`
	LikeCount   uint      `gorm:"comment:点赞数"`
}
