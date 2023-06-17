package login_register

import (
	"errors"
	"wechat-imitation/initialization"
	"wechat-imitation/model"
	Req "wechat-imitation/model/login_register/request"
)

type RegisterService struct{}

// 用户注册逻辑
func (r *RegisterService) Register(registerReq Req.RegisterReq) error {
	//查询是否存在当前用户
	accountResult := initialization.DB.Where("account = ? ", registerReq.Account).First(&model.User{})

	//判断是否存在
	if accountResult.RowsAffected > 0 { //如果存在
		return errors.New("账号已存在")
	} else { //不存在
		//创建用户
		user := model.User{
			Account:  registerReq.Account,
			Password: registerReq.Password,
			Username: registerReq.Username,
			Phone:    registerReq.Phone,
			Gender:   registerReq.Gender,
		}

		// 保存用户
		initialization.DB.Create(&user)
		return nil
	}
}
