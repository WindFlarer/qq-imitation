package request

type RegisterReq struct {
	Account  string ` json:"account"`
	Password string ` json:"password"`
	Username string ` json:"username"`
	Phone    string ` json:"phone"`
	Gender   uint8  ` json:"gender"`
}
