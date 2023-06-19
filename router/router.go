package router

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	// 初始化用户路由
	InitUserRouters(r)
	// 初始化好友列表路由
	InitFriendshipListRouters(r)
}
