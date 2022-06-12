package po

import "time"

type Love struct {
	ID        uint64    `gorm:"column:id;type:bigint unsigned auto_increment;primaryKey;not null;comment:主键id"`
	UserId    uint      `gorm:"type:bigint;comment:用户id"`
	VideoId   uint      `gorm:"type:bigint;comment:视频id"`
	CreatedAt time.Time `gorm:"column:ts;type:timestamp;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:update_ts;type:timestamp;not null;comment:修改时间"`
}
