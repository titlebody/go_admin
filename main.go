package main

import (
	"go_admin/dao"
	"go_admin/router"
)

func main() {
	//初始化mysql
	dao.InitMySQL()
	//初始化路由
	r := router.Init()

	//启动服务
	r.Run(":8080")

	//关闭

}
