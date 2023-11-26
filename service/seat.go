package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

type SeatGetReq struct {
	UserID int `json:"userID"`
}
type SeatGetData struct {
	HasSeat bool   `json:"hasSeat"`
	SeatID  int    `json:"seatID"`
	SeatNo  string `json:"seatNo"`
}

type SeatPostReq struct {
	UserID int `json:"userID"`
	SeatID int `json:"seatID"`
}

func SeatGetFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	//解析入参
	req, err := getSeatGetReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	seat, err := dao.ISeatInterface.GetSeatByUserID(req.UserID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	hasSeat := false
	seatID := 0
	seatNo := ""
	if len(seat) != 0 {
		hasSeat = true
		seatID = int(seat[0].ID)
		seatNo = seat[0].SeatNo
	}
	res.Data = &SeatGetData{
		HasSeat: hasSeat,
		SeatID:  seatID,
		SeatNo:  seatNo,
	}
	return
}

func SeatPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	//解析入参
	req, err := getSeatPostReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	seat, err := dao.ISeatInterface.GetSeatByUserID(req.UserID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	if len(seat) != 0 {
		res.Code = -1
		res.ErrorMsg = "该用户已有座位"
		return
	}
	seat, err = dao.ISeatInterface.GetSeatBySeatID(req.SeatID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	if len(seat) != 0 && seat[0].Status != int32(model.FreeStatus) {
		res.Code = -1
		res.ErrorMsg = "该座位已被占用"
		return
	}
	err = dao.ISeatInterface.UpdateSeatBySeatID(req.SeatID, req.UserID)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = "数据库操作失败"
		// TODO 补充错误日志
		return
	}
	return
}

func getSeatGetReq(r *http.Request) (req *SeatGetReq, err error) {
	req = new(SeatGetReq)
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
	req.UserID = userID.(int)
	return
}

func getSeatPostReq(r *http.Request) (req *SeatPostReq, err error) {
	req = new(SeatPostReq)
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
	seatID, ok := body["seatID"]
	if !ok {
		err = fmt.Errorf("缺少 seatID 参数")
		return
	}
	req.UserID = userID.(int)
	req.SeatID = seatID.(int)
	return
}
