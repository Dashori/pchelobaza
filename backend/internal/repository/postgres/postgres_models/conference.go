package postgresModel

import "time"

type ConferencePostgres struct {
	ConferenceId uint64    `db:"id"`
	UserId       uint64    `db:"id_user"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	Date         time.Time `db:"date"`
	Address      string    `db:"address"`
	MaxUsers     float64   `db:"maximum_users"`
	CurrentUsers float64   `db:"current_users"`
}

type ReviewPostgres struct {
	ReviewId     uint64    `db:"id"`
	ConferenceId uint64    `db:"id_conference"`
	UserId       uint64    `db:"id_user"`
	Login        string    `db:"login"`
	Name         string    `db:"name"`
	Surname      string    `db:"surname"`
	Date         time.Time `db:"date"`
	Description  string    `db:"description"`
}
