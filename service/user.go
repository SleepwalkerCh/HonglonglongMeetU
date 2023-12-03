package service

import (
	"fmt"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func GetNickNameFromUserInfoMap(userMap map[int]*model.UserModel, userID int) string {
	if userInfo, ok := userMap[userID]; ok {
		return userInfo.NickName
	}
	return ""

}

func GetGenderByUserID(userID int) (gender int, err error) {
	userMap, err := dao.IUserInterface.GetNormalUsersByIDList([]int{userID})
	if err != nil {
		return
	}
	if _, ok := userMap[userID]; !ok {
		err = fmt.Errorf("can not find corrent user,userID:%d", userID)
		return
	}
	gender = userMap[userID].Gender
	return
}
