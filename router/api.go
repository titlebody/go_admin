package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/middleware"
	"go_admin/service"
)

func Init() *gin.Engine {
	r := gin.Default()
	// 解决跨域问题
	r.Use(middleware.Cors())
	//api端口
	r.POST("/login", service.Login)

	r.POST("/setLogin", service.GetLogin)

	return r
}
