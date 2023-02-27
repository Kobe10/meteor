// +build !windows

package core

import (
	"meteor/global"
	"meteor/httptrack"
	"meteor/httptrack/syslog"
)

// 优雅重启服务
func InitHttpTrack() {
	//初始化配置文件

	// http通用包初始化
	httptrack.SetUp(httptrack.CommonConfig{
		Env:   global.GVA_VP.GetString("run_mode"),
		AppId: global.GVA_VP.GetString("appId"),
	})

	//日志
	syslog.Setup(syslog.LoggerConfig{
		LogPath: "",
		LogFile: "",
		Env:     global.GVA_VP.GetString("run_mode"),
	})
}
