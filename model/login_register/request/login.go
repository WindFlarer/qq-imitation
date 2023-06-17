package request

type LoginReq struct {
	Account  string `gorm:"column:account;type:varchar(255);primaryKey;" json:"account"`
	Password string `gorm:"column:password;type:varchar(255);" json:"password"`
}
