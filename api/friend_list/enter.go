package friend_list

import "wechat-imitation/service"

// login_register目录的管理结构体
type APIGroup struct {
	GetFriendshipListAPI
}

// 引入相关服务方便后续调用
var (
	getFriendshipListService = service.ServiceGroupApp.FriendListGroup.GetFriendshipListService
)
