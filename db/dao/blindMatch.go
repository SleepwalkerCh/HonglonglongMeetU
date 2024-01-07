package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const BlindMatchTableName = "blind_match"

type BlindMatchModelInterface interface {
	GetBlindMatchHistoryByUserIDAndGender(userID int, gender int) (blindMatchHistory []*model.BlindMatchModel, err error)
	GetBlindMatchHistoryByGenderAndTime(gender int, time string) (blindMatchHistory []*model.BlindMatchModel, err error)
	UpdateBlindMatchHistoryUserByID(ID, userID, gender int) (err error)
	CreateBlindMatchHistoryWithUserID(userID, gender int) (err error)
}

type BlindMatchModelInterfaceImp struct{}

var IBlindMatchInterface = &BlindMatchModelInterfaceImp{}

// GetBlindMatchHistoryByUserIDAndGender 按时间逆序获取某用户所用盲选记录
func (b *BlindMatchModelInterfaceImp) GetBlindMatchHistoryByUserIDAndGender(userID int, gender int) (blindMatchHistory []*model.BlindMatchModel, err error) {
	cli := db.Get()
	blindMatchHistory = make([]*model.BlindMatchModel, 0)
	query := cli.Table(BlindMatchTableName)
	if gender == model.MaleGender {
		query = query.Where("userid_male = ?", userID)
	} else {
		query = query.Where("userid_female = ?", userID)
	}
	if err = query.Order("created_at desc").Find(&blindMatchHistory).Error; err != nil {
		return
	}
	return
}

func (b *BlindMatchModelInterfaceImp) GetBlindMatchHistoryByGenderAndTime(gender int, time string) (blindMatchHistory []*model.BlindMatchModel, err error) {
	cli := db.Get()
	blindMatchHistory = make([]*model.BlindMatchModel, 0)
	query := cli.Table(BlindMatchTableName)
	if gender == model.MaleGender {
		query = query.Where("userid_male = ?", model.InitUserID)
	} else {
		query = query.Where("userid_female = ?", model.InitUserID)
	}
	if err = query.Where("created_at > ?", time).Find(&blindMatchHistory).Error; err != nil {
		return
	}
	return
}

func (b *BlindMatchModelInterfaceImp) UpdateBlindMatchHistoryUserByID(ID, userID, gender int) (err error) {
	cli := db.Get()
	query := cli.Table(BlindMatchTableName).Where("id = ?", ID)
	if gender == model.MaleGender {
		if err = query.Updates(map[string]interface{}{"userid_male": userID}).Error; err != nil {
			return
		}
	} else {
		if err = query.Updates(map[string]interface{}{"userid_female": userID}).Error; err != nil {
			return
		}
	}
	return
}

func (b *BlindMatchModelInterfaceImp) CreateBlindMatchHistoryWithUserID(userID, gender int) (err error) {
	cli := db.Get()
	blindMatchHistory := &model.BlindMatchModel{
		UserIDMale:   model.InitUserID,
		UserIDFemale: model.InitUserID,
	}
	if gender == model.MaleGender {
		blindMatchHistory.UserIDMale = userID
	} else {
		blindMatchHistory.UserIDFemale = userID
	}
	if err = cli.Table(DateHistoryTableName).Save(blindMatchHistory).Error; err != nil {
		return
	}
	return
}
