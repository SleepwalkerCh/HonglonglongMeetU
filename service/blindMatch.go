package service

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	rawBlindMatchHistory, err := dao.IBlindMatchInterface.GetBlindMatchHistoryByUserIDAndGender(req.UserID, int(gender))
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
		BlindMatching:     BlindMatching,
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

// TODO
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
	if len(seat) != 0 && seat[0].Status != model.FreeStatus) {
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
