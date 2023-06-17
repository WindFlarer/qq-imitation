package model

import (
	"time"
)

type User struct {
	Account   string `gorm:"column:account;type:varchar(255);primaryKey;" json:"account"`
	Password  string `gorm:"column:password;type:varchar(255);" json:"password"`
	Username  string `gorm:"column:username;type:varchar(255);" json:"username"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	Gender    string `gorm:"column:gender;type:varchar(255);" json:"gender"`
	Status    uint8  `gorm:"column:status;type:tinyint;" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "user"
}
