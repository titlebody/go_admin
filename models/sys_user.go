package models

import (
	"go_admin/define"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Username  string `gorm:"column:username;type:varchar(50);"  json:"username"`
	Password  string `gorm:"column:password;type:varchar(36);"  json:"password"`
	Phone     string `gorm:"column:phone;type:varchar(20);"  json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(255);"  json:"wxUnionId"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(255);"  json:"wxOpenId"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);"  json:"avatar"`
}

func (s *SysUser) TableName() string {
	return "sys_user"
}

// GetUserPassBox 根据用户名密码查询数据
func GetUserPassBox(user, password string) (*SysUser, error) {
	data := new(SysUser)
	err := define.DB.Where("username=? and password=?", user, password).First(&data).Error
	return data, err
}
