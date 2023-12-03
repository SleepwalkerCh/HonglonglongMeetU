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
	if err = cli.Table(DateHistoryTableName).Where("roomid = ?", roomID).Where("status = ?", status).Order("created_at desc").Find(dateHistory).Error; err != nil {
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
