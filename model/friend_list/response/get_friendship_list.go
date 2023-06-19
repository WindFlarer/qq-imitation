package response

import "time"

type GetFriendshipListRes struct {
	FriendAccount string ` json:"friendAccount"`
	Remark        string ` json:"remark"`
	Path          string ` json:"path"`
	Status        uint8  ` json:"status"`
	CreatedAt     time.Time
}
