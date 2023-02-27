package core

import (
	"fmt"
	"go.uber.org/zap"
	"meteor/global"
	"meteor/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化redis服务
	global.GVA_LOG.Info("init redis")
	initialize.Redis()

	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	流星启动成功
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
