package main

import (
	"GinChat/docs"
	"GinChat/router"
	"GinChat/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
func main() {
	// 初始化
	utils.InitConfig()
	utils.InitMySQL()
	// 路由请求
	router := router.Router()
	docs.SwaggerInfo.BasePath = ""
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// 启动
	router.Run("127.0.0.1:8080")
}
