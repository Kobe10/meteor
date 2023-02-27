package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}

//正常响应
func ResultResponse(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status":        GetStatus(errCode),
		"error_code":    errCode,
		"error_message": GetMsg(errCode),
		"data":          data,
	})
	return
}

func Result(status string, code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		status,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(IsOk, SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(IsOk, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(IsOk, SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(IsOk, SUCCESS, data, message, c)
}

func OkWithDetailedAndCode(code int, data interface{}, message string, c *gin.Context) {
	Result(IsOk, code, data, message, c)
}

func Fail(c *gin.Context) {
	Result(IsFail, ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(IsFail, ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(IsFail, ERROR, data, message, c)
}

func FailWithDetailedCode(data interface{}, message string, code int, c *gin.Context) {
	Result(IsFail, code, data, message, c)
}
