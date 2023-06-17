package login_register

import "wechat-imitation/service"

// login_register目录的管理结构体
type APIGroup struct {
	RegisterAPI
	LoginAPI
}

// 引入相关服务方便后续调用
var (
	loginService    = service.ServiceGroupApp.LoginRegisterGroup.LoginService
	registerService = service.ServiceGroupApp.LoginRegisterGroup.RegisterService
)
