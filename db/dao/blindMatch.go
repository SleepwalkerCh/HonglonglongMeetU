package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const BlindMatchTableName = "blind_match"

type BlindMatchModelInterface interface {
	GetBlindMatchHistoryByUserIDAndGender(userID int32, gender int) (blindMatchHistory []*model.BlindMatchModel, err error)
}

type BlindMatchModelInterfaceImp struct{}

var IBlindMatchInterface = &BlindMatchModelInterfaceImp{}

// GetBlindMatchHistoryByUserIDAndGender 按时间逆序获取某用户所用盲选记录
func (b *BlindMatchModelInterfaceImp) GetBlindMatchHistoryByUserIDAndGender(userID int32, gender int) (blindMatchHistory []*model.BlindMatchModel, err error) {
	cli := db.Get()
	blindMatchHistory = make([]*model.BlindMatchModel, 0)
	query := cli.Table(BlindMatchTableName)
	if gender == int(model.MaleGender) {
		query = query.Where("userid_male = ?", userID)
	} else {
		query = query.Where("userid_female = ?", userID)
	}
	if err = query.Order("created_at desc").Find(blindMatchHistory).Error; err != nil {
		return
	}
	return
}
