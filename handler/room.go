package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/service"
)

func RoomInfoHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodGet {
		res = service.RoomInfoGetFunc(r)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "内部错误,err:%v", err.Error())
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func DateRoomInHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.DateRoomInPostFunc(r)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "内部错误,err:%v", err.Error())
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func DateRoomOutHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.DateRoomOutPostFunc(r)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "内部错误,err:%v", err.Error())
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
