package router

import (
	"GinChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	indexRoutes := router.Group("/")
	userRoutes := router.Group("/user")
	utilsRoutes := router.Group("/utils")

	// 首页
	{
		indexRoutes.GET("", service.GetIndex)
	}
	// 用户
	{
		userRoutes.GET("", service.GetUserList)
		userRoutes.POST("", service.CreateUser)
	}
	// 工具类
	{
		utilsRoutes.GET("/swagger/*any", service.Swagger())
	}
	return router
}
