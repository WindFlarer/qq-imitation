package model

type FriendRequest struct {
	SenderAccount   string `gorm:"column:sender_account;type:varchar(255);" json:"senderAccount"`
	ReceiverAccount string `gorm:"column:receiver_account;type:varchar(255);" json:"receiverAccount"`
	Status          string `gorm:"column:status;type:varchar(255);" json:"status"`
}

func (f *FriendRequest) TableName() string {
	return "friend_request"
}
