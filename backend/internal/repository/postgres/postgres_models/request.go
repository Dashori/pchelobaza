package postgresModel

type RequestRostgres struct {
	RequestId   uint64 `db:"id"`
	UserLogin   string `db:"login"`
	Description string `db:"description"`
	Status      string `db:"status"`
}
