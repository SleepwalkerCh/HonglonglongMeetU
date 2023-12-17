package model

import "time"

// BlindMatchModel 匹配结果模型
type BlindMatchModel struct {
	ID           int       `gorm:"column:id" json:"id"`
	UserIDMale   int       `gorm:"column:userid_male" json:"userid_male"`
	UserIDFemale int       `gorm:"column:userid_female" json:"userid_female"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

const InitUserID int = 0
