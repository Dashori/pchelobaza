package server

import (
	models "backend/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	// "net/http"
)

func (s *services) GetFarms(c *gin.Context) {
	login, ok := c.GetQuery("login")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No login in the query!"))
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
		jsonBadRequestResponse(c, fmt.Errorf("No login in the query!"))
		return
	}
	skippedNum, _ := strconv.Atoi(skipped)

	res, err := s.Services.FarmService.GetUsersFarm(login, limitNum, skippedNum)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonUserFarmsOkResponse(c, res)
}

func (s *services) AddFarm(c *gin.Context) {
	var farm *models.Farm
	err := c.ShouldBindJSON(&farm)
	fmt.Println(farm)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	res, err := s.Services.FarmService.CreateFarm(farm)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonFarmCreateResponse(c, *res)
}
