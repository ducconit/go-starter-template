package models

type UserStatus int16

// User status
const (
	UserStatusBanned UserStatus = iota
	UserStatusInactive
	UserStatusActive
)
