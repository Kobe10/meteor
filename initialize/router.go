package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "meteor/docs"
	"meteor/global"
	"meteor/httptrack/syslog"
	"meteor/middleware"
	"meteor/router"
	_ "net/http"
)

// 初始化总路由
func Routers() *gin.Engine {
	//这个是默认的日志组件
	//var Router = gin.Default()

	//这个使用 基于zap的扩展路由日志中间件   gin.Recovery() 与 Logger() 为Gin框架的 gin.Default() 默认使用的全局中间件
	var Router = gin.New()
	Router.Use(middleware.ZapLogger(), gin.Recovery())

	//自定义日志中间件
	Router.Use(syslog.LoggerToFile())
	Router.Use(gin.Recovery())

	// 打开就能玩https了
	Router.Use(middleware.LoadTls())
	global.GVA_LOG.Info("use middleware logger")
	// 应用跨域中间件
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")

	// swagger 路由
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.ShortUrlRouter(ApiGroup)     // 创建短url
	router.ShortUrlJumpRouter(ApiGroup) // 跳转短url
	// 如果有多个routers  在后面依次添加
	//router.InitUserRouter(ApiGroup)                  // 注册用户路由

	global.GVA_LOG.Info("router register success")
	return Router
}
