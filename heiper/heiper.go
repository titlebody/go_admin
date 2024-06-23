package heiper

import (
	"github.com/dgrijalva/jwt-go"
	"go_admin/define"
)

func GenerateToken(id uint, name string, exprIeAt int64) (string, error) {
	uc := define.User{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exprIeAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(define.JwtKey))
}
