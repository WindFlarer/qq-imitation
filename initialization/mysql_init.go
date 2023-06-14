package initialization

import (
	"fmt"
	"wechat-imitation/config"
	"wechat-imitation/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := config.Conf.Mysql.Username +
		":" + config.Conf.Mysql.Password +
		"@tcp(" + config.Conf.Mysql.Path +
		":" + config.Conf.Mysql.Port +
		")/" + config.Conf.Mysql.Dbname +
		"?" + config.Conf.Mysql.Config

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	fmt.Println("连接数据库成功")
}

func InitTable() {

	//创建基本表
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Friendship{})
	DB.AutoMigrate(&model.Group{})
	DB.AutoMigrate(&model.FriendRequest{})
	DB.AutoMigrate(&model.GroupRequest{})
}
