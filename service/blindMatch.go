package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

type BlindMatchGetReq struct {
	UserID int `json:"userID"`
}

type BlindMatchRecord struct {
	ID             int    `json:"id"`
	UserIDMale     int    `json:"userid_male"`
	NicknameMale   string `json:"nickname_male"`
	UserIDFemale   int    `json:"userid_female"`
	NicknameFemale string `json:"nickname_female"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
type BlindMatchGetResp struct {
	BlindMatching     bool                `json:"blindMatching"`
	BlindMatchHistory []*BlindMatchRecord `json:"blindMatchHistory"`
}

type BlindMatchPostReq struct {
	UserID int `json:"userID"`
}
type BlindMatchPostResp struct {
	AvailableCnt int `json:"availableCnt"`
}

func BlindMatchGetFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	req, err := getBlindMatchGetReq(r)
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
	rawBlindMatchHistory, err := dao.IBlindMatchInterface.GetBlindMatchHistoryByUserIDAndGender(req.UserID, gender)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}

	blindMatchHistory, err := getBlindMatchHistoryFromRawData(rawBlindMatchHistory)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		return
	}
	res.Data = &BlindMatchGetResp{
		//BlindMatching:     BlindMatching,
		BlindMatchHistory: blindMatchHistory,
	}
	return
}

// TODO 抽象userID->userName
func getBlindMatchHistoryFromRawData(rawBlindMatchHistory []*model.BlindMatchModel) (blindMatchHistory []*BlindMatchRecord, err error) {
	userIDList := make([]int, 0)
	blindMatchHistory = make([]*BlindMatchRecord, 0)
	for _, blindMatch := range rawBlindMatchHistory {
		userIDList = append(userIDList, blindMatch.UserIDMale)
		userIDList = append(userIDList, blindMatch.UserIDFemale)
	}
	userMap, err := dao.IUserInterface.GetUsersByIDList(userIDList)
	if err != nil {
		return
	}
	for _, blindMatch := range rawBlindMatchHistory {
		blindMatchHistory = append(blindMatchHistory, &BlindMatchRecord{
			ID:             int(blindMatch.ID),
			UserIDMale:     int(blindMatch.UserIDMale),
			NicknameMale:   GetNickNameFromUserInfoMap(userMap, blindMatch.UserIDMale),
			UserIDFemale:   int(blindMatch.UserIDFemale),
			NicknameFemale: GetNickNameFromUserInfoMap(userMap, blindMatch.UserIDFemale),
			CreatedAt:      blindMatch.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      blindMatch.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return
}

// BlindMatchPostFunc TODO 先不考虑次数限制
func BlindMatchPostFunc(r *http.Request) (res *JsonResult) {
	res = &JsonResult{}
	res.Code = 0
	res.ErrorMsg = ""

	//解析入参
	req, err := getBlindMatchPostReq(r)
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
	// 时间点逻辑
	blindMatchHistory, err := dao.IBlindMatchInterface.GetBlindMatchHistoryByGenderAndTime(gender, time.Now().Add(-1*time.Hour).Format("2006-01-02T15:04:05Z07:00"))
	if len(blindMatchHistory) == 0 {
		// 未有异性待匹配中，需初始化一条记录
		err = dao.IBlindMatchInterface.CreateBlindMatchHistoryWithUserID(req.UserID, gender)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	} else {
		// 已有异性待匹配中，直接更改该条记录
		err = dao.IBlindMatchInterface.UpdateBlindMatchHistoryUserByID(blindMatchHistory[0].ID, req.UserID, gender)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
			return
		}
	}
	res.Data = &BlindMatchPostResp{AvailableCnt: 0}
	return
}

func getBlindMatchGetReq(r *http.Request) (req *BlindMatchGetReq, err error) {
	req = new(BlindMatchGetReq)
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

func getBlindMatchPostReq(r *http.Request) (req *BlindMatchPostReq, err error) {
	req = new(BlindMatchPostReq)
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
