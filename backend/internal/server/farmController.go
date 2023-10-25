package server

import (
	models "backend/internal/models"
	"backend/internal/server/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
		jsonBadRequestResponse(c, fmt.Errorf("No skipped in the query!"))
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

	login, _, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	farm.UserLogin = login

	res, err := s.Services.FarmService.CreateFarm(farm)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonFarmCreateResponse(c, *res)
}

func (s *services) GetFarmInfo(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No farm name in the query!"))
		return
	}

	res, err := s.Services.FarmService.GetFarm(name)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	honey, err := s.Services.HoneyService.GetFarmHoney(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("honey", honey)
	if !errorHandler(c, err) {
		return
	}
	res.Honey = honey

	jsonFarmInfoOkResponse(c, res)
}

func (s *services) PatchFarm(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		jsonBadRequestResponse(c, fmt.Errorf("No farm name in the query!"))
		return
	}

	login, _, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	var farm *models.Farm
	err = c.ShouldBindJSON(&farm)
	farm.Name = name
	farm.UserLogin = login

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	err = s.Services.FarmService.PatchFarm(farm)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonStatusOkResponse(c)
}
