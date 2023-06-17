package service

import "wechat-imitation/service/login_register"

// 管理所有Service结构体
type ServiceGroup struct {
	LoginRegisterGroup login_register.ServiceGroup
}

// Service管理对象
var ServiceGroupApp = new(ServiceGroup)
