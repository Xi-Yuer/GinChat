package main

import (
	"GinChat/router"
	"GinChat/utils"
)

func main() {
	// 初始化
	utils.InitConfig()
	utils.InitMySQL()
	// 路由请求
	router := router.Router()
	// 启动
	router.Run(":8080")
}
