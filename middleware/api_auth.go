package middleware

import (
	"github.com/gin-gonic/gin"
	"meteor/httptrack/syslog"
	"meteor/model/response"
	"meteor/service"
	"meteor/utils"
)

const (
	AppId           = "App-Id"
	AccessKeyId     = "Access-Key-Id"
	AccessKeySecret = "Access-Key-Secret"
)

/**
接口权限校验  md5明文校验
*/
func ApiAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpHeader := map[string][]string{}
		// 1、获取head所有属性
		for k, v := range c.Request.Header {
			httpHeader[k] = v
			//fmt.Print(k, v)
			//fmt.Print(",")
		}
		syslog.NewLogger(c).Info("当前请求header信息: ", httpHeader)
		// 1.1 获取应用Id
		appId := httpHeader[AppId]
		AccessKeyId := httpHeader[AccessKeyId]
		AccessKeySecret := httpHeader[AccessKeySecret]
		if len(appId) == 0 {
			// 返回，参数错误
			response.ResultResponse(c, 500, response.UrlAppIdOrSecretError, nil)
			c.Abort()
			return
		}
		// 2、鉴权处理 -->> md5加密 匹配
		app, err := service.FindAppByAppId(appId[0])
		if err != nil {
			// 返回，参数错误
			response.ResultResponse(c, 500, response.UrlAppIsNotExist, nil)
			c.Abort()
			return
		}
		if len(app.AccessKeySecret) == 0 || len(app.AccessKeyId) == 0 {
			// 返回，参数错误
			response.ResultResponse(c, 500, response.UrlAppIdOrSecretError, nil)
			c.Abort()
			return
		}
		appToken := utils.EncryptIdAndSecret(app.AccessKeyId, app.AccessKeySecret)
		userToken := utils.EncryptIdAndSecret(AccessKeyId[0], AccessKeySecret[0])
		if appToken != userToken {
			// 返回，参数错误
			response.ResultResponse(c, 500, response.UrlAppIsNotExist, nil)
			c.Abort()
			return
		}
		// 可以设置变量在context中
		c.Set("uid", app.Uid)
		c.Set("appId", app.AppId)
		c.Next()
	}
}
