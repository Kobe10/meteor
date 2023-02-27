package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"meteor/httptrack/sensitive_word"
	"meteor/utils"
)

func CheckSourceUrl(sourceUrl string, c *gin.Context) (status bool, err error) {
	//1、校验url合法性
	if !utils.CheckUrlIsValid(sourceUrl) {
		return false, errors.New("url不合法")
	}
	//2、敏感词校验
	api := sensitive_word.NewSensitiveWordApi(c)
	api.SetTimeout(sensitive_word.SensitiveWordTimeout)
	isSensitive, err := api.GetSensitiveApiResponse(sensitive_word.SensitiveWord{Content: sourceUrl})

	if !isSensitive {
		return false, errors.New("url不合法")
	}
	return true, err
}
