package dao

const MatchSettingTableName = "match_setting"

type MatchSettingModelInterface interface {
}

type MatchSettingModelInterfaceImp struct{}

var IMatchSettingInterface = &MatchSettingModelInterfaceImp{}
