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
		err == serviceErrors.ErrorGetAllRequests ||
		err == serviceErrors.ErrorHash ||
		err == serviceErrors.ErrorUserUpdate ||
		err == serviceErrors.ErrorUserCreate ||
		err == serviceErrors.ErrorGetUserByLogin ||
		err == serviceErrors.ErrorGetUserById ||
		err == serviceErrors.ErrorGetFarmByName ||
		err == serviceErrors.ErrorCreateFarm ||
		err == serviceErrors.ErrorGetHoney ||
		err == serviceErrors.ErrorCreateRequest ||
		err == serviceErrors.ErrorGetConferencesPagination ||
		err == serviceErrors.ErrorGetRequestsPagination ||
		err == serviceErrors.ErrorGetUserRequest ||
		err == serviceErrors.ErrorCreateConference ||
		err == serviceErrors.ErrorGetConference ||
		err == serviceErrors.ErrorEditConference ||
		err == serviceErrors.ErrorGetConferenceUsers ||
		err == serviceErrors.ErrorJoinConf ||
		err == serviceErrors.ErrorGetConferenceReviews ||
		err == serviceErrors.ErrorPatchConfUsers ||
		err == serviceErrors.ErrorRequestPatch ||
		err == serviceErrors.ErrorGetUsersFarm ||
		err == serviceErrors.ErrorFarmUpdate ||
		err == serviceErrors.ErrorCreateReview {
		jsonInternalServerErrorResponse(c, err)
		return false
	}

	if err == serviceErrors.UserDoesNotExists ||
		err == serviceErrors.InvalidPassword ||
		err == serviceErrors.ErrorConfirmPassword ||
		err == serviceErrors.FarmDoesNotExists ||
		err == serviceErrors.ErrorHoney ||
		err == serviceErrors.UserAlreadyBeemaster ||
		err == serviceErrors.ErrorPaginationParams ||
		err == serviceErrors.ErrorDateForConference ||
		err == serviceErrors.ErrorNameForConference ||
		err == serviceErrors.ErrorNoConference ||
		err == serviceErrors.ErrorOldConference ||
		err == serviceErrors.ErrorUsersForConference ||
		err == serviceErrors.ErrorNoPlace ||
		err == serviceErrors.ErrorConferenceJoin ||
		err == serviceErrors.RequestDoesNotExists ||
		err == serviceErrors.RequestErrorValue ||
		err == serviceErrors.ErrorRequestStatus {
		jsonBadRequestResponse(c, err)
		return false
	}

	if err == serviceErrors.ErrorNoYourConference ||
		err == serviceErrors.ErrorRoleForConference ||
		err == serviceErrors.ErrorFarmAccess {
		jsonPermResponse(c, err)
	}

	return true
}
