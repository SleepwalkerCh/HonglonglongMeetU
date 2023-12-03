package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
)

type DateStartPostReq struct {
	MaleUserID   int `json:"maleUserID"`
	FemaleUserID int `json:"femaleUserID"`
	RoomID       int `json:"roomID"`
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
