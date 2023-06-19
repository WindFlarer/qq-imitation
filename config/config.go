package config

import "github.com/golang-jwt/jwt"

type ConfStruct struct {
	Mysql Mysql `yaml:"mysql"`
}

// 管理所有的默认配置结构体
var Conf ConfStruct

type UserClaim struct {
	UserAccount string `json:"userAccount"`
	jwt.StandardClaims
}

// token密码
var JwtKey = "wechat-imitation-key"

// token过期时间
var TokenExpire = 86400
