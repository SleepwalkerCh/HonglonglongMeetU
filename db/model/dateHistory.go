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
	Status       int       `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// ResultMale/ResultFemale 0-未知 1-拒绝 2-同意

const (
	InitResult int = iota
	DenyResult
	AcceptResult
)

// DateHistory 0-约会中 1-约会结束 2-评价完成

const (
	DatingStatus int = iota
	FinishedStatus
)
