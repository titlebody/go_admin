package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_admin/define"
	"go_admin/heiper"
	"go_admin/models"
	"gorm.io/gorm"
)

// Login 用户登录
func Login(c *gin.Context) {
	i := new(LoginRequest)
	err := c.ShouldBindJSON(i)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// 根据账号密码查询
	sysUser, e := models.GetUserPassBox(i.Username, i.Password) // 查询用户
	if e != nil {                                               // 查询失败
		if errors.Is(e, gorm.ErrRecordNotFound) { //
			c.JSON(200, gin.H{
				"code":    400,
				"message": "用户名或者密码错误",
			})
			return
		}
	}

	token, es := heiper.GenerateToken(sysUser.ID, sysUser.Password, define.TokenExpire)
	if es != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": es.Error(),
		})
		return
	}

	//刷新token
	refreshToken, s := heiper.GenerateToken(sysUser.ID, sysUser.Password, define.RefreshTokenExpire)
	if s != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": es.Error(),
		})
		return
	}
	data := &LoginSum{
		Token:        token,
		RefreshToken: refreshToken,
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "登录成功",
		"data":    data,
	})
}
