package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const DateHistoryTableName = "date_history"

type DateHistoryModelInterface interface {
	CreateDateHistoryRecord(roomID, maleUserID, femaleUserID int) (err error)
	GetDateHistoryByRoomIDAndStatus(roomID int, status int) (dateHistory []*model.DateHistoryModel, err error)
	UpdateDateHistoryStatus(roomID, status int) (err error)
	UpdateDateHistoryByID(roomID int, updateMap map[string]interface{}) (err error)
	GetDateHistoryByUserIDAndGender(userID, gender int) (dateHistory []*model.DateHistoryModel, err error)
	UpdateDateHistoryResultByIDAndGender(ID, gender, result int) (err error)
	GetAllDatingDateHistory() (dateHistoryList []*model.DateHistoryModel, err error)
	GetAllDateHistory() (dateHistory []*model.DateHistoryModel, err error)
}

type DateHistoryModelInterfaceImp struct{}

var IDateHistoryInterface = &DateHistoryModelInterfaceImp{}

func (d DateHistoryModelInterfaceImp) CreateDateHistoryRecord(roomID, maleUserID, femaleUserID int) (err error) {
	cli := db.Get()
	dateHistory := &model.DateHistoryModel{
		UserIDMale:   maleUserID,
		UserIDFemale: femaleUserID,
		RoomID:       roomID,
		ResultMale:   model.InitResult,
		ResultFemale: model.InitResult,
		Status:       model.DatingStatus,
	}
	if err = cli.Table(DateHistoryTableName).Save(dateHistory).Error; err != nil {
		return
	}
	return
}

func (d DateHistoryModelInterfaceImp) GetDateHistoryByRoomIDAndStatus(roomID int, status int) (dateHistory []*model.DateHistoryModel, err error) {
	cli := db.Get()
	dateHistory = make([]*model.DateHistoryModel, 0)
	if err = cli.Table(DateHistoryTableName).Where("roomid = ?", roomID).Where("status = ?", status).Order("created_at desc").Find(&dateHistory).Error; err != nil {
		return
	}
	return
}

func (d DateHistoryModelInterfaceImp) UpdateDateHistoryStatus(roomID, status int) (err error) {
	cli := db.Get()
	if err = cli.Table(DateHistoryTableName).Where("roomid = ?", roomID).Updates(map[string]interface{}{"status": status}).Error; err != nil {
		return
	}
	return
}

func (d DateHistoryModelInterfaceImp) GetDateHistoryByUserIDAndGender(userID, gender int) (dateHistory []*model.DateHistoryModel, err error) {
	cli := db.Get()
	dateHistory = make([]*model.DateHistoryModel, 0)
	query := cli.Table(DateHistoryTableName).Where("status = ?", model.FinishedStatus)
	if gender == model.MaleGender {
		query = query.Where("userid_male = ?", userID)
	} else {
		query = query.Where("userid_female = ?", userID)
	}
	if err = query.Order("created_at desc").Find(&dateHistory).Error; err != nil {
		return
	}
	return
}

func (d DateHistoryModelInterfaceImp) UpdateDateHistoryResultByIDAndGender(ID, gender, result int) (err error) {
	cli := db.Get()
	query := cli.Table(DateHistoryTableName).Where("id = ?", ID)
	if gender == model.MaleGender {
		if err = query.Updates(map[string]interface{}{"result_male": result}).Error; err != nil {
			return
		}
	} else {
		if err = query.Updates(map[string]interface{}{"result_female": result}).Error; err != nil {
			return
		}
	}
	return
}

func (d DateHistoryModelInterfaceImp) GetAllDatingDateHistory() (dateHistory []*model.DateHistoryModel, err error) {
	cli := db.Get()
	dateHistory = make([]*model.DateHistoryModel, 0)
	if err = cli.Table(DateHistoryTableName).Where("status = ?", model.DatingStatus).Order("created_at desc").Find(&dateHistory).Error; err != nil {
		return
	}
	return
}

func (d DateHistoryModelInterfaceImp) GetAllDateHistory() (dateHistory []*model.DateHistoryModel, err error) {
	cli := db.Get()
	dateHistory = make([]*model.DateHistoryModel, 0)
	if err = cli.Table(DateHistoryTableName).Order("created_at desc").Find(&dateHistory).Error; err != nil {
		return
	}
	return
}

func (d DateHistoryModelInterfaceImp) UpdateDateHistoryByID(roomID int, updateMap map[string]interface{}) (err error) {
	cli := db.Get()
	if err = cli.Table(DateHistoryTableName).Where("id = ?", roomID).Updates(updateMap).Error; err != nil {
		return
	}
	return
}
