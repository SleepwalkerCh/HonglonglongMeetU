package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

type DateStartPostReq struct {
	MaleUserID   int `json:"maleUserID"`
	FemaleUserID int `json:"femaleUserID"`
	RoomID       int `json:"roomID"`
}
type DateStopPostReq struct {
	MaleUserID   int `json:"maleUserID"`
	FemaleUserID int `json:"femaleUserID"`
	RoomID       int `json:"roomID"`
}

type DateResultSubmitPostReq struct {
	UserID int `json:"userID"`
	RoomID int `json:"roomID"`
	Result int `json:"result"`
}

func DateStartPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getDateStartPostReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	err = dao.IDateHistoryInterface.CreateDateHistoryRecord(req.RoomID, req.MaleUserID, req.FemaleUserID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	res.Data = "success"
	return
}

func DateStopPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getDateStopPostReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	// dateHistory与入参不符
	if err = CheckDateHistory(req.RoomID, req.MaleUserID, req.FemaleUserID); err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	err = dao.IDateHistoryInterface.UpdateDateHistoryStatus(req.RoomID, model.FinishedStatus)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	res.Data = "success"
	return
}

func DateResultSubmitPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getDateResultSubmitPostReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	gender, err := GetGenderByUserID(req.UserID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	dateHistoryList, err := dao.IDateHistoryInterface.GetDateHistoryByUserIDAndGender(req.UserID, gender)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	isCorrect := false
	for _, dateHistory := range dateHistoryList {
		if dateHistory.RoomID == req.RoomID {
			if err = dao.IDateHistoryInterface.UpdateDateHistoryResultByIDAndGender(dateHistory.ID, gender, req.Result); err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
				return
			}
			isCorrect = true
		}
	}
	if !isCorrect {
		err = fmt.Errorf("request is in correct,userID:%v,roomID:%v", req.UserID, req.RoomID)
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	res.Data = "success"
	return
}

func CheckDateHistory(roomID, maleUserID, femaleUserID int) (err error) {

	dateHistory, err := dao.IDateHistoryInterface.GetDateHistoryByRoomIDAndStatus(roomID, model.DatingStatus)
	if err != nil {
		return
	}
	if len(dateHistory) != 1 || dateHistory[0].UserIDMale != maleUserID || dateHistory[0].UserIDFemale != femaleUserID {
		err = fmt.Errorf("dateHistory is incorrect,roomID:%d", roomID)
		return
	}
	return
}

func getDateStopPostReq(r *http.Request) (req *DateStopPostReq, err error) {
	req = new(DateStopPostReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	maleUserID, ok := body["maleUserID"]
	if !ok {
		err = fmt.Errorf("缺少 maleUserID 参数")
		return
	}
	femaleUserID, ok := body["femaleUserID"]
	if !ok {
		err = fmt.Errorf("缺少 femaleUserID 参数")
		return
	}
	roomID, ok := body["roomID"]
	if !ok {
		err = fmt.Errorf("缺少 roomID 参数")
		return
	}

	req.MaleUserID = maleUserID.(int)
	req.FemaleUserID = femaleUserID.(int)
	req.RoomID = roomID.(int)
	return
}

func getDateStartPostReq(r *http.Request) (req *DateStartPostReq, err error) {
	req = new(DateStartPostReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	maleUserID, ok := body["maleUserID"]
	if !ok {
		err = fmt.Errorf("缺少 maleUserID 参数")
		return
	}
	femaleUserID, ok := body["femaleUserID"]
	if !ok {
		err = fmt.Errorf("缺少 femaleUserID 参数")
		return
	}
	roomID, ok := body["roomID"]
	if !ok {
		err = fmt.Errorf("缺少 roomID 参数")
		return
	}

	req.MaleUserID = maleUserID.(int)
	req.FemaleUserID = femaleUserID.(int)
	req.RoomID = roomID.(int)
	return
}

func getDateResultSubmitPostReq(r *http.Request) (req *DateResultSubmitPostReq, err error) {
	req = new(DateResultSubmitPostReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	userID, ok := body["userID"]
	if !ok {
		err = fmt.Errorf("缺少 userID 参数")
		return
	}
	roomID, ok := body["roomID"]
	if !ok {
		err = fmt.Errorf("缺少 roomID 参数")
		return
	}
	result, ok := body["result"]
	if !ok {
		err = fmt.Errorf("缺少 result 参数")
		return
	}

	req.UserID = userID.(int)
	req.RoomID = roomID.(int)
	req.Result = result.(int)
	return
}
