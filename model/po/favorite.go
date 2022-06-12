package po

import "time"

type Favorite struct {
	Video_Id    uint      `gorm:"comment:视频id"`
	User_Id     uint      `gorm:"comment:用户id"`
	Like        uint      `gorm:"comment:点赞数量"`
	CreatedTime time.Time `gorm:"comment:创建时间"`
	UpdatedTime time.Time `gorm:"comment:修改时间"`
}
