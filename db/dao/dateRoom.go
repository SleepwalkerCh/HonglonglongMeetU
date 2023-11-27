package dao

const DateRoomTableName = "date_room"

type DateRoomModelInterface interface {
}

type DateRoomModelInterfaceImp struct{}

var IDateRoomInterface = &DateRoomModelInterfaceImp{}
