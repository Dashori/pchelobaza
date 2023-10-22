package postgresModel

type HoneyPostgres struct {
	HoneyId     uint64 `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
