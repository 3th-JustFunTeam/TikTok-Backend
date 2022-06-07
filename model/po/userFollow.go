package po

import "time"

type Follow struct {
	Id          uint      `gorm:"primarykey;comment:   主键id"`
	UserId      uint      `gorm:"type:bigint;comment:  用户id"`
	FollowingId uint      `gorm:"type:bigint;comment:关注的用户id"`
	CreatedAt   time.Time `gorm:"column:ts;type:timestamp;not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"column:update_ts;type:timestamp;not null;comment:修改时间"`
}
