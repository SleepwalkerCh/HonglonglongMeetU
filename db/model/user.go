package model

import "time"

// UserModel 用户模型
type UserModel struct {
	ID        int       `gorm:"column:id" json:"id"`
	NickName  string    `gorm:"column:nickname" json:"nickName"`
	RealName  string    `gorm:"column:realname" json:"realName"`
	Gender    int       `gorm:"column:gender" json:"gender"`
	UserType  int       `gorm:"column:user_type" json:"userType"`
	Status    int       `gorm:"column:status" json:"status"` // 0-异常 1-正常
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

const (
	MaleGender int = iota
)

const (
	AdminUserType int = iota
	NormalUserType
	InvalidUserType
)
