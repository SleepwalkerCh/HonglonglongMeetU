package service

import (
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/tools"
)

const (
	FreeStatus int = iota
	DatingStatus
)

type UserInfo struct {
	UserID   int    `json:"userID"`
	Status   int    `json:"status"`
	NickName string `json:"nickName"`
}

type AllUserStatusGetResp struct {
	UserInfo []*UserInfo `json:"userInfo"`
	Count    int         `json:"count"`
}

func AllUserStatusGetFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	allNormalUsers, err := dao.IUserInterface.GetAllNormalUsers()
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	dateHistoryList, err := dao.IDateHistoryInterface.GetAllDatingDateHistory()
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	datingUserList := make([]int, 0)
	for _, dateHistory := range dateHistoryList {
		datingUserList = append(datingUserList, dateHistory.UserIDMale)
		datingUserList = append(datingUserList, dateHistory.UserIDFemale)
	}

	userInfoList := make([]*UserInfo, 0)
	for _, normalUser := range allNormalUsers {
		status := FreeStatus
		if tools.InIntList(normalUser.ID, datingUserList) {
			status = DatingStatus
		}
		userInfoList = append(userInfoList, &UserInfo{
			UserID:   normalUser.ID,
			Status:   status,
			NickName: normalUser.NickName,
		})
	}
	res.Data = &AllUserStatusGetResp{
		UserInfo: userInfoList,
		Count:    len(userInfoList),
	}
	return
}

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
