package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

type RoomInfoGetReq struct {
	ShowAll bool `json:"showAll"`
	RoomID  int  `json:"roomID"`
}
type RoomInfo struct {
	RoomID         int    `json:"roomID"`
	RoomName       string `json:"roomName"`
	MaleUserID     int    `json:"maleUserID"`
	MaleUserName   string `json:"maleUserName"`
	FemaleUserID   int    `json:"femaleUserID"`
	FemaleUserName string `json:"femaleUserName"`
}

type RoomInfoGetResp struct {
	IsAllRoom bool        `json:"isAllRoom"`
	RoomInfo  []*RoomInfo `json:"roomInfo"`
}

type DateRoomInPostReq struct {
	UserID int `json:"userID"`
	RoomID int `json:"roomID"`
}
type DateRoomOutPostReq struct {
	UserID int `json:"userID"`
	RoomID int `json:"roomID"`
}

func RoomInfoGetFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getRoomInfoGetReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	dateRooms := make([]*model.DateRoomModel, 0)
	if req.ShowAll {
		dateRooms, err = dao.IDateRoomInterface.GetAllDateRoom()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	} else {
		dateRooms, err = dao.IDateRoomInterface.GetDateRoomByID(req.RoomID)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	roomInfo := MakeDateRoomInfo(dateRooms)
	res.Data = &RoomInfoGetResp{
		IsAllRoom: req.ShowAll,
		RoomInfo:  roomInfo,
	}
	return
}

func DateRoomInPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getDateRoomInPostReq(r)
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
	dateRooms, err := dao.IDateRoomInterface.GetDateRoomByID(req.RoomID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	if len(dateRooms) == 0 {
		err = fmt.Errorf("roomID is invalid,roomID:%v", req.RoomID)
		return
	}
	dateRoom := dateRooms[0]
	if gender == model.MaleGender {
		if dateRoom.UserIDMale != 0 {
			err = fmt.Errorf("room is fulled,roomID:%v", req.RoomID)
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	} else {
		if dateRoom.UserIDFemale != 0 {
			err = fmt.Errorf("room is fulled,roomID:%v", req.RoomID)
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	//可能有并发问题
	err = dao.IDateRoomInterface.UpdateRoomWithUserIDAndGender(req.RoomID, req.UserID, gender, dao.InAction)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	return
}
func DateRoomOutPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getDateRoomOutPostReq(r)
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
	dateRooms, err := dao.IDateRoomInterface.GetDateRoomByID(req.RoomID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	if len(dateRooms) == 0 {
		err = fmt.Errorf("roomID is invalid,roomID:%v", req.RoomID)
		return
	}
	dateRoom := dateRooms[0]
	if gender == model.MaleGender {
		if dateRoom.UserIDMale != req.UserID {
			err = fmt.Errorf("this user is not in this room,roomID:%v,userID:%v", req.RoomID, req.UserID)
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	} else {
		if dateRoom.UserIDFemale != req.UserID {
			err = fmt.Errorf("this user is not in this room,roomID:%v,userID:%v", req.RoomID, req.UserID)
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	//可能有并发问题
	err = dao.IDateRoomInterface.UpdateRoomWithUserIDAndGender(req.RoomID, req.UserID, gender, dao.OutAction)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	return
}

func getRoomInfoGetReq(r *http.Request) (req *RoomInfoGetReq, err error) {
	req = new(RoomInfoGetReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	showAll, ok := body["showAll"]
	if !ok {
		err = fmt.Errorf("缺少 showAll 参数")
		return
	}
	roomID, ok := body["roomID"]
	if !ok {
		err = fmt.Errorf("缺少 roomID 参数")
		return
	}
	req.ShowAll = showAll.(bool)
	req.RoomID = int(roomID.(float64))
	return
}

func getDateRoomInPostReq(r *http.Request) (req *DateRoomInPostReq, err error) {
	req = new(DateRoomInPostReq)
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
	req.UserID = int(userID.(float64))
	req.RoomID = int(roomID.(float64))
	return
}
func getDateRoomOutPostReq(r *http.Request) (req *DateRoomOutPostReq, err error) {
	req = new(DateRoomOutPostReq)
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
	req.UserID = int(userID.(float64))
	req.RoomID = int(roomID.(float64))
	return
}

func MakeDateRoomInfo(dateRooms []*model.DateRoomModel) (roomInfo []*RoomInfo) {
	userIDList := make([]int, 0)
	for _, dateRoom := range dateRooms {
		userIDList = append(userIDList, dateRoom.UserIDMale)
		userIDList = append(userIDList, dateRoom.UserIDFemale)
	}
	userMap, err := dao.IUserInterface.GetUsersByIDList(userIDList)
	if err != nil {
		return
	}
	roomInfo = make([]*RoomInfo, 0)
	for _, dateRoom := range dateRooms {
		roomInfo = append(roomInfo, &RoomInfo{
			RoomID:         dateRoom.ID,
			RoomName:       dateRoom.RoomName,
			MaleUserID:     dateRoom.UserIDMale,
			MaleUserName:   GetNickNameFromUserInfoMap(userMap, dateRoom.UserIDMale),
			FemaleUserID:   dateRoom.UserIDFemale,
			FemaleUserName: GetNickNameFromUserInfoMap(userMap, dateRoom.UserIDFemale),
		})
	}
	return
}
