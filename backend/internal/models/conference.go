package models

import "time"

type Conference struct {
	ConferenceId uint64
	Name         string
	UserId       uint64
	UserLogin    string
	Description  string
	Date         time.Time
	Address      string
	MaxUsers     int
	CurrentUsers int
}

type Review struct {
	ReviewId       uint64
	ConferenceId   uint64
	ConferenceName string
	UserId         uint64
	Login          string
	Name           string
	Surname        string
	Date           time.Time
	Description    string
}
