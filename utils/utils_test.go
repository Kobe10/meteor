package utils

import (
	"fmt"
	"meteor/core"
	"meteor/global"
	"meteor/httptrack"
	"testing"
)

/**
url正确性校验 单测
*/
func Test_checkUrlIsValid(t *testing.T) {
	url1 := "http://www.baidu.com"
	url2 := "www.baidu.com"
	fmt.Print(CheckUrlIsValid(url1))
	fmt.Print(CheckUrlIsValid(url2))
}

// 初始化62进制随机字符
func Test_init62Str(t *testing.T) {
	init62Str()
}

// 62进制转换
func Test_10to62(t *testing.T) {
	println(DecimalTo62(62, 62))
}

// 10进制转换
func Test_62to10(t *testing.T) {
	println(anyToDecimal("tFF", 62))
}

func TestMain(m *testing.M) {
	global.GVA_LOG = core.Zap() // 初始化zap日志库

	httptrack.SetUp(httptrack.CommonConfig{
		Env:   "",
		AppId: "test",
	})
	m.Run()
}
