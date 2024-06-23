package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/middleware"
	"go_admin/service"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/", service.Login)
	return r
}
