package serviceErrors

import "errors"

var (
	UserDoesNotExists     = errors.New("Service error! Такого пользователя не существует!")
	UserAlreadyExists     = errors.New("Service error! Пользователь уже существует в базе!")
	ErrorUserCreate       = errors.New("Service error! Не удалось добавить нового пользователя!")
	ErrorUserUpdate       = errors.New("Service error! Не удалось обновить информацию о пользователе!")
	ErrorConfirmPassword  = errors.New("Service error! Пароли не совпадают!")
	ErrorPaginationParams = errors.New("Service error! Неверные параметры для пагинации!")

	// Create + login
	ErrorGetUserByLogin = errors.New("Service error! Ошибка при получении пользователя по логину!")

	ErrorCreateRequest         = errors.New("Service error! Не удалось создать новую заявку!")
	ErrorGetAllRequests        = errors.New("Service error! Ошибка при получении пользовательских заявок!")
	ErrorGetRequestsPagination = errors.New("Service error! Ошибка при получении пользовательских заявок с пагинацией!")
	ErrorGetUserRequest        = errors.New("Service error! Ошибка при получении заявки пользователя!")
	UserAlreadyBeemaster       = errors.New("Service error! Пользователь уже является beemaster!")
	RequestDoesNotExists       = errors.New("Service error! Заявки от пользователя нет!")
	RequestAlreadyExists       = errors.New("Service error! Заявка от пользователя уже существует!")
	ErrorRequestStatus         = errors.New("Service error! Заявку нельзя редактировать, так как она не находится в статусе ожидания!")
	ErrorRequestPatch          = errors.New("Service error! Не удалось обновить статус по заявке!")

	ErrorCreateFarm    = errors.New("Service error! Не удалось создать новую ферму!")
	FarmDoesNotExists  = errors.New("Service error! Такой фермы не существует!")
	ErrorGetFarmByName = errors.New("Service error! Ошибка при получении фермы!")
	ErrorGetUsersFarm  = errors.New("Service error! Ошибка при получении ферм пользователя!")
	FarmAlreadyExists  = errors.New("Service error! Ферма с таким названием уже существует!")
	ErrorFarmAccess    = errors.New("Service error! Ферма не принадлежит Вам, вы не можете ее изменять!")
	ErrorFarmUpdate    = errors.New("Service error! Не удалось обновить информацию о ферме!")

	ErrorHash       = errors.New("Service error! Ошибка получения хэша для пароля!")
	InvalidPassword = errors.New("Service error! Неверный пароль!")
)
