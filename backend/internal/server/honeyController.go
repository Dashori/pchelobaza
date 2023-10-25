package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *services) GetHoney(c *gin.Context) {
	res, err := s.Services.HoneyService.GetAllHoney()
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonUserHoneyOkResponse(c, res)
}
