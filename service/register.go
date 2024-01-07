package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/tools"
)

type SignupReq struct {
	Gender     int    `json:"gender"`
	NickName   string `json:"nickName"`
	RealName   string `json:"realName"`
	InviteCode string `json:"inviteCode"`
	Code       string `json:"code"`
}

type LoginReq struct {
	Code string `json:"code"`
}

type LoginResp struct {
	UserID   int    `json:"userID"`
	NickName string `json:"nickName"`
	Gender   int    `json:"gender"`
}

func SignupPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	//解析入参
	req, err := getSignupReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	//入参校验
	err = validateSignupReq(req)
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
	_, openid, err := callLoginService(req.Code)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	// 写入DB
	needInsertUser := &model.UserModel{
		NickName: req.NickName,
		RealName: req.RealName,
		Gender:   req.Gender,
		UserType: userType,
		OpenID:   openid,
		Status:   model.NormalStatus,
	}
	err = dao.IUserInterface.InsertUser(needInsertUser)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = "数据库操作失败"
		return
	}
	res.Code = 0
	res.ErrorMsg = ""
	res.Data = "success"
	return
}

func getSignupReq(r *http.Request) (req *SignupReq, err error) {
	req = new(SignupReq)
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
	code, ok := body["code"]
	if !ok {
		err = fmt.Errorf("缺少 code 参数")
		return
	}
	req.Gender = int(gender.(float64))
	req.NickName = nickName.(string)
	req.RealName = realName.(string)
	req.InviteCode = inviteCode.(string)
	req.Code = code.(string)
	return
}

func validateSignupReq(req *SignupReq) (err error) {
	// gender必须为0,1
	if !tools.InList(req.Gender, []int{0, 1}) {
		err = fmt.Errorf("gender参数不合法")
		return
	}
	// 昵称不得大于12字符
	if len(req.NickName) > 12 {
		err = fmt.Errorf("昵称不得过长")
		return
	}
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

func LoginPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	//解析入参
	req, err := getLoginReq(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	_, openid, err := callLoginService(req.Code)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	user, err := dao.IUserInterface.GetUserByOpenID(openid)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	if user == nil {
		res.Code = 1 // 需要先注册
		res.ErrorMsg = "该用户需要注册"
		return
	}
	res.Code = 0
	res.ErrorMsg = ""
	res.Data = &LoginResp{
		UserID:   user.ID,
		NickName: user.NickName,
		Gender:   user.Gender,
	}
	return
}

func getLoginReq(r *http.Request) (req *LoginReq, err error) {
	req = new(LoginReq)
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err = decoder.Decode(&body); err != nil {
		return
	}
	defer r.Body.Close()

	code, ok := body["code"]
	if !ok {
		err = fmt.Errorf("缺少 code 参数")
		return
	}
	req.Code = code.(string)
	return
}

func callLoginService(code string) (sessionKey, openid string, err error) {
	// 创建查询参数
	baseURL := "https://api.weixin.qq.com/sns/jscode2session"
	params := url.Values{}

	params.Add("js_code", code)
	params.Add("appid", tools.AppID)
	params.Add("secret", tools.AppSecret)
	params.Add("grant_type", "authorization_code")

	// 将查询参数添加到基础URL
	reqURL := baseURL + "?" + params.Encode()
	resp, err := http.Get(reqURL) // 示例URL，可以替换为您要请求的URL
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取响应体内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	//{
	//	"openid":"xxxxxx",
	//	"session_key":"xxxxx",
	//	"unionid":"xxxxx",
	//	"errcode":0,
	//	"errmsg":"xxxxx"
	//}
	if errCode, ok := result["errcode"]; ok && errCode != 0 {
		errMsg := result["errmsg"]
		err = fmt.Errorf("[callLoginService]Code:%d-Msg:%s", errCode, errMsg)
		return
	}
	sessionKey = result["session_key"].(string)
	openid = result["openid"].(string)
	return
}
