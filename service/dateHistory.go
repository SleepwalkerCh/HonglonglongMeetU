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
	RoomID         int    `json:"roomID"`
	Status         int    `json:"status"`
	ResultMale     int    `json:"resultMale"`
	ResultFemale   int    `json:"resultFemale"`
	DateTime       string `json:"dateTime"`
}
type DateHistoryGetResp struct {
	DateInfo []*DateInfo `json:"dateInfo"`
}

type DateHistoryPostReq struct {
	OperateUserID int `json:"operateUserID"`
	DateID        int `json:"dateID"`
	UserIDMale    int `json:"userIDMale"`
	UserIDFemale  int `json:"userIDFemale"`
	RoomID        int `json:"roomID"`
	Status        int `json:"status"`
	ResultMale    int `json:"resultMale"`
	ResultFemale  int `json:"resultFemale"`
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
			RoomID:         dateHistory.RoomID,
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
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""
	req, err := getDateHistoryPostReq(r)
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
	needUpdateDateHistory := false
	updateMap := make(map[string]interface{})
	if req.UserIDMale != -1 {
		updateMap["userid_male"] = req.UserIDMale
		needUpdateDateHistory = true
	}
	if req.UserIDFemale != -1 {
		updateMap["userid_female"] = req.UserIDFemale
		needUpdateDateHistory = true
	}
	if req.Status != -1 {
		updateMap["status"] = req.Status
		needUpdateDateHistory = true
	}
	if req.RoomID != -1 {
		updateMap["roomid"] = req.RoomID
		needUpdateDateHistory = true
	}
	if req.ResultMale != -1 {
		updateMap["result_male"] = req.ResultMale
		needUpdateDateHistory = true
	}
	if req.ResultFemale != -1 {
		updateMap["result_female"] = req.ResultFemale
		needUpdateDateHistory = true
	}
	if needUpdateDateHistory {
		err = dao.IDateHistoryInterface.UpdateDateHistoryByID(req.DateID, updateMap)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	res.Data = "success"
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

func getDateHistoryPostReq(r *http.Request) (req *DateHistoryPostReq, err error) {
	req = new(DateHistoryPostReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	operateUserID, ok := body["operateUserID"]
	if !ok {
		err = fmt.Errorf("缺少 operateUserID 参数")
		return
	}
	dateID, ok := body["dateID"]
	if !ok {
		err = fmt.Errorf("缺少 dateID 参数")
		return
	}
	req.OperateUserID = operateUserID.(int)
	req.DateID = dateID.(int)
	// init req
	req.UserIDMale = -1
	req.UserIDFemale = -1
	req.RoomID = -1
	req.Status = -1
	req.ResultMale = -1
	req.ResultFemale = -1
	userIDMale, ok := body["userIDMale"]
	if ok {
		req.UserIDMale = userIDMale.(int)
	}
	userIDFemale, ok := body["userIDFemale"]
	if ok {
		req.UserIDFemale = userIDFemale.(int)
	}
	roomID, ok := body["roomID"]
	if ok {
		req.RoomID = roomID.(int)
	}
	status, ok := body["status"]
	if ok {
		req.Status = status.(int)
	}
	resultMale, ok := body["resultMale"]
	if ok {
		req.ResultMale = resultMale.(int)
	}
	resultFemale, ok := body["resultFemale"]
	if !ok {
		req.ResultFemale = resultFemale.(int)
	}
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
