package api

import (
	"wechat-imitation/api/friend_list"
	"wechat-imitation/api/login_register"
)

// 管理所有API结构体
type APIGroup struct {
	LoginRegisterGroup  login_register.APIGroup
	FriendshipListGroup friend_list.APIGroup
}

// API管理对象
var APIGroupApp = new(APIGroup)
