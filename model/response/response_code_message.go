package response

const (
	ERROR         = 500
	SUCCESS       = 200
	InvalidParams = 400
	VERSION       = 10000
)

const (
	IsFail = "fail"
	IsOk   = "success"
)

// 错误编码
const (
	UrlParamsError        = 10001
	UrlIsNotValid         = 10002
	UrlAppIsNotExist      = 10003
	UrlAppIdOrSecretError = 10004
)

// 错误信息
const ()

var MsgFlags = map[int]string{
	SUCCESS:               "",
	ERROR:                 "服务器开小差了",
	InvalidParams:         "请求参数错误",
	VERSION:               "header版本号错误",
	UrlParamsError:        "url格式不正确",
	UrlIsNotValid:         "url无效",
	UrlAppIsNotExist:      "应用不存在，请创建应用后再试",
	UrlAppIdOrSecretError: "应用账号密码配置错误，请联系管理员配置",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

func GetStatus(code int) string {
	if code == SUCCESS {
		return "success"
	} else {
		return "failure"
	}
}
