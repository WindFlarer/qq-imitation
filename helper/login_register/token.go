package login_register

import (
	"errors"
	"time"
	"wechat-imitation/config"

	"github.com/golang-jwt/jwt"
)

// 生成令牌
func GenerateToken(userAccount string, second int) (string, error) {
	//生成载体
	uc := config.UserClaim{
		UserAccount: userAccount,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(second) * time.Second).Unix(),
		},
	}

	//生成令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析令牌
func AnalyzeToken(token string) (*config.UserClaim, error) {
	uc := new(config.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
