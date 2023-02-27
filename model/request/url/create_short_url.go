package url

// api分页条件查询及排序结构体
type CreateShortUrlReq struct {
	//request.AuthInfo
	SourceUrl string `json:"sourceUrl"`
}

// api分页条件查询及排序结构体
type ShortUrlDTO struct {
	SourceUrl    string `json:"sourceUrl"`
	SourceUrlMD5 string `json:"sourceUrlMD5"`
	Uid          string `json:"uid"`
	AppId        string `json:"appId"`
}
