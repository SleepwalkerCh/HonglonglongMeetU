package dao

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const UserTableName = "user"

type UserModelInterface interface {
	GetAllUsers() ([]*model.UserModel, error)
	InsertUser(user *model.UserModel) error
	GetNormalUsersByIDList(userID []int32) (userMap map[int32]*model.UserModel, err error)
}

type UserModelInterfaceImp struct{}

var IUserInterface = &UserModelInterfaceImp{}

func (u *UserModelInterfaceImp) GetAllUsers() (users []*model.UserModel, err error) {
	cli := db.Get()
	users = make([]*model.UserModel, 0)
	if err = cli.Table(UserTableName).Where("status = ?", model.NormalStatus).Find(users).Error; err != nil {
		return
	}
	return
}

func (u *UserModelInterfaceImp) InsertUser(user *model.UserModel) (err error) {
	cli := db.Get()
	if err = cli.Table(UserTableName).Save(user).Error; err != nil {
		return
	}
	return
}

func (u *UserModelInterfaceImp) GetNormalUsersByIDList(userID []int32) (userMap map[int32]*model.UserModel, err error) {
	cli := db.Get()
	userMap = make(map[int32]*model.UserModel)
	users := make([]*model.UserModel, 0)
	if err = cli.Table(UserTableName).
		Where("id = ?", userID).
		Where("status = ?", model.NormalStatus).
		Where("user_type = ?", model.NormalUserType).
		Find(users).Error; err != nil {
		return
	}
	if len(users) == 0 {
		err = fmt.Errorf("[GetNormalUserByID]Not found normalUser by userid:%d", userID)
		return
	}
	for _, user := range users {
		userMap[user.ID] = user
	}
	return
}
