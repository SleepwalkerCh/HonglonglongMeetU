package model

import "time"

// SeatModel 座位模型
type SeatModel struct {
	ID        int32     `gorm:"column:id" json:"id"`
	SeatNo    string    `gorm:"column:seatNo" json:"seat_no"`
	UserID    int32     `gorm:"column:userID" json:"user_id"`
	Status    int32     `gorm:"column:status" json:"status"` // 0-空闲 1-已被占用
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updated_at"`
}

type SeatStatus int32

const (
	FreeStatus SeatStatus = iota
	OccupiedStatus
)
