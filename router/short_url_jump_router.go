package router

import (
	"github.com/gin-gonic/gin"
	"meteor/controllers/inside/v1"
)

func ShortUrlJumpRouter(Router *gin.RouterGroup) {
	// 路由分组 针对短链相关的api 进行权限拦截校验

	DRouter := Router.Group("/url")
	{
		DRouter.GET("/:url", v1.ShortUrlJump) // 创建url
	}
}
