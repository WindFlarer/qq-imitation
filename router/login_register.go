package router

import (
	API "wechat-imitation/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouters(r *gin.Engine) {

	user := r.Group("/user")
	{
		user.POST("/register", API.APIGroupApp.LoginRegisterGroup.RegisterAPI.Register) // 用户注册接口
		user.POST("/login", API.APIGroupApp.LoginRegisterGroup.LoginAPI.Login)          // 用户登录接口
	}
}
