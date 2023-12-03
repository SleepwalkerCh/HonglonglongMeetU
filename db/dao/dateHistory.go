package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const DateHistoryTableName = "date_history"

type DateHistoryModelInterface interface {
	CreateDateHistoryRecord(roomID, maleUserID, femaleUserID int) (err error)
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
