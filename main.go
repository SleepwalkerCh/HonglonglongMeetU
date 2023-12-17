package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/handler"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/count", handler.CounterHandler)
	// 1.(功能0)录入信息并提交 昵称+真名+性别+验证码(区分是否为嘉宾或工作人员) todo
	http.HandleFunc("/api/register", handler.RegisterHandler)
	// 2.(功能1)展示当前座位情况 Done
	// 3.(功能1)录入选定座位 Done
	// 4.(功能1)展示本人当前座位 Done
	http.HandleFunc("/api/allSeat", handler.AllSeatHandler)
	http.HandleFunc("/api/seat", handler.SeatHandler)
	// 5.(功能2)发起匹配 todo
	// 6.(功能2)匹配成功回调 含微信消息推送 todo
	// 7.(功能2)展示匹配成功信息页 todo
	http.HandleFunc("/api/blindMatch", handler.BlindMatchHandler)
	http.HandleFunc("/api/callback/blindMatch", handler.AllSeatHandler)
	// 8.(功能3)展示当前房间状态 Done
	// 9.(功能3)选择房间/座位 Done
	// 21.(功能3)退出约会房间（补） Done
	// 10.(功能3)确认约会 Done
	// 11.(功能3)确认结束约会 Done
	// 12.(功能3)提交约会结果 Done
	// 13.(功能3)约会结果回调 含微信消息推送 todo
	http.HandleFunc("/api/roomInfo", handler.RoomInfoHandler)
	http.HandleFunc("/api/dateRoomIn", handler.DateRoomInHandler)
	http.HandleFunc("/api/dateRoomOut", handler.DateRoomOutHandler)
	http.HandleFunc("/api/dateStart", handler.DateStartHandler)
	http.HandleFunc("/api/dateStop", handler.DateStopHandler)
	http.HandleFunc("/api/dateResultSubmit", handler.DateResultSubmitHandler)

	// 14.(功能4)展示所有用户当前状态 Done
	http.HandleFunc("/api/allUserStatus", handler.AllUserStatusHandler)

	// 15.(功能5)查看所有嘉宾信息和状态 Done
	// 16.(功能5)编辑所有嘉宾信息和状态 Done
	// 17.(功能5)查看约会历史 Done
	// 18.(功能5)编辑约会历史状态(待定) Done
	// 19.(功能5)查看时间段-匹配次数
	// 20.(功能5)编辑时间段-匹配次数
	http.HandleFunc("/api/userInfo", handler.UserInfoHandler)
	http.HandleFunc("/api/dateHistory", handler.DateHistoryHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
