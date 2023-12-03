package model

import "time"

// DateRoomModel 当前约会房间模型
type DateRoomModel struct {
	ID           int       `gorm:"column:id" json:"id"`
	RoomName     string    `gorm:"column:room_name" json:"room_name"`
	UserIDMale   int       `gorm:"column:userid_male" json:"userid_male"`
	UserIDFemale int       `gorm:"column:userid_female" json:"userid_female"`
	Status       int       `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
