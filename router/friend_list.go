package router

import (
	API "wechat-imitation/api"
	"wechat-imitation/middleware"

	"github.com/gin-gonic/gin"
)

func InitFriendshipListRouters(r *gin.Engine) {

	friendshipList := r.Group("/friendshipList")
	friendshipList.Use(middleware.AuthMiddleware())
	{
		friendshipList.GET("/getFriendshipList", API.APIGroupApp.FriendshipListGroup.GetFriendshipListAPI.GetFriendshipList) // 好友列表接口
	}
}
