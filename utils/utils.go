package utils

import (
	"net/http"
	"net/url"
)

/**
正则表达式判断url是否有效
*/
func CheckUrlIsValid(result string) bool {
	if result != "" {
		// 解析url  无效返回err
		uStruct, err := url.ParseRequestURI(result)
		if err != nil {
			//global.GVA_LOG.Error("【CheckUrlIsValid, url格式错误】", zap.Any("err", err))
			return false
		}
		//global.GVA_LOG.Info("【CheckUrlIsValid-url解析结果】", zap.Any("uStruct", uStruct))
		println(uStruct)
		return true
	}
	return false
}

/**
获取http请求的完整路径
*/
func GetFullUrlAddr(w http.ResponseWriter, r *http.Request) string {
	return "http://" + r.Host + r.RequestURI
}
