package service

import (
	"wechat-imitation/service/friend_list"
	"wechat-imitation/service/login_register"
)

// 管理所有Service结构体
type ServiceGroup struct {
	LoginRegisterGroup login_register.ServiceGroup
	FriendListGroup    friend_list.ServiceGroup
}

// Service管理对象
var ServiceGroupApp = new(ServiceGroup)
