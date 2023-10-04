package router

import (
	"GinChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	userRoutes := router.Group("/user")
	utilsRoutes := router.Group("/utils")
	// 用户
	{
		userRoutes.GET("", service.GetUserList)
		userRoutes.POST("", service.CreateUser)
		userRoutes.DELETE(":id", service.DeleteUser)
		userRoutes.PATCH(":id", service.UpdateUser)
		userRoutes.POST("/login", service.UserLogin)
	}
	// 工具类
	{
		utilsRoutes.GET("/swagger/*any", service.Swagger())
	}
	return router
}
