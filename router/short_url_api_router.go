package router

import (
	"github.com/gin-gonic/gin"
	"meteor/controllers/inside/v1"
	"meteor/middleware"
)

func ShortUrlRouter(Router *gin.RouterGroup) {
	// 路由分组 针对短链相关的api 进行权限拦截校验
	DRouter := Router.Group("api/short_url").Use(middleware.ApiAuthorize())
	{
		DRouter.POST("/create", v1.CreateShortUrl) // 创建url
	}
}
