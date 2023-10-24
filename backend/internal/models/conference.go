package models

import "time"

type Conference struct {
	ConferenceId uint64
	Name         string
	UserId       uint64
	Description  string
	Date         time.Time
	Address      string
	MaxUsers     float64
	CurrentUsers float64
}

type Review struct {
	ReviewId     uint64
	ConferenceId uint64
	UserId       uint64
	Login        string
	Name         string
	Surname      string
	Date         time.Time
	Description  string
}
