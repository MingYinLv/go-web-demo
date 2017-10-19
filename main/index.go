package main

import (
	"github.com/gin-gonic/gin"
	"router/handle"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/users", handle.GetUserList)
	}
	{
		user := v1.Group("/user")
		user.GET("/:id", handle.GetUserById)
		user.POST("/", handle.AddUser)
		user.PUT("/:id", handle.UpdateUser)
		user.DELETE("/:id", handle.DeleteUserById)
	}
	router.Run()
}
