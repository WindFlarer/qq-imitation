package model

type Group struct {
	GroupAccount   string `gorm:"column:group_account;type:varchar(255);" json:"groupAccount"`
	Groupname      string `gorm:"column:groupname;type:varchar(255);" json:"groupname"`
	Username       string `gorm:"column:username;type:varchar(255);" json:"username"`
	CreaterAccount string `gorm:"column:creater_account;type:varchar(255);" json:"createrAccount"`
	GroupMeMber    string `gorm:"column:group_member;type:text;" json:"groupMember"`
	Path           string `gorm:"column:path;type:varchar(255);" json:"path"`
}

func (g *Group) TableName() string {
	return "group"
}
