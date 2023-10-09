package postgresModel

import "time"

type UserPostgres struct {
	UserId       uint64    `db:"id"`
	Login        string    `db:"login"`
	Password     string    `db:"password"`
	Name         string    `db:"name"`
	Contact      string    `db:"contact"`
	RegisteredAt time.Time `db:"registered_at"`
	Role         string    `db:"role"`
}
