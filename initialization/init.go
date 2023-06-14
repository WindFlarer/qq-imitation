package initialization

func InitAll() {
	// 初始化
	LoadConfig() // 读取yaml
	InitDB()     // 连接数据库
	InitTable()  // 初始化表
}
