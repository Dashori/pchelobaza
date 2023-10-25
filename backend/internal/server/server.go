package server

import (
	// registry "backend/cmd/registry"

	"backend/internal/app"
	// "backend/internal/server/controllers"
	"backend/internal/server/middlewares"
	// "fmt"
	"github.com/gin-gonic/gin"
)

type services struct {
	Services *app.AppServiceFields
}

func SetupServer(a *app.App) *gin.Engine {
	s := services{a.Services}

	router := gin.Default()

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

	}

	return router
}
