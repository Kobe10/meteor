package main

import (
	"meteor/core"
	"meteor/global"
	"meteor/initialize"
)

func init() {
	global.GVA_VP = core.Viper()        // 初始化Viper
	global.GVA_ENV_VP = core.EnvViper() // 初始化Viper
}

//@title 短链API文档
//@version 0.0.1
//@license.name  meteor
//@description This is a sample Server pets
//@in header
//@name x-token
//@BasePath /
func main() {
	//global.GVA_VP = core.Viper()        // 初始化Viper
	//global.GVA_ENV_VP = core.EnvViper() // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	core.InitHttpTrack()              // 初始化http通用组件
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	core.RunWindowsServer()

}
