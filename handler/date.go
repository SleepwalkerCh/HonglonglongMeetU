package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/service"
)

func DateStartHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.DateStartPostFunc(r)
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

// DateStopHandler 仅更新dateHistory状态,需配合用户离开约会房间使用
func DateStopHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.DateStopPostFunc(r)
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

func DateResultSubmitHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.DateResultSubmitPostFunc(r)
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
