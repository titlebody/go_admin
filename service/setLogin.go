package service

import (
	"github.com/gin-gonic/gin"
	"go_admin/models"
)

func GetLogin(c *gin.Context) {
	i := new(LoginRequest)
	err := c.ShouldBindJSON(&i)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "数据不能为空",
		})
		return
	}
	//添加进入数据库
	e := models.AddUser(i.Username, i.Password)
	if e != nil {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "添加失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}
