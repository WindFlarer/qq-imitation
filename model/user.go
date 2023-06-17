package model

import (
	"time"
)

type User struct {
	Account   string `gorm:"column:account;type:varchar(255);primaryKey;" json:"account"`
	Password  string `gorm:"column:password;type:varchar(255);" json:"password"`
	Username  string `gorm:"column:username;type:varchar(255);" json:"username"`
	Phone     string `gorm:"column:phone;type:varchar(255);" json:"phone"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	Gender    uint8  `gorm:"column:gender;type:tinyint;" json:"gender"`
	Status    uint8  `gorm:"column:status;type:tinyint;" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "user"
}
