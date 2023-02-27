package httptrack

import (
	"fmt"
	"meteor/httptrack/syslog"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const (
	maxIdleConns        int = 300
	maxIdleConnsPerHost int = 200
	idleConnTimeout     int = 300 //keepalive 5 min
)

type IService interface {
	GetResponse() (entry interface{}, err error)
	ParseRequest(params map[string]interface{}) bool
	GetApiName() (req string)
}

type HttpService struct {
	IService
	ApiUri  string
	Timeout time.Duration
	syslog.LogHandler
}

type CommonConfig struct {
	Env   string
	AppId string
}

type CommonHttpInfoLog struct {
	ReqDUrl   string // url
	ReqParams string // 请求参数
	ReqCost   int64  // 请求花费时间
	RespCode  int    // 响应码
}

var httpClient *http.Client

var httpClientConfig CommonConfig

func SetUp(config CommonConfig) {

	httpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   400 * time.Millisecond,
				KeepAlive: 300 * time.Second,
			}).DialContext,
			MaxIdleConns:        maxIdleConns,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(idleConnTimeout) * time.Second,
		},
		Timeout: 1500 * time.Millisecond,
	}

	httpClientConfig = config

}

func GetHttpClient() *http.Client {
	return httpClient
}

func GetEnv() string {
	return httpClientConfig.Env
}

func GetAppId() string {
	return httpClientConfig.AppId
}

//设置超时时间
func (p *HttpService) SetTimeout(time time.Duration) {
	p.Timeout = time
}

//通过解析结构体构建查询参数 依据字段tag ：json构建
func (p *HttpService) BuildGetParams(paramsData interface{}) {
	v := reflect.ValueOf(paramsData)
	count := v.NumField()
	params := make([]string, 0)
	typ := reflect.TypeOf(paramsData)
	for i := 0; i < count; i++ {
		f := v.Field(i)
		tag := typ.Field(i).Tag.Get("json")
		params = append(params, fmt.Sprintf("%v=%v", tag, f))
	}
	if len(params) > 0 {
		p.ApiUri += "?" + strings.Join(params, "&")
	}
}
