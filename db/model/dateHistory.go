package model

import "time"

// DateHistoryModel 约会历史模型
type DateHistoryModel struct {
	ID           int       `gorm:"column:id" json:"id"`
	UserIDMale   int       `gorm:"column:userid_male" json:"userid_male"`
	UserIDFemale int       `gorm:"column:userid_female" json:"userid_female"`
	RoomID       int       `gorm:"column:roomid" json:"roomid"`
	ResultMale   int       `gorm:"column:result_male" json:"result_male"`
	ResultFemale int       `gorm:"column:result_female" json:"result_female"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
