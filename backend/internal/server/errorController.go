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

	if err == serviceErrors.FarmAlreadyExists ||
		err == serviceErrors.RequestAlreadyExists ||
		err == serviceErrors.UserAlreadyExists {
		jsonAlreadyExistsResponse(c, err)
		return false
	}

	if err == dbErrors.ErrorInitDB ||
		err == serviceErrors.ErrorHash ||
		err == serviceErrors.ErrorUserUpdate ||
		err == serviceErrors.ErrorUserCreate ||
		err == serviceErrors.ErrorGetUserByLogin ||
		err == serviceErrors.ErrorGetFarmByName ||
		err == serviceErrors.ErrorCreateFarm ||
		err == serviceErrors.ErrorGetHoney ||
		err == serviceErrors.ErrorCreateRequest ||
		err == serviceErrors.ErrorGetConferencesPagination ||
		err == serviceErrors.ErrorCreateConference ||
		err == serviceErrors.ErrorGetConference ||
		err == serviceErrors.ErrorEditConference ||
		err == serviceErrors.ErrorGetConferenceUsers ||
		err == serviceErrors.ErrorJoinConf ||
		err == serviceErrors.ErrorGetConferenceReviews ||
		err == serviceErrors.ErrorPatchConfUsers {
		jsonInternalServerErrorResponse(c, err)
		return false
	}

	if err == serviceErrors.UserDoesNotExists ||
		err == serviceErrors.InvalidPassword ||
		err == serviceErrors.ErrorConfirmPassword ||
		err == serviceErrors.ErrorFarmAccess ||
		err == serviceErrors.FarmDoesNotExists ||
		err == serviceErrors.ErrorHoney ||
		err == serviceErrors.UserAlreadyBeemaster ||
		err == serviceErrors.ErrorPaginationParams ||
		err == serviceErrors.ErrorRoleForConference ||
		err == serviceErrors.ErrorDateForConference ||
		err == serviceErrors.ErrorUsersForConference ||
		err == serviceErrors.ErrorNameForConference ||
		err == serviceErrors.ErrorNoConference ||
		err == serviceErrors.ErrorOldConference ||
		err == serviceErrors.ErrorDateForConference ||
		err == serviceErrors.ErrorUsersForConference ||
		err == serviceErrors.ErrorNoPlace ||
		err == serviceErrors.ErrorConferenceJoin ||
		err == serviceErrors.RequestErrorValue {
		jsonBadRequestResponse(c, err)
		return false
	}

	if err == serviceErrors.ErrorNoYourConference ||
		err == serviceErrors.ErrorRoleForConference {
		jsonPermResponse(c, err)
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
