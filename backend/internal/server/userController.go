package server

import (
	models "backend/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

	tempId, ok := c.GetQuery("id")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No id in the query!"))
		return
	}

	id, err := strconv.Atoi(tempId)
	if err != nil {
		jsonBadRequestResponse(c, fmt.Errorf("Error id in the query!"))
		return
	}

	var user *models.User
	err = c.ShouldBindJSON(&user)
	// user.Login = login
	user.UserId = uint64(id)
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
