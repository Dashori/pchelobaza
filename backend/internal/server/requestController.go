package server

import (
	models "backend/internal/models"
	"backend/internal/server/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *services) AddRequest(c *gin.Context) {
	var request *models.Request
	err := c.ShouldBindJSON(&request)
	fmt.Println(request)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	login, _, id, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	request.UserLogin = login
	request.UserId = id

	res, err := s.Services.RequestService.CreateRequest(request)
	if !errorHandler(c, err) {
		return
	}

	jsonRequestCreateResponse(c, res)
}

// login limit skipped
func (s *services) GetRequest(c *gin.Context) {
	login, role, _, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	loginFlag := true

	queryLogin, ok := c.GetQuery("login")
	if !ok {
		loginFlag = false
	}

	if loginFlag && login != queryLogin && role != "beeadmin" {
		jsonGetRequestPermResponse(c)
		return
	}

	if loginFlag {
		res, err := s.Services.RequestService.GetUserRequest(queryLogin)
		if !errorHandler(c, err) {
			return
		}
		jsonGetRequestResponse(c, res)
		return
	}

	limit, ok := c.GetQuery("limit")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No limit in the query!"))
		return
	}
	limitNum, _ := strconv.Atoi(limit)

	skipped, ok := c.GetQuery("skipped")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No skipped in the query!"))
		return
	}
	skippedNum, _ := strconv.Atoi(skipped)

	if role != "beeadmin" {
		jsonGetRequestPermResponse(c)
		return
	}

	if limitNum != 0 {
		res, err := s.Services.RequestService.GetRequestsPagination(limitNum, skippedNum)
		if !errorHandler(c, err) {
			return
		}
		jsonGetRequestsResponse(c, res)
		return
	}
}

func (s *services) PatchRequest(c *gin.Context) {
	_, role, id, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	if role != "beeadmin" {
		jsonGetRequestPermResponse(c)
		return
	}

	login, ok := c.GetQuery("login")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No login in the query!"))
		return
	}

	var request *models.Request
	err = c.ShouldBindJSON(&request)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	request.UserLogin = login
	request.UserId = id

	err = s.Services.RequestService.PatchUserRequest(*request)
	if !errorHandler(c, err) {
		return
	}

	jsonStatusOkResponse(c)
}
