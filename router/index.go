package router

import (
	"GinChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/index", service.GetIndex)
	return router
}
