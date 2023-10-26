package server

import (
	"backend/internal/app"
	"backend/internal/server/middlewares"
	"encoding/json"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

type services struct {
	Services *app.AppServiceFields
}

func SetupServer(a *app.App) *gin.Engine {
	s := services{a.Services}

	router := gin.Default()

	swagger := ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/docs/swagger.yaml"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.GET("swagger/*any", swagger)
	router.GET("docs/*any", getOpenApi)

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", s.Login)
			auth.POST("/signup", s.SignUp)
			// auth.POST("/logout", s.DeleteFileContent)
		}

		user := api.Group("/users")
		{
			user.Use(middlewares.JwtAuthMiddleware())
			user.GET("", s.GetUser)
			user.PATCH("", s.PatchUser)
		}

		farm := api.Group("/farms")
		{
			farm.Use(middlewares.JwtAuthMiddleware())
			farm.GET("", s.GetFarms)
			farm.POST("", s.AddFarm)
			farm.GET("/", s.GetFarmInfo)
			farm.PATCH("/", s.PatchFarm)
		}

		honey := api.Group("/honey")
		{
			honey.Use(middlewares.JwtAuthMiddleware())
			honey.GET("", s.GetHoney)
		}

		request := api.Group("/requests")
		{
			request.Use(middlewares.JwtAuthMiddleware())
			request.POST("", s.AddRequest)
			request.GET("", s.GetRequest)
			request.PATCH("", s.PatchRequest)
		}

		conference := api.Group("/conferences")
		{
			conference.Use(middlewares.JwtAuthMiddleware())
			conference.GET("", s.GetAllConferences)
			conference.POST("", s.CreateConference)
			conference.GET("/:name", s.GetConference)
			conference.PATCH("/:name", s.PatchConference)
			conference.GET("/:name/participants", s.GetConferenceUsers)
			conference.POST("/:name/participants", s.PatchConferenceUsers)
			conference.GET("/:name/reviews", s.GetConferenceReviews)
			conference.POST("/:name/reviews", s.AddReview)
		}

	}

	return router
}

func getOpenApi(c *gin.Context) {
	plan, err := os.ReadFile("openapi.json")
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	var data interface{}
	err = json.Unmarshal(plan, &data)
	if err != nil {
		jsonInternalServerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
