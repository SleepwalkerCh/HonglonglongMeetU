package model

import "time"

// UserModel 用户模型
type UserModel struct {
	ID        int32     `gorm:"column:id" json:"id"`
	NickName  string    `gorm:"column:nickName" json:"nickName"`
	RealName  string    `gorm:"column:realName" json:"realName"`
	Gender    int32     `gorm:"column:gender" json:"gender"`
	UserType  int32     `gorm:"column:userType" json:"userType"`
	Status    int32     `gorm:"column:status" json:"status"` // 0-异常 1-正常
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type Gender int32

const (
	MaleGender Gender = iota
)

type UserType int32

const (
	AdminUserType UserType = iota
	NormalUserType
	InvalidUserType
)
