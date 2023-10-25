package server

import (
	"github.com/gin-gonic/gin"

	dbErrors "backend/internal/pkg/errors/db_errors"
	servicesErrors "backend/internal/pkg/errors/services_errors"
)

func errorHandler(c *gin.Context, err error) bool {

	if err == nil {
		return true
	}

	if err == dbErrors.ErrorInitDB ||
		err == servicesErrors.ErrorHash ||
		err == servicesErrors.ErrorUserUpdate ||
		err == servicesErrors.ErrorUserCreate ||
		err == servicesErrors.ErrorGetUserByLogin ||
		err == servicesErrors.ErrorGetFarmByName{
		jsonInternalServerErrorResponse(c, err)
		return false
	}

	if err == servicesErrors.UserDoesNotExists ||
		err == servicesErrors.UserAlreadyExists ||
		err == servicesErrors.InvalidPassword ||
		err == servicesErrors.ErrorConfirmPassword ||
		err == servicesErrors.FarmAlreadyExists ||
		err == servicesErrors.ErrorFarmAccess || 
		err == servicesErrors.FarmDoesNotExists {
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
