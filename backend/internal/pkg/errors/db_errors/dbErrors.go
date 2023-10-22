package dbErrors

import "errors"

var (
	ErrorParseConfig = errors.New("Ошибка при чтении конфига!")
	ErrorInitDB      = errors.New("Ошибка подключения к базе данных!")

	ErrorSetRole = errors.New("Ошибка базы данных! Невозможно установить роль!")
)
