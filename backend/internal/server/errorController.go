package server

import (
	"github.com/gin-gonic/gin"

	dbErrors "backend/internal/pkg/errors/db_errors"
	serviceErrors "backend/internal/pkg/errors/services_errors"
)

func errorHandler(c *gin.Context, err error) bool {

	if err == nil {
		return true
	}

	if err == dbErrors.ErrorInitDB ||
		err == serviceErrors.ErrorHash ||
		err == serviceErrors.ErrorUserUpdate ||
		err == serviceErrors.ErrorUserCreate ||
		err == serviceErrors.ErrorGetUserByLogin ||
		err == serviceErrors.ErrorGetFarmByName ||
		err == serviceErrors.ErrorCreateFarm ||
		err == serviceErrors.ErrorGetHoney ||
		err == serviceErrors.ErrorCreateRequest {
		jsonInternalServerErrorResponse(c, err)
		return false
	}

	if err == serviceErrors.UserDoesNotExists ||
		err == serviceErrors.UserAlreadyExists ||
		err == serviceErrors.InvalidPassword ||
		err == serviceErrors.ErrorConfirmPassword ||
		err == serviceErrors.FarmAlreadyExists ||
		err == serviceErrors.ErrorFarmAccess ||
		err == serviceErrors.FarmDoesNotExists ||
		err == serviceErrors.ErrorHoney ||
		err == serviceErrors.RequestAlreadyExists ||
		err == serviceErrors.UserAlreadyBeemaster {
		jsonBadRequestResponse(c, err)
		return false
	}

	return true
}

func errorHandlerClientAuth(c *gin.Context, err error, role string) bool {

	if err != nil {
		jsonUnauthorizedResponse(c, err)
		return false
	}

	if role != "client" {
		jsonBadRoleResponse(c, role)
		return false
	}

	return true
}

func errorHandlerDoctorAuth(c *gin.Context, err error, role string) bool {

	if err != nil {
		jsonUnauthorizedResponse(c, err)
		return false
	}

	if role != "doctor" {
		jsonBadRoleResponse(c, role)
		return false
	}

	return true
}
