package syslog

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_NewLogger(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	NewLogger(c)
	t.Log("test")
}

func TestMain(t *testing.M) {
	Setup(LoggerConfig{
		LogPath: "/app/logs",
		LogFile: "util-test",
	})
	fmt.Println("start")
	t.Run()
	fmt.Println("end")
}
