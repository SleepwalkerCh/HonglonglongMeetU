package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const DateRoomTableName = "date_room"
const (
	InAction int = iota
	OutAction
)

type DateRoomModelInterface interface {
	GetAllDateRoom() (dateRooms []*model.DateRoomModel, err error)
	GetDateRoomByID(roomID int) (dateRooms []*model.DateRoomModel, err error)
	UpdateRoomWithUserIDAndGender(roomID, userID, gender int, action int) (err error)
}

type DateRoomModelInterfaceImp struct{}

var IDateRoomInterface = &DateRoomModelInterfaceImp{}

func (d *DateRoomModelInterfaceImp) GetAllDateRoom() (dateRooms []*model.DateRoomModel, err error) {
	cli := db.Get()
	dateRooms = make([]*model.DateRoomModel, 0)
	if err = cli.Table(DateRoomTableName).Find(dateRooms).Error; err != nil {
		return
	}
	return
}

func (d *DateRoomModelInterfaceImp) GetDateRoomByID(roomID int) (dateRooms []*model.DateRoomModel, err error) {
	cli := db.Get()
	dateRooms = make([]*model.DateRoomModel, 0)
	if err = cli.Table(DateRoomTableName).Where("id = ?", roomID).Find(dateRooms).Error; err != nil {
		return
	}
	return
}

// action = in 用户进入房间 action = out 用户出房间
func (d *DateRoomModelInterfaceImp) UpdateRoomWithUserIDAndGender(roomID, userID, gender int, action int) (err error) {
	cli := db.Get()
	needUpdateUserID := 0
	if action == InAction {
		needUpdateUserID = userID
	}
	query := cli.Table(DateRoomTableName).Where("id = ?", roomID)
	if gender == model.MaleGender {
		if err = query.Updates(map[string]interface{}{"userid_male": needUpdateUserID}).Error; err != nil {
			return
		}
	} else {
		if err = query.Updates(map[string]interface{}{"userid_female": needUpdateUserID}).Error; err != nil {
			return
		}
	}
	return
}
