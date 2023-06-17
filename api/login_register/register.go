package login_register

import (
	"wechat-imitation/model/common/response"
	Req "wechat-imitation/model/login_register/request"

	"github.com/gin-gonic/gin"
)

type RegisterAPI struct{}

func (r *RegisterAPI) Register(c *gin.Context) {
	var registerReq Req.RegisterReq
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 注册逻辑
	err = registerService.Register(registerReq)
	// 注册失败
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 注册成功
	response.OkWithMessage("注册成功", c)
}
