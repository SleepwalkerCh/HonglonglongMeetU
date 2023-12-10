package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
)

type DateHistoryGetReq struct {
	OperateUserID int `json:"operateUserID"`
}

type DateInfo struct {
	DateID         int    `json:"dateID"`
	UserIDMale     int    `json:"userIDMale"`
	UserIDFemale   int    `json:"userIDFemale"`
	NicknameMale   string `json:"nicknameMale"`
	NicknameFemale string `json:"nicknameFemale"`
	RoomName       string `json:"roomName"`
	Status         int    `json:"status"`
	ResultMale     int    `json:"resultMale"`
	ResultFemale   int    `json:"resultFemale"`
	DateTime       string `json:"dateTime"`
}
type DateHistoryGetResp struct {
	DateInfo []*DateInfo `json:"dateInfo"`
}

func DateHistoryGetFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""
	req, err := getDateHistoryGetReq(r)
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
	dateHistoryList, err := dao.IDateHistoryInterface.GetAllDateHistory()
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	userIDList := make([]int, 0)
	for _, dateHistory := range dateHistoryList {
		userIDList = append(userIDList, dateHistory.UserIDMale)
		userIDList = append(userIDList, dateHistory.UserIDFemale)
	}
	userMap, err := dao.IUserInterface.GetUsersByIDList(userIDList)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	data := make([]*DateInfo, 0)
	for _, dateHistory := range dateHistoryList {
		data = append(data, &DateInfo{
			DateID:         dateHistory.ID,
			UserIDMale:     dateHistory.UserIDMale,
			UserIDFemale:   dateHistory.UserIDFemale,
			NicknameMale:   GetNickNameFromUserInfoMap(userMap, dateHistory.UserIDMale),
			NicknameFemale: GetNickNameFromUserInfoMap(userMap, dateHistory.UserIDFemale),
			RoomName:       "",
			Status:         dateHistory.Status,
			ResultMale:     dateHistory.ResultMale,
			ResultFemale:   dateHistory.ResultFemale,
			DateTime:       dateHistory.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	res.Data = data
	return
}

func DateHistoryPostFunc(r *http.Request) (res *JsonResult) {

	return
}
func getDateHistoryGetReq(r *http.Request) (req *DateHistoryGetReq, err error) {
	req = new(DateHistoryGetReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	operateUserID, ok := body["operateUserID"]
	if !ok {
		err = fmt.Errorf("缺少 operateUserID 参数")
		return
	}
	req.OperateUserID = operateUserID.(int)
	return
}

//func getDateHistoryPostReq(r *http.Request) (req *DateHistoryPostReq, err error) {
//	req = new(UserInfoPostReq)
//	decoder := json.NewDecoder(r.Body)
//	body := make(map[string]interface{})
//	if err = decoder.Decode(&body); err != nil {
//		return
//	}
//	defer r.Body.Close()
//	// init request
//	req.Status = -1
//	req.SeatID = -1
//
//	operateUserID, ok := body["operateUserID"]
//	if !ok {
//		err = fmt.Errorf("缺少 operateUserID 参数")
//		return
//	}
//	userID, ok := body["userID"]
//	if !ok {
//		err = fmt.Errorf("缺少 userID 参数")
//		return
//	}
//	req.OperateUserID = operateUserID.(int)
//	req.UserID = userID.(int)
//	status, ok := body["status"]
//	if ok {
//		req.Status = status.(int)
//	}
//	nickName, ok := body["nickName"]
//	if ok {
//		req.NickName = nickName.(string)
//	}
//	realName, ok := body["realName"]
//	if ok {
//		req.RealName = realName.(string)
//	}
//	seatID, ok := body["seatID"]
//	if !ok {
//		req.SeatID = seatID.(int)
//	}
//	return
//}
