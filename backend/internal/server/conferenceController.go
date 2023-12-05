package server

import (
	models "backend/internal/models"
	"backend/internal/server/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *services) GetAllConferences(c *gin.Context) {
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

	res, err := s.Services.ConferenceService.GetAllConferences(limitNum, skippedNum)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonConferencesOkResponse(c, res)
}

func (s *services) GetConference(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
		return
	}

	res, err := s.Services.ConferenceService.GetConferenceByName(name)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonConferenceOkResponse(c, res)
}

func (s *services) GetConferenceUsers(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
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

	res, err := s.Services.ConferenceService.GetConferenceUsers(name, limitNum, skippedNum)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonConferenceUsersOkResponse(c, res)
}

func (s *services) GetConferenceReviews(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
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

	res, err := s.Services.ConferenceService.GetConferenceReviews(name, limitNum, skippedNum)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonConferenceReviewsOkResponse(c, res)
}

func (s *services) CreateConference(c *gin.Context) {
	var conference *models.Conference
	err := c.ShouldBindJSON(&conference)
	fmt.Println(conference)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	login, _, id, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	conference.UserLogin = login
	conference.UserId = id

	res, err := s.Services.ConferenceService.CreateConference(conference)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonConferenceCreateResponse(c, res)
}

func (s *services) PatchConference(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
		return
	}

	var conference *models.Conference
	err := c.ShouldBindJSON(&conference)
	fmt.Println(conference)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	login, _, id, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	conference.UserLogin = login
	conference.UserId = id

	err = s.Services.ConferenceService.PatchConference(conference)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonStatusOkResponse(c)
}

func (s *services) PatchConferenceUsers(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
		return
	}

	login, _, _, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	err = s.Services.ConferenceService.PatchConferenceUsers(name, login)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonConferenceUserCreateResponse(c)
}

func (s *services) AddReview(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		jsonBadRequestResponse(c, fmt.Errorf("No name in the path!"))
		return
	}

	var review *models.Review
	err := c.ShouldBindJSON(&review)
	fmt.Println(review)

	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	login, _, id, err := middlewares.ExtractTokenIdAndRole(c)
	if err != nil {
		jsonUnauthorizedResponse(c, nil)
		return
	}

	review.Login = login
	review.UserId = id
	review.ConferenceName = name

	res, err := s.Services.ConferenceService.CreateReview(review)
	if err != nil {
		fmt.Println(err)
	}

	if !errorHandler(c, err) {
		return
	}

	jsonReviewCreateResponse(c, res)
}
