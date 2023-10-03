package main

import (
	"GinChat/common"
	"GinChat/router"
	"GinChat/utils"
)

// @title  标题
// @version 1.0 版本
// @description  描述
// @BasePath /  基础路径
func main() {
	// 初始化
	utils.InitConfig()
	utils.InitMySQL()
	common.CreateTables()
	// 路由请求
	router := router.Router()
	// 启动
	router.Run("127.0.0.1:8080")
}
