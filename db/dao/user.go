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
	GetUsersByIDList(userID []int) (userMap map[int]*model.UserModel, err error)
	GetAllNormalUsers() ([]*model.UserModel, error)
	UpdateUserByID(userID int, updateMap map[string]interface{}) (err error)
	GetUserByOpenID(openID string) (user *model.UserModel, err error)
}

type UserModelInterfaceImp struct{}

var IUserInterface = &UserModelInterfaceImp{}

func (u *UserModelInterfaceImp) GetAllUsers() (users []*model.UserModel, err error) {
	cli := db.Get()
	users = make([]*model.UserModel, 0)
	if err = cli.Table(UserTableName).Where("status = ?", model.NormalStatus).Find(&users).Error; err != nil {
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

func (u *UserModelInterfaceImp) GetUsersByIDList(userID []int) (userMap map[int]*model.UserModel, err error) {
	cli := db.Get()
	userMap = make(map[int]*model.UserModel)
	users := make([]*model.UserModel, 0)
	if err = cli.Table(UserTableName).
		Where("id = ?", userID).
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

func (u *UserModelInterfaceImp) GetAllNormalUsers() (users []*model.UserModel, err error) {
	cli := db.Get()
	users = make([]*model.UserModel, 0)
	if err = cli.Table(UserTableName).Where("user_type = ?", model.NormalUserType).Where("status = ?", model.NormalStatus).Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *UserModelInterfaceImp) UpdateUserByID(userID int, updateMap map[string]interface{}) (err error) {
	cli := db.Get()
	if err = cli.Table(UserTableName).Where("id = ?", userID).Updates(updateMap).Error; err != nil {
		return
	}
	return
}

func (u *UserModelInterfaceImp) GetUserByOpenID(openID string) (user *model.UserModel, err error) {
	userList := make([]*model.UserModel, 0)
	cli := db.Get()
	if err = cli.Table(UserTableName).Where("openid = ?", openID).Find(&userList).Error; err != nil {
		return
	}
	if len(userList) > 0 {
		user = userList[0]
	}
	return
}
