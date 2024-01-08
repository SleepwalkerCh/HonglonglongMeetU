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
	OpenID    string    `gorm:"column:openid" json:"openID"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

const (
	MaleGender int = iota
	FemaleGender
)

const (
	AdminUserType int = iota
	NormalUserType
	InvalidUserType
)
