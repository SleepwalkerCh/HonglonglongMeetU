package model

import "time"

// CounterModel 计数器模型
type CounterModel struct {
	Id        int       `gorm:"column:id" json:"id"`
	Count     int       `gorm:"column:count" json:"count"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
