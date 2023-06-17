package login_register

import (
	"errors"
	"wechat-imitation/initialization"
	"wechat-imitation/model"
	Req "wechat-imitation/model/login_register/request"
)

type LoginService struct{}

// 用户登录逻辑
func (l *LoginService) Login(loginReq Req.LoginReq) error {
	//查询是否存在当前用户
	accountResult := initialization.DB.Where("account = ? ", loginReq.Account).First(&model.User{})

	//判断是否存在
	if accountResult.RowsAffected > 0 { //如果存在
		if passwordResult := initialization.DB.Where("account =? AND password =?", loginReq.Account, loginReq.Password).First(&model.User{}); passwordResult.RowsAffected > 0 { //密码正确
			//生成token
			// token, _ := helpers.GenerateToken(userBasic.ID, userBasic.UserName, config.TokenExpire)
			return nil
		} else { //密码错误
			return errors.New("密码错误")
		}
	} else { //不存在
		return errors.New("用户不存在")
	}
}
