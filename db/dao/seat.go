package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const SeatTableName = "seat"

type SeatModelInterface interface {
	GetSeatByUserID(userID int) ([]*model.SeatModel, error)
	GetSeatBySeatID(seatID int) ([]*model.SeatModel, error)
	GetAllSeats() ([]*model.SeatModel, error)
	UpdateSeatBySeatID(seatID, userID int) error
	RemoveSeatByUserID(userID int) (err error)
}

type SeatModelInterfaceImp struct{}

var ISeatInterface = &SeatModelInterfaceImp{}

func (s *SeatModelInterfaceImp) GetSeatByUserID(userID int) (seatInfo []*model.SeatModel, err error) {
	cli := db.Get()
	seatInfo = make([]*model.SeatModel, 0)
	if err = cli.Table(SeatTableName).Where("userid = ?", userID).Find(&seatInfo).Error; err != nil {
		return
	}
	return
}

func (s *SeatModelInterfaceImp) GetAllSeats() (seats []*model.SeatModel, err error) {
	cli := db.Get()
	seats = make([]*model.SeatModel, 0)
	if err = cli.Table(SeatTableName).Find(&seats).Error; err != nil {
		return
	}
	return
}

func (s *SeatModelInterfaceImp) GetSeatBySeatID(seatID int) (seatInfo []*model.SeatModel, err error) {
	cli := db.Get()
	seatInfo = make([]*model.SeatModel, 0)
	if err = cli.Table(SeatTableName).Where("id = ?", seatID).Find(&seatInfo).Error; err != nil {
		return
	}
	return
}

func (s *SeatModelInterfaceImp) UpdateSeatBySeatID(seatID, userID int) (err error) {
	cli := db.Get()
	if err = cli.Table(SeatTableName).Where("id = ?", seatID).Updates(map[string]interface{}{"userID": userID, "status": model.OccupiedStatus}).Error; err != nil {
		return
	}
	return
}

func (s *SeatModelInterfaceImp) RemoveSeatByUserID(userID int) (err error) {
	cli := db.Get()
	if err = cli.Table(SeatTableName).Where("userid = ?", userID).Updates(map[string]interface{}{"userID": 0, "status": model.FreeStatus}).Error; err != nil {
		return
	}
	return
}
