package server

import (
	"backend/internal/app"
	"backend/internal/server/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type services struct {
	Services *app.AppServiceFields
}

func SetupServer(a *app.App) *gin.Engine {
	s := services{a.Services}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", s.Login)
			auth.POST("/signup", s.SignUp)
		}

		user := api.Group("/users")
		{
			user.Use(middlewares.JwtAuthMiddleware())
			user.GET("/:login", s.GetUser)
			user.PATCH("/:login", s.PatchUser)
		}

		farm := api.Group("/farms")
		{
			farm.Use(middlewares.JwtAuthMiddleware())
			farm.GET("", s.GetFarms)
			farm.POST("", s.AddFarm)
			farm.GET("/:name", s.GetFarmInfo)
			farm.PATCH("/:name", s.PatchFarm)
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
