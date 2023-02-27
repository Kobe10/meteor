package model

import "meteor/global"

// 自增主键表 利用生成
type ShortUrl struct {
	UrlCode      string `json:"url_code" form:"url_code" gorm:"column:url_code;comment:url编码"`
	Uid          string `json:"uid" form:"uid" gorm:"column:uid;comment:用户uid"`
	Phone        string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号"`
	ShortUrl     string `json:"short_url" form:"short_url" gorm:"column:short_url;comment:短链地址"`
	AppId        string `json:"app_id" form:"app_id" gorm:"column:app_id;comment:应用id"`
	SourceUrlMd5 string `json:"source_url_md5" form:"source_url_md5" gorm:"column:source_url_md5;comment:长链地址MD5加密"`
	global.GVA_MODEL
}

func (ShortUrl) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "short_url"
}
