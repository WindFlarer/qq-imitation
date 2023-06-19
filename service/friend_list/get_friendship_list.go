package friend_list

import (
	"errors"
	"wechat-imitation/initialization"
	"wechat-imitation/model"
	Res "wechat-imitation/model/friend_list/response"
)

type GetFriendshipListService struct{}

func (g *GetFriendshipListService) GetFriendshipList(userAccount string) ([]Res.GetFriendshipListRes, error) {
	// 好友列表
	var getFriendshipListRes []Res.GetFriendshipListRes
	// 查询好友列表
	result := initialization.DB.Model(model.Friendship{}).Where("user_account = ?", userAccount).Find(&getFriendshipListRes)
	if result.RowsAffected > 0 {
		return getFriendshipListRes, nil
	} else {
		return nil, errors.New("好友列表为空")
	}
}
