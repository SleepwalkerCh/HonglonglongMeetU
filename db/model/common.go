package model

type Status int32

// User
const (
	AbnormalStatus Status = iota
	NormalStatus
)

// MatchSetting DateRoom
const (
	InactiveStatus Status = iota
	ActiveStatus
)

// Seat
const (
	FreeStatus Status = iota
	OccupiedStatus
)
