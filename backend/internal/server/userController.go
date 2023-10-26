package server

import (
	models "backend/internal/models"
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

	var user *models.User
	err := c.ShouldBindJSON(&user)
	user.Login = login
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
