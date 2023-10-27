package postgresModel

type RequestRostgres struct {
	RequestId   uint64 `db:"id"`
	UserId      uint64 `db:"id_user"`
	UserLogin   string `db:"login"`
	Description string `db:"description"`
	Status      string `db:"status"`
}
