package model

type Friendship struct {
	UserAccount   string `gorm:"column:user_account;type:varchar(255);" json:"userAccount"`
	FriendAccount string `gorm:"column:friend_account;type:varchar(255);" json:"friendAccount"`
	Remark        string `gorm:"column:remark;type:varchar(255);" json:"remark"`
	Path          string `gorm:"column:path;type:varchar(255);" json:"path"`
}

func (f *Friendship) TableName() string {
	return "friendship"
}
