package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetKey(c *gin.Context) {
	// redisRepo, ok := c.MustGet("redis_repo").(redisrepo.RedisRepository)
	// if ok != true {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
	// 		Message: "Failed to get redis_repo. Please, try again later"})
	// 	return
	// }
	fmt.Printf("AAAAAA")
}

func SetupServer() *gin.Engine {

	router := gin.Default()

	redisApi := router.Group("/")
	{
		redisApi.GET("/get_key", GetKey)
		// redisApi.POST("/set_key", controllers.SetKey)
		// redisApi.DELETE("/del_key", controllers.DelKey)
	}

	return router
}
