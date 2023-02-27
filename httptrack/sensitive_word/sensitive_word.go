package sensitive_word

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"meteor/httptrack"
	"meteor/httptrack/responses"
	"net/http"
	"runtime/debug"
	"time"
)

const SensitiveWordTimeout = 500 * time.Millisecond
const apiUrl = "/api/sensitive"

//api结构体
type SensitiveWordApi struct {
	httptrack.HttpService
	Params SensitiveWord
}

// 敏感词请求实体
type SensitiveWord struct {
	Content string `json:"content" form:"content"`
}

// 敏感词结果响应体
type SensitiveResponse struct {
	Data    SensitiveListEntry `json:"data"`
	Code    int64              `json:"code"`
	Message string             `json:"message"`
}

type SensitiveListEntry struct {
	Sensitive bool             `json:"sensitive"`
	IkList    []SensitiveEntry `json:"ikList"`
}

type SensitiveEntry struct {
	Word      string `json:"word"`
	Count     int64  `json:"count"`
	WordLevel string `json:"wordLevel"`
}

func NewSensitiveWordApi(ctx *gin.Context) *SensitiveWordApi {
	a := new(SensitiveWordApi)
	//url初始化
	a.ApiUri = GetHost() + apiUrl
	//设置logger
	a.SetLogger(ctx)
	//设置默认超时
	//a.SetTimeout(SensitiveWordTimeout)
	a.SetTimeout(1)
	return a
}

// 敏感词过滤
// url  请求url
func (params *SensitiveWordApi) GetResponse() (bool, error) {
	var response SensitiveResponse
	start := time.Now()

	////1、请求数据拼接
	//sensitiveConfig := global.GVA_CONFIG.Sensitive
	////1.1 敏感词url
	//sensitiveUrl := sensitiveConfig.Domain + sensitiveConfig.Sensitive
	sensitiveUrl := params.ApiUri

	//2、调用敏感词系统接口
	httpClient := httptrack.GetHttpClient()
	j, _ := responses.GetJsonParser().Marshal(params.Params)

	buffer := bytes.NewBuffer(j)

	ctx, _ := context.WithTimeout(context.TODO(), params.Timeout)
	req, err := http.NewRequestWithContext(ctx, "POST", sensitiveUrl, buffer)
	if err != nil {
		params.GetLogger().Error("NewRequestWithContext Err:", err.Error(), ",url:", sensitiveUrl, ",params:"+string(j))
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("authorization", GetAccessKey())
	req.Header.Set("cache-control", "no-cache")

	resp, err := httpClient.Do(req)
	if err != nil {
		params.GetLogger().Error("sentinelWordFilter-调用敏感词系统异常 Request Err:", err.Error(), ",url:", sensitiveUrl, ",params:", string(j))
		return false, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	commonHttpInfoLog := new(httptrack.CommonHttpInfoLog)
	commonHttpInfoLog.ReqCost = duration
	commonHttpInfoLog.RespCode = resp.StatusCode
	commonHttpInfoLog.ReqDUrl = sensitiveUrl
	commonHttpInfoLog.ReqParams = string(j)
	params.GetLogger().WithFields(map[string]interface{}{
		"message":   "调用敏感词系统详细参数",
		"reqDUrl":   commonHttpInfoLog.ReqDUrl,
		"reqParams": commonHttpInfoLog.ReqParams,
		"reqCost":   duration,
		"reqCode":   response.Code,
	}).Info()

	// 3、 判断当前请求是否失败
	if resp.StatusCode != http.StatusOK {
		params.GetLogger().Error("NewRequestWithContext Err:", ",url:", sensitiveUrl, ",params:", string(j))
		return false, errors.New("Request Err:" + resp.Status)
	}
	// 4、 出参结果转换，判断是否包含敏感词
	if err := responses.GetJsonParser().NewDecoder(resp.Body).Decode(&response); err != nil {
		params.GetLogger().Error("Json Decode Err:", string(debug.Stack()))
		return false, err
	}

	//code 20000
	if response.Code != 10000 {
		jsonData, _ := responses.GetJsonParser().MarshalToString(response)
		params.GetLogger().Error("调用敏感词系统异常， ", "Data Err:", jsonData)
		return false, errors.New("Data Err: ")
	}
	// 判断是否过滤词语出现
	if response.Data.Sensitive {
		// 存在敏感词
		return false, nil
	}

	return true, nil
}

//使用参数请求数据
func (params *SensitiveWordApi) GetSensitiveApiResponse(p SensitiveWord) (bool, error) {
	params.Params = p
	//if params.Params insta {
	//	return false, nil
	//}
	real, err := params.GetResponse()

	if err != nil {
		return false, errors.New("url不合法")
	}
	return real, nil
}
