package define

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
)

var (
	// JwtKey 密钥
	JwtKey = "sys-admin"
	// TokenExpire Token有效期限
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// RefreshTokenExpire 刷新Token有效期限
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	DB                 *gorm.DB
)

type User struct {
	Id                 uint
	Name               string
	IsAdmin            bool // 是否授权
	jwt.StandardClaims      // jwt.StandardClaims
}
