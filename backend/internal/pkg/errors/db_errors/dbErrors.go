package dbErrors

import "errors"

var (
	ErrorInitDB = errors.New("Ошибка подключения к базе данных!")
)
