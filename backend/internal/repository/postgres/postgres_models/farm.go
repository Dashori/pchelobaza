package postgresModel

type FarmPostgres struct {
	FarmId      uint64 `db:"id"`
	UserId      uint64 `db:"id_user"`
	UserLogin   string `db:"login"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Address     string `db:"address"`
}
