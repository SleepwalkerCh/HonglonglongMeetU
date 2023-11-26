package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const UserTableName = "user"

type UserModelInterface interface {
	GetAllUsers() ([]*model.UserModel, error)
	InsertUser(user *model.UserModel) error
}

type UserModelInterfaceImp struct{}

var IUserInterface = &UserModelInterfaceImp{}

func (u *UserModelInterfaceImp) GetAllUsers() (users []*model.UserModel, err error) {
	cli := db.Get()
	users = make([]*model.UserModel, 0)
	if err = cli.Table(UserTableName).Where("status = ?", 1).Find(users).Error; err != nil {
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
