package dbErrors

import "errors"

var (
	ErrorParseConfig = errors.New("Ошибка при чтении конфига!")
	ErrorInitDB      = errors.New("Ошибка подключения к базе данных!")

	ErrorSetRole = errors.New("Ошибка базы данных! Невозможно установить роль!")

	ErrorInsert = errors.New("Ошибка базы данных! Невалидная операция insert!")
	ErrorDelete = errors.New("Ошибка базы данных! Невалидная операция delete!")
	ErrorSelect = errors.New("Ошибка базы данных! Невалидная операция select")
	ErrorUpdate = errors.New("Ошибка базы данных! Невалидная операция update")

	ErrorCopy = errors.New("Ошибка при копировании полей из базы данных!")
)
