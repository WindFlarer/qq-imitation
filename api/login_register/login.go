package login_register

import (
	"wechat-imitation/model/common/response"
	Req "wechat-imitation/model/login_register/request"

	"github.com/gin-gonic/gin"
)

type LoginAPI struct{}

// 用户登录接口
func (r *LoginAPI) Login(c *gin.Context) {
	var loginReq Req.LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 登录逻辑
	err = loginService.Login(loginReq)
	// 登录失败
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 登录成功
	response.OkWithMessage("登录成功", c)
}
