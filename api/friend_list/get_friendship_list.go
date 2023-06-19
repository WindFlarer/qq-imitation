package friend_list

import (
	"wechat-imitation/model/common/response"
	Res "wechat-imitation/model/friend_list/response"

	"github.com/gin-gonic/gin"
)

type GetFriendshipListAPI struct{}

// 获取好友列表接口
func (g *GetFriendshipListAPI) GetFriendshipList(c *gin.Context) {
	// var getFriendshipListReq Req.GetFriendshipListReq
	// err := c.ShouldBindJSON(&getFriendshipListReq)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	var getFriendshipListRes []Res.GetFriendshipListRes

	// 获取好友列表逻辑
	getFriendshipListRes, err := getFriendshipListService.GetFriendshipList(c.Request.Header.Get("UserAccount"))
	// 获取失败
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取成功
	response.OkWithData(getFriendshipListRes, c)
}
