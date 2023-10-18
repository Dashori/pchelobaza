package repoErrors

import "errors"

var (
	EntityDoesNotExists = errors.New("Repository error! Такой сущности нет в базе данных!")

	ErrorGetClientByLogin = errors.New("Repository error! Ошибка при получении клиента по логину")
	ErrorGetDoctorByLogin = errors.New("Repository error! Ошибка при получении доктора по логину")

	ErrorGetAllByClient = errors.New("Repository error! Ошибка при получении всех питомцев клиента!")
	ErrorGetAllByDoctor = errors.New("Repository error! Ошибка при получении всех записей доктора")
)
