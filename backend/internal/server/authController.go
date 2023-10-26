package server

import (
	"backend/internal/models"
	"backend/internal/server/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *services) Login(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := s.Services.UserService.Login(user.Login, user.Password)

	if !errorHandler(c, err) {
		return
	}

	token, err := middlewares.GenerateToken(res.Login, res.UserId, res.Role)

	if err != nil {
		jsonBadRequestResponse(c, err)
		return
	}

	jsonUserLoginOkResponse(c, token)
}

func (s *services) SignUp(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	fmt.Println(user)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := s.Services.UserService.Create(user)
	if !errorHandler(c, err) {
		return
	}

	token, err := middlewares.GenerateToken(res.Login, res.UserId, res.Role)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	jsonUserCreateResponse(c, token)
}
