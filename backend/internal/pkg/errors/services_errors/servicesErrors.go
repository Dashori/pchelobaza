package serviceErrors

import "errors"

var (
	// records
	ErrorCreateRecordTime = errors.New("Service error! Неверное время для записи!")
	TimeIsTaken           = errors.New("Service error! Данное время уже занято!")
	ErrorDoctorTime       = errors.New("Service error! Неверное время для записи к этому доктору!")

	NotUserPet = errors.New("Service error! Данный питомец не принадлежит вам!")

	// entity does not exists
	DoctorDoesNotExists = errors.New("Service error! Такого доктора не существует!")
	UserDoesNotExists   = errors.New("Service error! Такого пользователя не существует!")
	RecordDoesNotExists = errors.New("Service error! Такой записи не существует!")
	PetDoesNotExists    = errors.New("Service error! Такого питомца не существует!")

	// entity already exists
	PetAlreadyExists    = errors.New("Service error! Питомец уже существует в базе!")
	DoctorAlreadyExists = errors.New("Service error! Доктор уже существует в базе!")
	UserAlreadyExists   = errors.New("Service error! Пользователь уже существует в базе!")


	ErrorUserCreate = errors.New("Service error! Не удалось добавить нового пользователя!")
	ErrorUserUpdate      = errors.New("Service error! Не удалось обновить информацию о пользователе!")
	ErrorConfirmPassword = errors.New("Service error! Пароли не совпадают!")

	// Create + login
	ErrorGetUserByLogin   = errors.New("Service error! Ошибка при получении пользователя по логину!")


	ErrorHash             = errors.New("Service error! Ошибка получения хэша для пароля!")
	InvalidPassword       = errors.New("Service error! Неверный пароль!")

)
