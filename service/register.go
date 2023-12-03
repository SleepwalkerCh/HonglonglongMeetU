package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/tools"
)

type RegisterReq struct {
	Gender     int    `json:"gender"`
	NickName   string `json:"nickName"`
	RealName   string `json:"realName"`
	InviteCode string `json:"inviteCode"`
}

func RegisterPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	//解析入参
	req, err := getRegisterReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	//入参校验
	err = validateRegisterReq(req)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	// 校验nickName不重复
	allUsers, err := dao.IUserInterface.GetAllUsers()
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		//fmt.Fprintf(w, "err:%v", err)
		return
	}
	allUserName := make([]string, 0)
	for _, user := range allUsers {
		allUserName = append(allUserName, user.NickName)
	}
	if tools.InList(req.NickName, allUserName) {
		res.Code = -1
		res.ErrorMsg = "昵称已存在"
		return
	}
	// 校验inviteCode
	userType := checkInviteCode(req.InviteCode)
	if userType == model.InvalidUserType {
		res.Code = -1
		res.ErrorMsg = "邀请码不合法"
		return
	}
	// 写入DB
	needInsertUser := &model.UserModel{
		NickName: req.NickName,
		RealName: req.RealName,
		Gender:   req.Gender,
		UserType: int(userType),
		Status:   model.NormalStatus,
	}
	err = dao.IUserInterface.InsertUser(needInsertUser)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = "数据库操作失败"
		// TODO 补充错误日志
		return
	}
	//实际绑定操作
	//TODO 调用微信接口绑定账号 1126

	res.Code = 0
	res.ErrorMsg = ""
	res.Data = "success"
	return
}

func getRegisterReq(r *http.Request) (req *RegisterReq, err error) {
	req = new(RegisterReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	gender, ok := body["gender"]
	if !ok {
		err = fmt.Errorf("缺少 gender 参数")
		return
	}
	nickName, ok := body["nickName"]
	if !ok {
		err = fmt.Errorf("缺少 nickName 参数")
		return
	}
	realName, ok := body["realName"]
	if !ok {
		err = fmt.Errorf("缺少 realName 参数")
		return
	}
	inviteCode, ok := body["inviteCode"]
	if !ok {
		err = fmt.Errorf("缺少 inviteCode 参数")
		return
	}
	req.Gender = gender.(int)
	req.NickName = nickName.(string)
	req.RealName = realName.(string)
	req.InviteCode = inviteCode.(string)
	return
}

func validateRegisterReq(req *RegisterReq) (err error) {
	// gender必须为0,1
	if tools.InList(req.Gender, []int{0, 1}) {
		err = fmt.Errorf("gender参数不合法")
		return
	}
	// 昵称不得大于12字符
	if len(req.NickName) > 12 {
		err = fmt.Errorf("昵称不得过长")
		return
	}
	// 真名不得大于10字符
	//if len(req.RealName)>10{
	//	err = fmt.Errorf("真名不得过长")
	//	return
	//}
	return
}

// checkInviteCode 先把嘉宾及管理员邀请码写死
func checkInviteCode(inviteCode string) int {
	//admin Base64后前五位
	if inviteCode == "YWRta" {
		return model.AdminUserType
	} else if inviteCode == "bm9yb" { //normal Base64后前五位
		return model.NormalUserType
	}
	return model.InvalidUserType
}
