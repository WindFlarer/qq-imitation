package model

type GroupRequest struct {
	SenderAccount   string `gorm:"column:sender_account;type:varchar(255);" json:"senderAccount"`
	ReceiverAccount string `gorm:"column:receiver_account;type:varchar(255);" json:"receiverAccount"`
	GroupAccount    string `gorm:"column:group_account;type:varchar(255);" json:"groupAccount"`
	Status          string `gorm:"column:status;type:varchar(255);" json:"status"`
}

func (g *GroupRequest) TableName() string {
	return "group_request"
}
