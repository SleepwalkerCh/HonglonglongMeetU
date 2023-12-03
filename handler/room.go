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
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func DateRoomHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.DateRoomPostFunc(r)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
