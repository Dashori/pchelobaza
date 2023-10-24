package repoErrors

import "errors"

var (
	EntityDoesNotExists = errors.New("Repository error! Такой сущности нет в базе данных!")
)
