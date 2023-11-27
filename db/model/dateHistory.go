package model

import "time"

// DateHistoryModel 约会历史模型
type DateHistoryModel struct {
	ID           int32     `gorm:"column:id" json:"id"`
	UserIDMale   int32     `gorm:"column:userid_male" json:"userid_male"`
	UserIDFemale int32     `gorm:"column:userid_female" json:"userid_female"`
	RoomID       int32     `gorm:"column:roomid" json:"roomid"`
	ResultMale   int32     `gorm:"column:result_male" json:"result_male"`
	ResultFemale int32     `gorm:"column:result_female" json:"result_female"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
