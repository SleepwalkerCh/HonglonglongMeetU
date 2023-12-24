package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/service"
)

// SignupHandler Request
//
//	{
//	   "gender":0,//0-男 1-女
//	   "nickName":"追风少年",
//	   "realName":"张三",
//	   "inviteCode":"123456"//邀请码，确认为嘉宾还是工作人员
//	}
//
// SignupHandler 注册接口
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.SignupPostFunc(r)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误,json序列化失败")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	res := &service.JsonResult{}

	if r.Method == http.MethodPost {
		res = service.LoginPostFunc(r)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误,json序列化失败")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
