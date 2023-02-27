package v1

import (
	"github.com/gin-gonic/gin"
	"meteor/model/request/url"
	"meteor/model/response"
	"meteor/service"
	"meteor/utils"
)

// @Tags Demo
// @Summary 长url转短url
// @accept application/json
// @Produce application/json
// @Param data body url.CreateShortUrlReq true "生成短链"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router  /api/short_url/create [post]
func CreateShortUrl(c *gin.Context) {
	var createShortUrl url.CreateShortUrlReq
	// 1、转换json
	_ = c.ShouldBindJSON(&createShortUrl)
	// 2、url校验
	//urlStatus, err := service.CheckSourceUrl(createShortUrl.SourceUrl, c)
	//if err != nil {
	//	response.ResultResponse(c, 500, response.UrlParamsError, nil)
	//	return
	//}
	//if !urlStatus {
	//	response.ResultResponse(c, 500, response.UrlParamsError, nil)
	//	return
	//}
	// 3、生成短链
	var shortUrlDTO url.ShortUrlDTO
	shortUrlDTO = TransformShortUrlDTO(&createShortUrl, c)
	shortUrl, generateErr := service.GenerateShortChains(&shortUrlDTO, c)
	if generateErr != nil {
		response.ResultResponse(c, 500, response.UrlIsNotValid, nil)
		return
	}
	response.ResultResponse(c, 200, 200, shortUrl)
	return
}

/**
DTO对象转换
*/
func TransformShortUrlDTO(createShortUrl *url.CreateShortUrlReq, c *gin.Context) url.ShortUrlDTO {
	var shortUrlDTO url.ShortUrlDTO
	shortUrlDTO.SourceUrl = createShortUrl.SourceUrl
	//1、对长url进行md5加密 ，后续流程都通过md5来查询和插入
	shortUrlDTO.SourceUrlMD5 = utils.MD5V([]byte(createShortUrl.SourceUrl))

	//2、从context中获取appId和用户信息
	shortUrlDTO.AppId = c.GetString("appId")
	shortUrlDTO.Uid = c.GetString("uid")
	return shortUrlDTO
}

// @Tags Demo
// @Summary 短链跳转
// @accept application/json
// @Produce application/json
// @Param url path string true "test"
// @Success 302  json jsong "返回结果"
// @Router  /url/:url [get]
func ShortUrlJump(c *gin.Context) {
	// 1、获取完整请求地址， 缓存中查询是否存在该短链地址
	fullAddr := utils.GetFullUrlAddr(c.Writer, c.Request)
	sourceUrl, err := service.FindSourceChainsFormRedisAndDb(fullAddr)
	if err != nil {
		//response.FailWithDetailedCode(nil, response.UrlRuleIsNotValid, response.UrlIsNotValid, c)
		return
	}
	if len(sourceUrl) == 0 {
		//response.FailWithDetailedCode(nil, response.UrlRuleIsNotValid, response.UrlIsNotValid, c)
		return
	}
	// 2、记录请求信息

	// 3、自动跳转
	c.Redirect(301, sourceUrl)
	c.Abort()
	//response.ResultResponse(c, 302, 200, sourceUrl)
}
