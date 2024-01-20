package service

import (
	"encoding/json"
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

type UserInfoDetail struct {
	UserID   int    `json:"userID"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
	NickName string `json:"nickName"`
	RealName string `json:"realName"`
	SeatID   int    `json:"seatID"`
	SeatNo   string `json:"seatNo"`
}

type UserInfoGetResp struct {
	UserInfo []*UserInfoDetail `json:"userInfo"`
	Count    int               `json:"count"`
}

type UserInfoPostReq struct {
	OperateUserID int    `json:"operateUserID"`
	UserID        int    `json:"userID"`
	Status        int    `json:"status"`
	NickName      string `json:"nickName"`
	RealName      string `json:"realName"`
	SeatID        int    `json:"seatID"`
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

func UserInfoGetFunc(r *http.Request) (res *JsonResult) {
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
	allSeatInfo, err := dao.ISeatInterface.GetAllSeats()
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	seatUserMap := make(map[int]*model.SeatModel)
	for _, seatInfo := range allSeatInfo {
		seatUserMap[seatInfo.UserID] = seatInfo
	}

	userInfoDetail := make([]*UserInfoDetail, 0)
	for _, user := range allNormalUsers {
		status := FreeStatus
		if tools.InIntList(user.ID, datingUserList) {
			status = DatingStatus
		}
		seatID, seatNo := 0, ""
		if seat, ok := seatUserMap[user.ID]; ok {
			seatID = seat.ID
			seatNo = seat.SeatNo
		}
		userInfoDetail = append(userInfoDetail, &UserInfoDetail{
			UserID:   user.ID,
			Gender:   user.Gender,
			Status:   status,
			NickName: user.NickName,
			RealName: user.RealName,
			SeatID:   seatID,
			SeatNo:   seatNo,
		})
	}

	res.Data = &UserInfoGetResp{
		UserInfo: userInfoDetail,
		Count:    len(userInfoDetail),
	}
	return
}

func UserInfoPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""
	req, err := getUserInfoPostReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	if !CheckAdminByUserID(req.OperateUserID) {
		err = fmt.Errorf("this operator is not admin")
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	needUpdateUser := false
	updateMap := make(map[string]interface{})
	if req.Status != -1 {
		updateMap["status"] = req.Status
		needUpdateUser = true
	}
	if req.NickName != "" {
		updateMap["nickname"] = req.NickName
		needUpdateUser = true
	}
	if req.NickName != "" {
		updateMap["realname"] = req.RealName
		needUpdateUser = true
	}
	if needUpdateUser {
		err = dao.IUserInterface.UpdateUserByID(req.UserID, updateMap)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	needUpdateSeat := false
	if req.SeatID != -1 {
		needUpdateSeat = true
	}
	if needUpdateSeat {
		seatInfos, err := dao.ISeatInterface.GetSeatBySeatID(req.SeatID)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
		if len(seatInfos) == 0 {
			err = fmt.Errorf("can not find seatID:%v", req.SeatID)
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
		if seatInfos[0].Status != model.FreeStatus {
			err = fmt.Errorf("can not change to seatID:%v", req.SeatID)
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
		// 移除原有座位
		if err = dao.ISeatInterface.RemoveSeatByUserID(req.UserID); err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
		// 放置入新的座位
		if err = dao.ISeatInterface.UpdateSeatBySeatID(req.SeatID, req.UserID); err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	return
}

func getUserInfoPostReq(r *http.Request) (req *UserInfoPostReq, err error) {
	req = new(UserInfoPostReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()
	// init request
	req.Status = -1
	req.SeatID = -1

	operateUserID, ok := body["operateUserID"]
	if !ok {
		err = fmt.Errorf("缺少 operateUserID 参数")
		return
	}
	userID, ok := body["userID"]
	if !ok {
		err = fmt.Errorf("缺少 userID 参数")
		return
	}
	req.OperateUserID = int(operateUserID.(float64))
	req.UserID = int(userID.(float64))
	status, ok := body["status"]
	if ok {
		req.Status = int(status.(float64))
	}
	nickName, ok := body["nickName"]
	if ok {
		req.NickName = nickName.(string)
	}
	realName, ok := body["realName"]
	if ok {
		req.RealName = realName.(string)
	}
	seatID, ok := body["seatID"]
	if !ok {
		req.SeatID = int(seatID.(float64))
	}
	return
}
func CheckAdminByUserID(userID int) (isAdmin bool) {
	userInfos, err := dao.IUserInterface.GetUsersByIDList([]int{userID})
	if err != nil {
		return false
	}
	if userInfo, ok := userInfos[userID]; ok {
		if userInfo.UserType == model.AdminUserType {
			return true
		}
	}
	return false
}

func GetNickNameFromUserInfoMap(userMap map[int]*model.UserModel, userID int) string {
	if userInfo, ok := userMap[userID]; ok {
		return userInfo.NickName
	}
	return ""

}

func GetGenderByUserID(userID int) (gender int, err error) {
	userMap, err := dao.IUserInterface.GetUsersByIDList([]int{userID})
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
