package server

import (
	models "backend/internal/models"
	"backend/internal/server/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *services) GetUser(c *gin.Context) {
	login := c.Param("login")
	if login == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
		return
	}

	res, err := s.Services.UserService.GetUserByLogin(login)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonUserInfoOkResponse(c, res)
}

func (s *services) PatchUser(c *gin.Context) {
	login := c.Param("login")
	if login == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
		return
	}

	loginToken, _, id, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	if login != loginToken {
		jsonBadRequestResponse(c, fmt.Errorf("Login from query and token not the same!"))
		return
	}

	var user *models.User
	err = c.ShouldBindJSON(&user)
	user.UserId = id
	fmt.Println(user)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	err = s.Services.UserService.Update(user)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonStatusOkResponse(c)
}
