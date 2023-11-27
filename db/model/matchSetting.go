package model

import "time"

// MatchSettingModel 时间段-匹配次数配置表模型
type MatchSettingModel struct {
	ID        int32     `gorm:"column:id" json:"id"`
	Count     string    `gorm:"column:count" json:"count"`
	Status    int32     `gorm:"column:status" json:"status"` // 0-未启用 1-启用
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
