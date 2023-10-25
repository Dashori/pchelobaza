package serviceErrors

import "errors"

var (
	// bad request
	UserDoesNotExists    = errors.New("Service error! Такого пользователя не существует!")
	InvalidPassword      = errors.New("Service error! Неверный пароль!")
	ErrorConfirmPassword = errors.New("Service error! Пароли не совпадают!")

	UserAlreadyExists = errors.New("Service error! Пользователь уже существует в базе!")
	ErrorUserCreate   = errors.New("Service error! Не удалось добавить нового пользователя!")
	ErrorUserUpdate   = errors.New("Service error! Не удалось обновить информацию о пользователе!")

	ErrorPaginationParams = errors.New("Service error! Неверные параметры для пагинации!")

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

	ErrorHoney    = errors.New("Service error! Такого меда не существует!")
	ErrorGetHoney = errors.New("Service error! Ошибка при получении меда")

	ErrorGetConferencesPagination = errors.New("Service error! Ошибка при получении конференций с пагинацией!")
	ErrorRoleForConference        = errors.New("Service error! Ошибка при создании конференции -- пользователь не является мастером!")
	ErrorDateForConference        = errors.New("Service error! Ошибка при создании конференции -- нельзя создать конференцию на прощедшую дату!")
	ErrorUsersForConference       = errors.New("Service error! Ошибка при создании конференции -- минимальное количество мест для участников 10!")
	ErrorNameForConference        = errors.New("Service error! Ошибка при создании конференции -- конференция с таким названием уже существует!")
	ErrorCreateConference         = errors.New("Service error! Не удалось создать новую конференцию!")
	ErrorGetConference            = errors.New("Service error! Ошибка при получении конференции!")
	ErrorNoConference             = errors.New("Service error! Конференции с таким названием нет!")
	ErrorNoYourConference         = errors.New("Service error! Конференция не принадлежит Вам, Вы не можете ее редактировать!")
	ErrorOldConference            = errors.New("Service error! Ошибка при изменении конференции, эта конференция уже прошла!")
	ErrorEditConference           = errors.New("Service error! Ошибка при изменении конференции!")
	ErrorGetConferenceUsers       = errors.New("Service error! Ошибка при получении пользователей конференции!")
	ErrorGetConferenceReviews     = errors.New("Service error! Ошибка при получении комментариев конференции!")
	ErrorConferenceJoin           = errors.New("Service error! Данный пользователь уже зарегестрирован на конференцию!")
	ErrorNoPlace                  = errors.New("Service error! Все места на конференцию заняты!")
	ErrorJoinConf                 = errors.New("Service error! Ошибка при записи на конференцию!")
	ErrorCreateReview             = errors.New("Service error! Ошибка при добавлении комментария!")

	ErrorHash = errors.New("Service error! Ошибка получения хэша для пароля!")
)
