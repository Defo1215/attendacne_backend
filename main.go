package main

import (
	"attendance/app/controllers"
	"attendance/database"
	"attendance/routes"
)

func main() {
	database.InitMySQL()          // 初始化MySQL数据库连接
	controllers.InitQrCodeTasks() // 初始化二维码任务

	r := routes.InitRouter() // 初始化路由

	_ = r.Run(":8080")
}
