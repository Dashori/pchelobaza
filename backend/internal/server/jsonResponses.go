package server

import (
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonStatusOkResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func jsonPermResponse(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
}

func jsonAlreadyExistsResponse(c *gin.Context, err error) {
	c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
}

// errors

func jsonInternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func jsonBadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func jsonUnauthorizedResponse(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}

// user
func jsonUserInfoOkResponse(c *gin.Context, user *models.User) {
	c.JSON(http.StatusOK, gin.H{"userId": user.UserId, "login": user.Login, "name": user.Name, "surname": user.Surname,
		"contact": user.Contact, "registration_date": user.RegisteredAt, "role": user.Role})
}

func jsonUserCreateResponse(c *gin.Context, token string) {
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func jsonUserLoginOkResponse(c *gin.Context, token string) {
	c.JSON(http.StatusOK, gin.H{"Token": token})
}

// farm
func jsonUserFarmsOkResponse(c *gin.Context, farms []models.Farm) {
	c.JSON(http.StatusOK, gin.H{"farms": farms})
}

func jsonFarmInfoOkResponse(c *gin.Context, farm *models.Farm) {
	c.JSON(http.StatusOK, gin.H{"farmId": farm.FarmId, "name": farm.Name, "description": farm.Description,
		"address": farm.Address, "user": farm.UserLogin, "honey": farm.Honey})
}

func jsonFarmCreateResponse(c *gin.Context, farm models.Farm) {
	c.JSON(http.StatusCreated, gin.H{"farmId": farm.FarmId, "name": farm.Name, "description": farm.Description,
		"address": farm.Address, "userLogin": farm.UserLogin})
}

// honey
func jsonUserHoneyOkResponse(c *gin.Context, honey []models.Honey) {
	c.JSON(http.StatusOK, gin.H{"honey": honey})
}

// request
func jsonRequestCreateResponse(c *gin.Context, request *models.Request) {
	c.JSON(http.StatusCreated, gin.H{"description": request.Description, "status": request.Status})
}

func jsonGetRequestResponse(c *gin.Context, request *models.Request) {
	c.JSON(http.StatusOK, gin.H{"description": request.Description, "status": request.Status})
}

func jsonGetRequestsResponse(c *gin.Context, requests []models.Request) {
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func jsonGetRequestPermResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"err": "You can't see this, permission denied!"})
}

// conferences
func jsonConferencesOkResponse(c *gin.Context, conferences []models.Conference) {
	c.JSON(http.StatusOK, gin.H{"conferences": conferences})
}

func jsonConferenceOkResponse(c *gin.Context, conference *models.Conference) {
	c.JSON(http.StatusOK, gin.H{"name": conference.Name, "userLogin": conference.UserLogin,
		"description": conference.Description, "date": conference.Date,
		"address": conference.Address, "maxUsers": conference.MaxUsers,
		"currentUsers": conference.CurrentUsers})
}

func jsonConferenceUsersOkResponse(c *gin.Context, users []models.User) {
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func jsonConferenceReviewsOkResponse(c *gin.Context, reviews []models.Review) {
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func jsonConferenceCreateResponse(c *gin.Context, conference *models.Conference) {
	c.JSON(http.StatusCreated, gin.H{"name": conference.Name, "userLogin": conference.UserLogin,
		"description": conference.Description, "date": conference.Date,
		"address": conference.Address, "maxUsers": conference.MaxUsers,
		"currentUsers": conference.CurrentUsers})
}

func jsonConferenceUserCreateResponse(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}

func jsonReviewCreateResponse(c *gin.Context, review *models.Review) {
	c.JSON(http.StatusCreated, gin.H{"conferenceName": review.ConferenceName, "login": review.Login,
		"name": review.Name, "surname": review.Surname, "date": review.Date, "description": review.Description})
}
