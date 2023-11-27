package dao

const DateHistoryTableName = "date_history"

type DateHistoryModelInterface interface {
}

type DateHistoryModelInterfaceImp struct{}

var IDateHistoryInterface = &DateHistoryModelInterfaceImp{}
