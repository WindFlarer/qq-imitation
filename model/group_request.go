package model

import "time"

type GroupRequest struct {
	SenderAccount   string `gorm:"column:sender_account;type:varchar(255);" json:"senderAccount"`
	ReceiverAccount string `gorm:"column:receiver_account;type:varchar(255);" json:"receiverAccount"`
	GroupAccount    string `gorm:"column:group_account;type:varchar(255);" json:"groupAccount"`
	Status          uint8  `gorm:"column:status;type:tinyint;" json:"status"`
	CreatedAt       time.Time
}

func (g *GroupRequest) TableName() string {
	return "group_request"
}
