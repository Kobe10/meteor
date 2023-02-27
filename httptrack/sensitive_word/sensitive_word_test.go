package sensitive_word

import (
	"github.com/gin-gonic/gin"
	"meteor/httptrack"
	"meteor/httptrack/syslog"
	"net/http/httptest"
	//"net/http/httptest"
	"testing"
	//"time"
)

func Test_sentinelWordFilter(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	api := NewSensitiveWordApi(c)
	api.SetTimeout(10000)
	got, err := api.GetSensitiveApiResponse(SensitiveWord{Content: "http://www.baidu.com"})
	if err != nil {
		println(err.Error())
	}
	println(got)

	got1, err := api.GetSensitiveApiResponse(SensitiveWord{Content: "http://www.baidu.com/苍井空"})
	if err != nil {
		println(err.Error())
	}
	println(got1)
}

func TestMain(m *testing.M) {

	// http通用包初始化
	httptrack.SetUp(httptrack.CommonConfig{
		Env:   "feature",
		AppId: "meteor",
	})
	//日志
	syslog.Setup(syslog.LoggerConfig{
		LogPath: "/User/fuzhiqiang/logs",
		LogFile: "app",
		Env:     "feature",
	})

	m.Run()
}
