package server

import (
	// registry "backend/cmd/registry"

	"backend/internal/app"
	// "backend/internal/server/controllers"
	"backend/internal/server/middlewares"
	"fmt"
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
			// auth.POST("/logout", controllers.DeleteFileContent)
		}

		user := api.Group("/users")
		{
			user.Use(middlewares.JwtAuthMiddleware())
			user.GET("", s.GetUser)
			user.PATCH("", s.PatchUser)
		}

		farm := api.Group("/farms")
		{
			farm.GET("", s.GetFarms)
			farm.POST("", s.AddFarm)
		}

		api.POST("/setRole", s.setRole)

		// api.GET("/doctors", t.getAllDoctors)
		// api.POST("/doctor/create", t.createDoctor)
		// api.POST("/doctor/login", t.loginDoctor)

		// doctor := api.Group("/doctor")
		// doctor.Use(middlewares.JwtAuthMiddleware())
		// doctor.GET("/info", t.doctorInfo)
		// doctor.GET("/records", t.doctorRecords)
		// doctor.PATCH("/shedule", t.doctorShedule)

		// api.POST("/client/create", t.createClient)
		// api.POST("/client/login", t.loginClient)

		// client := api.Group("/client")
		// client.Use(middlewares.JwtAuthMiddleware())
		// client.GET("/info", t.infoClient)
		// client.GET("/records", t.ClientRecords)
		// client.GET("/pets", t.ClientPets)
		// client.POST("/record", t.NewRecord)
		// client.POST("/pet", t.NewPet)
		// client.DELETE("/pet", t.DeletePet)
	}

	// port := a.Config.Port
	// adress := a.Config.Address
	// err := router.Run(adress + port)

	// if err != nil {
	// 	return nil
	// }

	return router
}

func (s *services) setRole(c *gin.Context) {
	// var role *Role
	// err := c.ShouldBindJSON(&role)
	// if err != nil {
	// 	jsonInternalServerErrorResponse(c, err)
	// 	return
	// }

	// if role.Role == "doctor" {
	// 	err = t.Services.DoctorService.SetRole()
	// } else if role.Role == "client" {
	// 	err = t.Services.ClientService.SetRole()
	// } else {
	// 	jsonBadRequestResponse(c, fmt.Errorf("Такой роли не существует!"))
	// }

	// if err != nil {
	// 	jsonInternalServerErrorResponse(c, err)
	// 	return
	// }

	// jsonStatusOkResponse(c)
	fmt.Println("bbbb")
	return
}
