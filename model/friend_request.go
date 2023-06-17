package model

import "time"

type FriendRequest struct {
	SenderAccount   string `gorm:"column:sender_account;type:varchar(255);" json:"senderAccount"`
	ReceiverAccount string `gorm:"column:receiver_account;type:varchar(255);" json:"receiverAccount"`
	Status          uint8  `gorm:"column:status;type:tinyint;" json:"status"`
	CreatedAt       time.Time
}

func (f *FriendRequest) TableName() string {
	return "friend_request"
}
